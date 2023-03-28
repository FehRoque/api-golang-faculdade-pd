package database

import (
	"fmt"
	"log"
	"os"

	"example.com/m/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "postgres"
	user     = "postgres"
	password = "password"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable", user, password, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}
	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)

	Database = DbInstance{
		Db: db,
	}
}

func Setup(db *gorm.DB) {
	db.AutoMigrate(&models.Aluno{}, &models.Curso{}, &models.Disciplina{}, &models.DisciplinaMatricula{}, &models.Faculdade{}, &models.Professor{})
	seed(db)
}

func seed(db *gorm.DB) {
	faculdades := models.Faculdades{
		{
			Nome: "Faculdade PD - MG",
			Cnpj: "123456789",
		},
		{
			Nome: "Faculdade PD - SP",
			Cnpj: "154354",
		},
		{
			Nome: "Faculdade PD - Passei Direto (EAD)",
			Cnpj: "7070.777",
		},
	}

	for _, faculdade := range faculdades {
		db.Create(&faculdade)
	}

	var passeiDiretoEAD, passeiDiretoSP, passeiDiretoMG models.Faculdade

	db.First(&passeiDiretoEAD, "Nome = ?", "Faculdade PD - Passei Direto (EAD)")
	db.First(&passeiDiretoSP, "Nome = ?", "Faculdade PD - SP")
	db.First(&passeiDiretoMG, "Nome = ?", "Faculdade PD - MG")

	professores := models.Professores{
		{Nome: "Felipe Roque", Formacao: "Graduado"},
		{Nome: "Vinicius Roque", Formacao: "Mestrado"},
	}

	for _, professor := range professores {
		db.Create(&professor)
	}

	var ProfessorFelipe, ProfessorVinicius models.Professor
	db.First(&ProfessorFelipe, "Nome = ?", "Felipe Roque")
	db.First(&ProfessorVinicius, "Nome = ?", "Vinicius Roque")

	cursos := models.Cursos{
		{Nome: "Análise e Desenvolvimento de Sistemas", FaculdadeID: passeiDiretoSP.ID},
		{Nome: "Análise e Desenvolvimento de Sistemas", FaculdadeID: passeiDiretoMG.ID},
		{Nome: "Análise e Desenvolvimento de Sistemas (EAD)", FaculdadeID: passeiDiretoEAD.ID},
		{Nome: "Robótica", FaculdadeID: passeiDiretoMG.ID},
		{Nome: "Robótica (EAD)", FaculdadeID: passeiDiretoEAD.ID},
		{Nome: "RPA Development", FaculdadeID: passeiDiretoSP.ID},
	}

	for _, curso := range cursos {
		db.Create(&curso)
	}

	var cursoAdsSP, cursoAdsEad, cursoRoboticaEad, cursoRoboticaMG models.Curso

	db.First(&cursoAdsSP, "Nome = ? AND faculdade_id = ?", "Análise e Desenvolvimento de Sistemas", passeiDiretoSP.ID)
	db.First(&cursoAdsEad, "Nome = ?", "Análise e Desenvolvimento de Sistemas (EAD)")
	db.First(&cursoRoboticaEad, "Nome = ?", "Robótica (EAD)")
	db.First(&cursoRoboticaMG, "Nome = ? AND faculdade_id = ?", "Robótica", passeiDiretoMG.ID)

	disciplinas := models.Disciplinas{
		{Nome: "Lógica de Programação", ProfessorID: ProfessorVinicius.ID},
		{Nome: "Programação Avançada", ProfessorID: ProfessorFelipe.ID},
		{Nome: "Robótica", ProfessorID: ProfessorFelipe.ID},
		{Nome: "Robótica Avançada", ProfessorID: ProfessorFelipe.ID},
		{Nome: "RPA Básico", ProfessorID: ProfessorFelipe.ID},
		{Nome: "RPA Avançado", ProfessorID: ProfessorVinicius.ID},
	}

	for _, disciplina := range disciplinas {
		db.Create(&disciplina)
	}

	var disciplinaLogicaProgramacao, disciplinaProgramacaoAvancada, disciplinaRobotica, disciplinaRoboticaAvancada, disciplinaRPABasico, disciplinaRPAAvancado models.Disciplina

	db.First(&disciplinaLogicaProgramacao, "Nome = ?", "Lógica de Programação")
	db.First(&disciplinaProgramacaoAvancada, "Nome = ?", "Programação Avançada")
	db.First(&disciplinaRobotica, "Nome = ?", "Robótica")
	db.First(&disciplinaRoboticaAvancada, "Nome = ?", "Robótica Avançada")
	db.First(&disciplinaRPABasico, "Nome = ?", "RPA Básico")
	db.First(&disciplinaRPAAvancado, "Nome = ?", "RPA Avançado")

	alunos := models.Alunos{
		{Nome: "Felipe Ramos Roque", Cpf: "987654321"},
		{Nome: "Nicolas Roque", Cpf: "123456789"},
		{Nome: "Matheus Brasileiro", Cpf: "30-333-777.666"},
	}

	for _, aluno := range alunos {
		db.Create(&aluno)
	}

	var Felipe, Nicolas models.Aluno
	db.First(&Felipe, "Nome = ?", "Felipe Ramos Roque")
	db.First(&Nicolas, "Nome = ?", "Nicolas Roque")
}
