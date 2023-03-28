package models

import (
	"time"
)

type DisciplinaMatricula struct {
	AlunoID           uint      `gorm:"primaryKey" json:"aluno_id"`
	CursoDisciplinaID uint      `gorm:"primaryKey" json:"curso_disciplina_id"`
	DataMatricula     time.Time `json:"data_matricula"`
}

type DisciplinaMatriculas []DisciplinaMatricula

type CreateDisciplinaMatriculaSchema struct {
	CursoDisciplinaID string `json:"curso_disciplina_id" validate:"required"`
	AlunoID           string `json:"aluno_id" validate:"required"`
}

type UpdateDisciplinaMatriculaSchema struct {
	CursoDisciplinaID string `json:"curso_disciplina_id,omitempty"`
	AlunoID           string `json:"aluno_id,omitempty"`
}
