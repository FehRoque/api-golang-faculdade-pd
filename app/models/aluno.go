package models

import (
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome                 string                `json:"nome"`
	Cpf                  string                `json:"cpf"`
	DisciplinaMatriculas []DisciplinaMatricula `json:"disciplina_matricula"`
}

type Alunos []Aluno

type CreateAlunoSchema struct {
	Nome string `json:"nome" validate:"required"`
	Cpf  string `json:"cpf" validate:"required"`
}

type UpdateAlunoSchema struct {
	Nome string `json:"nome,omitempty"`
	Cpf  string `json:"cpf,omitempty"`
}
