package models

import (
	"gorm.io/gorm"
)

type Faculdade struct {
	gorm.Model
	Nome string `json:"nome"`
	Cnpj string `json:"cnpj"`
}

type Faculdades []Faculdade

type CreateFaculdadeSchema struct {
	Nome string `json:"nome" validate:"required"`
	Cnpj string `json:"cnpj" validate:"required"`
}

type UpdateFaculdadeSchema struct {
	Nome string `json:"nome,omitempty"`
	Cnpj string `json:"cnpj,omitempty"`
}
