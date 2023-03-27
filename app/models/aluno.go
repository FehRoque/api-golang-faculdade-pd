package models

import (
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	Cpf  string `json:"cpf"`
}

type Alunos struct {
	Alunos []Aluno `json:"alunos"`
}

type CreateAlunoSchema struct {
	Nome  string `json:"nome" validate:"required"`
	Cpf   string `json:"cpf" validate:"required"`
}

type UpdateAlunoSchema struct {
	Nome  string `json:"nome,omitempty"`
	Cpf   string `json:"cpf,omitempty"`
}
