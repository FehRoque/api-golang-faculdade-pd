package routes

import (
	"time"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateDisciplina(c *fiber.Ctx) error {
	var payload *models.CreateDisciplinaSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var professor models.Professor
	result := database.Database.Db.First(&professor, "id = ?", payload.ProfessorID)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newDisciplina := models.Disciplina{
		Nome:        payload.Nome,
		ProfessorID: professor.ID,
	}

	result = database.Database.Db.Create(&newDisciplina)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"disciplina": newDisciplina}})
}

func DeleteDisciplina(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.Database.Db.Delete(&models.Disciplina{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func FindAlunosByDisciplina(c *fiber.Ctx) error {
	id := c.Params("id")

	var disciplina models.Disciplina
	result := database.Database.Db.First(&disciplina, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var alunos models.Alunos

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"disciplina": disciplina, "alunos": alunos}})
}

func GetDisciplinaById(c *fiber.Ctx) error {
	id := c.Params("id")
	var disciplina models.Disciplina

	result := database.Database.Db.First(&disciplina, "id = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"disciplina": disciplina}})
}

func GetDisciplinas(c *fiber.Ctx) error {
	var disciplinas []models.Disciplina

	results := database.Database.Db.Find(&disciplinas)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(disciplinas), "disciplinas": disciplinas})
}

func UpdateDisciplina(c *fiber.Ctx) error {
	id := c.Params("id")
	var payload *models.UpdateDisciplinaSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var disciplina models.Disciplina

	result := database.Database.Db.First(&disciplina, "id = ?", id)
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

	if payload.ProfessorID != "" {
		updates["professor_id"] = payload.ProfessorID
	}

	updates["updated_at"] = time.Now()
	database.Database.Db.Model(&disciplina).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"disciplina": disciplina}})
}
