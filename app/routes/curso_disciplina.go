package routes

import (
	"time"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCursoDisciplina(c *fiber.Ctx) error {
	var payload *models.CreateCursoDisciplinaSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var curso models.Curso
	result := database.Database.Db.First(&curso, "id = ?", payload.CursoID)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "CursoID not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var disciplina models.Disciplina
	result = database.Database.Db.First(&disciplina, "id = ?", payload.DisciplinaID)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "DisciplinaID not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newCursoDisciplina := models.CursoDisciplina{
		CursoID:      curso.ID,
		DisciplinaID: disciplina.ID,
	}

	result = database.Database.Db.Create(&newCursoDisciplina)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"curso_disciplina": newCursoDisciplina}})
}

func DeleteCursoDisciplina(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.Database.Db.Delete(&models.Curso{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func GetCursoDisciplinaById(c *fiber.Ctx) error {
	id := c.Params("id")

	var cursoDisciplina models.CursoDisciplina
	result := database.Database.Db.First(&cursoDisciplina, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "id not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"curso_disciplina": cursoDisciplina}})
}

func GetCursoDisciplinas(c *fiber.Ctx) error {
	var cursoDisciplinas []models.CursoDisciplina

	results := database.Database.Db.Find(&cursoDisciplinas)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(cursoDisciplinas), "curso_disciplinas": cursoDisciplinas})
}

func UpdateCursoDisciplina(c *fiber.Ctx) error {
	id := c.Params("id")
	var payload *models.UpdateCursoDisciplinaSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var cursoDisciplina models.CursoDisciplina
	result := database.Database.Db.First(&cursoDisciplina, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhum curso existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})

	if payload.CursoID != "" {
		var curso models.Curso

		result = database.Database.Db.First(&curso, "id = ?", payload.CursoID)
		if err := result.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "CursoID not found."})
			}
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		updates["curso_id"] = payload.CursoID
	}

	if payload.DisciplinaID != "" {
		var disciplina models.Disciplina

		result = database.Database.Db.First(&disciplina, "id = ?", payload.DisciplinaID)
		if err := result.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "DisciplinaID noot found."})
			}
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		updates["disciplina_id"] = disciplina.ID
	}

	updates["updated_at"] = time.Now()

	database.Database.Db.Model(&cursoDisciplina).Updates(updates)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"curso_disciplina": cursoDisciplina}})
}
