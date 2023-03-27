package routes

import (
	"time"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func CreateFaculdade(c *fiber.Ctx) error {
	var payload *models.CreateFaculdadeSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}
	
	newFaculdade := models.Faculdade {
		Nome: payload.Nome,
		Cnpj: payload.Cnpj,
		Cursos: payload.Cursos,
	}

	result := database.Database.Db.Create(&newFaculdade)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"faculdade": newFaculdade}})
}

func DeleteFaculdade(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.Database.Db.Delete(&models.Faculdade{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhuma faculdade existente com esse ID."})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func FindAlunosByFaculdade(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func FindCursosByFaculdade(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func GetFaculdadeById(c *fiber.Ctx) error {
	id := c.Params("id")

	var faculdade models.Faculdade
	result := database.Database.Db.First(&faculdade, "id = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhuma faculdade existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"faculdade": faculdade}})
}


func GetFaculdades(c *fiber.Ctx) error {
	var faculdades []models.Faculdade
	
	results := database.Database.Db.Find(&faculdades)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(faculdades), "faculdades": faculdades})
}

func GetTotalAlunosFaculdade(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateFaculdade(c *fiber.Ctx) error {
	id := c.Params("id")
	var payload *models.UpdateFaculdadeSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	
	var faculdade models.Faculdade
	result := database.Database.Db.First(&faculdade, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhuma faculdade existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Nome != "" {
		updates["nome"] = payload.Nome
	}

	if payload.Cnpj != "" {
		updates["cnpj"] = payload.Cnpj
	}

	if len(payload.Cursos.Cursos) != 0 {
		updates["cursos"] = payload.Cursos
	}

	updates["updated_at"] = time.Now()
	database.Database.Db.Model(&faculdade).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"faculdade": faculdade}})
}
