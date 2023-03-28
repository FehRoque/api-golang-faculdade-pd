package models

import (
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nome        string        `json:"nome"`
	FaculdadeID uint          `gorm:"foreignKey:FaculdadeID" json:"faculdade"`
	Disciplinas []*Disciplina `gorm:"many2many:curso_disciplinas;" json:"disciplinas"`
}

type Cursos []Curso

type CreateCursoSchema struct {
	Nome        string `json:"nome" validate:"required"`
	FaculdadeID string `json:"faculdade_id"`
}

type UpdateCursoSchema struct {
	Nome        string `json:"nome,omitempty"`
	FaculdadeID string `json:"faculdade_id,omitempty"`
}
