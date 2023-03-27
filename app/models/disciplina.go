package models

import "gorm.io/gorm"

type Disciplina struct {
	gorm.Model
	Nome        string `json:"nome"`
	Professores Professores   `json:"professores"`
}

type Disciplinas struct {
	Disciplinas []Disciplina `json:"disciplinas"`
}

type CreateDisciplinaSchema struct {
	Nome string `json:"nome" validate:"required"`
	Professores Professores `json:"professores"`
}

type UpdateDisciplinaSchema struct {
	Nome string `json:"nome,omitempty"`
	Professores Professores `json:"professores,omitempty"`
}