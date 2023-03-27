package models

import (
	"gorm.io/gorm"
)


type Curso struct {
	gorm.Model
	Nome         string `json:"nome"`
	FaculdadeId  uint `json:"faculdade_id"`
	Faculdade  Faculdade `json:"faculdade"`
}

type Cursos struct {
	Cursos []Curso `json:"cursos"`
}

type CreateCursoSchema struct {
	Nome string `json:"nome" validate:"required"`
	FaculdadeId uint `json:"faculdade_id" validate:"required"`
}

type UpdateCursoSchema struct {
	Nome string `json:"nome,omitempty"`
	FaculdadeId uint `json:"faculdade_id,omitempty"`
}