package models

import "gorm.io/gorm"

type CursoDisciplina struct {
	gorm.Model
	Curso      Curso      `gorm:"foreignKey:CursoRefer"`
	Disciplina Disciplina `gorm:"foreignKey:DisciplinaRefer"`
}
