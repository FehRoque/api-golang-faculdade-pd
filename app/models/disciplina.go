package models

import "gorm.io/gorm"

type Disciplina struct {
	gorm.Model
	Nome        string   `json:"nome"`
	ProfessorID uint     `gorm:"foreignKey:ProfessorID" json:"professor_id"`
	Curso       []*Curso `gorm:"many2many:curso_disciplinas;"`
}

type Disciplinas []Disciplina

type CreateDisciplinaSchema struct {
	Nome        string `json:"nome" validate:"required"`
	ProfessorID string `json:"professor_id"`
}

type UpdateDisciplinaSchema struct {
	Nome        string `json:"nome,omitempty"`
	ProfessorID string `json:"professor_id,omitempty"`
}
