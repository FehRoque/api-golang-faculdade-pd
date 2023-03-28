package main

import (
	"log"

	"example.com/m/database"
	"example.com/m/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome API!")
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", welcome)

	alunos := api.Group("/alunos")

	alunos.Get("", routes.GetAlunos)
	alunos.Post("", routes.CreateAluno)

	alunos.Route("/:id", func(router fiber.Router) {
		router.Delete("", routes.DeleteAluno)
		router.Get("", routes.GetAlunoById)
		router.Get("/disciplinas", routes.FindDisciplinasByAluno)
		router.Patch("", routes.UpdateAluno)
	})

	cursos := api.Group("/cursos")

	cursos.Get("", routes.GetCursos)
	cursos.Post("", routes.CreateCurso)

	cursos.Route("/:id", func(router fiber.Router) {
		router.Delete("", routes.DeleteCurso)
		router.Get("", routes.GetCursoById)
		router.Get("/alunos", routes.FindAlunosByCurso)
		router.Patch("", routes.UpdateCurso)
	})

	cursoDisciplinas := api.Group("/curso_disciplinas")

	cursoDisciplinas.Get("", routes.GetCursoDisciplinas)
	cursoDisciplinas.Post("", routes.CreateCursoDisciplina)

	cursoDisciplinas.Route("/:id", func(router fiber.Router) {
		router.Delete("", routes.DeleteCursoDisciplina)
		router.Get("", routes.GetCursoDisciplinaById)
		router.Patch("", routes.UpdateCursoDisciplina)
	})

	disciplinas := api.Group("/disciplinas")

	disciplinas.Get("", routes.GetDisciplinas)
	disciplinas.Post("", routes.CreateDisciplina)

	disciplinas.Route("/:id", func(router fiber.Router) {
		router.Delete("", routes.DeleteDisciplina)
		router.Get("", routes.GetDisciplinaById)
		router.Get("/alunos", routes.FindAlunosByDisciplina)
		router.Patch("", routes.UpdateDisciplina)
	})

	disciplinaMatriculas := api.Group("/disciplina_matriculas")

	disciplinaMatriculas.Get("", routes.GetDisciplinaMatriculas)
	disciplinaMatriculas.Post("", routes.CreateDisciplinaMatricula)

	disciplinaMatriculas.Route("/:id", func(router fiber.Router) {
		router.Delete("", routes.DeleteDisciplinaMatricula)
		router.Get("", routes.GetDisciplinaMatriculaById)
		router.Patch("", routes.UpdateDisciplinaMatricula)
	})

	faculdades := api.Group("/faculdades")

	faculdades.Get("", routes.GetFaculdades)
	faculdades.Post("", routes.CreateFaculdade)

	faculdades.Route("/:id", func(router fiber.Router) {
		router.Delete("", routes.DeleteFaculdade)
		router.Get("", routes.GetFaculdadeById)
		router.Get("/alunos", routes.FindAlunosByFaculdade)
		router.Get("/cursos", routes.FindCursosByFaculdade)
		router.Patch("", routes.UpdateFaculdade)
	})

	professores := api.Group("/professores")

	professores.Get("", routes.GetProfessores)
	professores.Post("", routes.CreateProfessor)

	professores.Route("/:id", func(router fiber.Router) {
		router.Delete("", routes.DeleteProfessor)
		router.Get("", routes.GetProfessorById)
		router.Patch("", routes.UpdateProfessor)
	})
}

func main() {
	database.ConnectDb()
	database.Setup(database.Database.Db)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
