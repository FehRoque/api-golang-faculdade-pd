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
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "password"
	dbname    = "pd"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable", user, password, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
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
	db.AutoMigrate(&models.Aluno{}, &models.Curso{}, &models.Disciplina{}, &models.Faculdade{}, &models.Professor{})
	// seed(db)
}


func seed(db *gorm.DB) {
	faculdades := []models.Faculdade {
		{
			Nome: "Faculdade PD - Passei Direto",
			Cnpj: "7070.777",
		},
		{
			Nome: "PD",
			Cnpj: "154354",
		},
	}
	
	for _, faculdade := range faculdades {
		db.Create(&faculdade)
	}
				
	var passeiDireto, PD models.Faculdade
	db.First(&passeiDireto, "Nome = ?", "Passei Direto")
	db.First(&PD, "Nome = ?", "PD")

	alunos := []models.Aluno {
		{ Nome: "Felipe Ramos Roque", Cpf: "486.681.328-81"},
		{ Nome: "Nicolas Roque", Cpf: "12345678"},
		{ Nome: "Matuê", Cpf: "30-333-777.666"},
	}

	for _, aluno := range alunos {
		db.Create(&aluno)
	}

	var Felipe, Nicolas models.Aluno
	db.First(&Felipe, "Nome = ?", "Felipe")
	db.First(&Nicolas, "Nome = ?", "Nicolas")

	professores := []models.Professor {
		{ Nome: "Felipe Roque", Formacao: "Graduado"},
		{ Nome: "Vinicius Roque", Formacao: "Mestrado"},
		{ Nome: "Matuê", Formacao: "Trapper"},
	}

	for _, professor := range professores {
		db.Create(&professor)
	}

	var ProfessorFelipe, ProfessorVinicius models.Professor
	db.First(&ProfessorFelipe, "Nome = ?", "Felipe Roque")
	db.First(&ProfessorVinicius, "Nome = ?", "Vinicius Roque")



	// disciplinas := []models.Disciplina{
	// 	{Nome: "Programação Avançada", ProfessorID: ProfessorFelipe.ID},
	// 	{Nome: "Robótica", ProfessorID: ProfessorFelipe.ID},
	// 	{Nome: "Lógica de Programação", ProfessorID: ProfessorVinicius.ID},
	// 	{Nome: "RPA", ProfessorID: ProfessorVinicius.ID},
	// }

	// for _, disciplina := range disciplinas {
	// 	db.Create(&disciplina)
	// }
}
