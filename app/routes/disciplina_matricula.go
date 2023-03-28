package routes

import (
	"time"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateDisciplinaMatricula(c *fiber.Ctx) error {
	var payload *models.CreateDisciplinaMatriculaSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var aluno models.Aluno
	result := database.Database.Db.First(&aluno, "id = ?", payload.AlunoID)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhum Curso/Disciplina existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var cursoDisciplina models.CursoDisciplina
	result = database.Database.Db.First(&cursoDisciplina, "id = ?", payload.CursoDisciplinaID)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhum Curso/Disciplina existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newDisciplinaMatricula := models.DisciplinaMatricula{
		AlunoID:           aluno.ID,
		CursoDisciplinaID: cursoDisciplina.ID,
	}

	result = database.Database.Db.Create(&newDisciplinaMatricula)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"disciplina_matricula": newDisciplinaMatricula}})
}

func DeleteDisciplinaMatricula(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.Database.Db.Delete(&models.DisciplinaMatricula{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhum aluno existente com esse ID."})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func GetDisciplinaMatriculaById(c *fiber.Ctx) error {
	id := c.Params("id")
	var disciplinaMatricula models.DisciplinaMatricula

	result := database.Database.Db.First(&disciplinaMatricula, "id = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Nenhuma disciplina existente com esse ID."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"disciplina_matricula": disciplinaMatricula}})
}

func GetDisciplinaMatriculas(c *fiber.Ctx) error {
	var disciplinaMatriculas models.DisciplinaMatriculas

	results := database.Database.Db.Find(&disciplinaMatriculas)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(disciplinaMatriculas), "disciplina_matriculas": disciplinaMatriculas})
}

func UpdateDisciplinaMatricula(c *fiber.Ctx) error {
	id := c.Params("id")
	var payload *models.UpdateDisciplinaMatriculaSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var disciplinaMatricula models.DisciplinaMatricula

	result := database.Database.Db.First(&disciplinaMatricula, "id = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "ID not found."})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})

	if payload.CursoDisciplinaID != "" {
		var cursoDisciplina models.CursoDisciplina

		result = database.Database.Db.First(&cursoDisciplina, "id = ?", payload.CursoDisciplinaID)
		if err := result.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "CursoDisciplinaID not found."})
			}
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		updates["curso_disciplina_id"] = payload.CursoDisciplinaID
	}

	if payload.AlunoID != "" {
		var aluno models.Aluno

		result = database.Database.Db.First(&aluno, "id = ?", payload.AlunoID)
		if err := result.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "AlunoID not found."})
			}
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		updates["aluno_id"] = payload.AlunoID
	}

	updates["updated_at"] = time.Now()
	database.Database.Db.Model(&disciplinaMatricula).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"disciplina_matricula": disciplinaMatricula}})
}
