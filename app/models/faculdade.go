package models

import (
	"gorm.io/gorm"
)

type Faculdade struct {
	gorm.Model
	Nome   string `json:"nome"`
	Cnpj   string `json:"cnpj"`
	Cursos Cursos `json:"cursos"`
}

type Faculdades struct {
	Faculdades []Faculdade `json:"faculdades"`
}

type CreateFaculdadeSchema struct {
	Nome   string `json:"nome" validate:"required"`
	Cnpj   string `json:"cnpj" validate:"required"`
	Cursos Cursos `json:"cursos,omitempty"`
}

type UpdateFaculdadeSchema struct {
	Nome   string `json:"nome,omitempty"`	
	Cnpj   string `json:"cnpj,omitempty"`	
	Cursos Cursos `json:"cursos,omitempty"`	
}

