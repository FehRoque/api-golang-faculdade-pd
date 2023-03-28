package routes

import (
	"time"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProfessor(c *fiber.Ctx) error {
	var payload *models.CreateProfessorSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newProfessor := models.Professor{
		Nome:     payload.Nome,
		Formacao: payload.Formacao,
	}

	result := database.Database.Db.Create(&newProfessor)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"professor": newProfessor}})
}

func DeleteProfessor(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.Database.Db.Delete(&models.Professor{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func GetProfessorById(c *fiber.Ctx) error {
	id := c.Params("id")

	var professor models.Professor
	result := database.Database.Db.First(&professor, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"professor": professor}})
}

func GetProfessores(c *fiber.Ctx) error {
	var professores []models.Professor

	results := database.Database.Db.Find(&professores)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(professores), "professores": professores})
}

func UpdateProfessor(c *fiber.Ctx) error {
	id := c.Params("id")

	var payload *models.UpdateProfessorSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var professor models.Professor
	result := database.Database.Db.First(&professor, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Nome != "" {
		updates["nome"] = payload.Nome
	}

	if payload.Formacao != "" {
		updates["formacao"] = payload.Formacao
	}

	updates["updated_at"] = time.Now()
	database.Database.Db.Model(&professor).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"professor": professor}})
}
