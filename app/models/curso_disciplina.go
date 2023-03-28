package models

import "gorm.io/gorm"

type CursoDisciplina struct {
	gorm.Model
	CursoID      uint `gorm:"primaryKey" json:"curso_id"`
	DisciplinaID uint `gorm:"primaryKey" json:"disciplina_id"`
}

type CursoDisciplinas []CursoDisciplina

type CreateCursoDisciplinaSchema struct {
	CursoID      string `json:"curso_id" validate:"required"`
	DisciplinaID string `json:"disciplina_id" validate:"required"`
}

type UpdateCursoDisciplinaSchema struct {
	CursoID      string `json:"curso_id,omitempty"`
	DisciplinaID string `json:"disciplina_id,omitempty"`
}
