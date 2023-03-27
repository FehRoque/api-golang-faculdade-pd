package models

import (
	"time"

	"gorm.io/gorm"
)

type DisciplinaMatricula struct {
	gorm.Model
	CursoDisciplina CursoDisciplina `gorm:"foreignKey:CursoDisciplinaRefer"`
	Alunos Alunos
	DataMatricula time.Time `json:"data_matricula"`
}
