package routes

import (
	"time"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func CreateAluno(c *fiber.Ctx) error {
	var payload *models.CreateAlunoSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	
	newAluno := models.Aluno {
		Nome: payload.Nome,
		Cpf: payload.Cpf,
	}

	result := database.Database.Db.Create(&newAluno)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"aluno": newAluno}})
}

func DeleteAluno(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.Database.Db.Delete(&models.Aluno{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhum aluno existente com esse ID."})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func FindDisciplinasByAluno(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func GetAlunoById(c *fiber.Ctx) error {
	id := c.Params("id")

	var aluno models.Aluno
	result := database.Database.Db.First(&aluno, "id = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhum aluno existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"aluno": aluno}})
}

func GetAlunos(c *fiber.Ctx) error {
	var alunos []models.Aluno
	
	results := database.Database.Db.Find(&alunos)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(alunos), "alunos": alunos})
}

func UpdateAluno(c *fiber.Ctx) error {
	id := c.Params("id")

	var payload *models.UpdateAlunoSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	
	var aluno models.Aluno
	result := database.Database.Db.First(&aluno, "id = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhum aluno existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Nome != "" {
		updates["nome"] = payload.Nome
	}

	if payload.Cpf != "" {
		updates["cpf"] = payload.Cpf
	}

	updates["updated_at"] = time.Now()

	database.Database.Db.Model(&aluno).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"aluno": aluno}})
}
