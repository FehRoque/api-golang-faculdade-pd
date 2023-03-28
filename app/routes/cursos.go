package routes

import (
	"time"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCurso(c *fiber.Ctx) error {
	var payload *models.CreateCursoSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var faculdade models.Faculdade
	result := database.Database.Db.First(&faculdade, "id = ?", payload.FaculdadeID)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhuma faculdade existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newCurso := models.Curso{
		Nome:        payload.Nome,
		FaculdadeID: faculdade.ID,
	}

	result = database.Database.Db.Create(&newCurso)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"curso": newCurso}})
}

func DeleteCurso(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.Database.Db.Delete(&models.Curso{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func FindAlunosByCurso(c *fiber.Ctx) error {
	id := c.Params("id")

	var curso models.Curso
	result := database.Database.Db.First(&curso, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var alunos models.Alunos

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"curso": curso, "alunos": alunos}})
}

func GetCursoById(c *fiber.Ctx) error {
	id := c.Params("id")

	var curso models.Curso
	result := database.Database.Db.First(&curso, "id = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"curso": curso}})
}

func GetCursos(c *fiber.Ctx) error {
	var cursos []models.Curso

	results := database.Database.Db.Find(&cursos)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(cursos), "cursos": cursos})
}

func UpdateCurso(c *fiber.Ctx) error {
	id := c.Params("id")

	var payload *models.UpdateCursoSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var curso models.Curso
	result := database.Database.Db.First(&curso, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhum curso existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Nome != "" {
		updates["nome"] = payload.Nome
	}

	if payload.FaculdadeID != "" {
		var faculdade models.Faculdade
		result = database.Database.Db.First(&faculdade, "id = ?", payload.FaculdadeID)
		if err := result.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhuma faculdade existente com esse ID."})
			}
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		updates["faculdade_id"] = faculdade.ID
	}

	updates["updated_at"] = time.Now()

	database.Database.Db.Model(&curso).Updates(updates)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"curso": curso}})
}
