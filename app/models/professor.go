package models

import (
	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model
	Nome     string `json:"nome"`
	Formacao string `json:"formacao"`
}

type Professores struct {
	Professores []Professor `json:"professores"`
}

type CreateProfessorSchema struct {
	Nome     string `json:"nome" validate:"required"`
	Formacao string `json:"formacao"`
}


type UpdateProfessorSchema struct {
	Nome     string `json:"nome,omitempty"`
	Formacao string `json:"formacao,omitempty"`
}
