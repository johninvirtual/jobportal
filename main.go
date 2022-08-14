package main

import (
	"embed"
	"jobportal/database"
	"jobportal/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// Embed a directory
//
//go:embed ui/build/*
var embedBuildDir embed.FS

func setUpRoutes(app *fiber.App) {
	// api routes
	app.Post("/job", routes.AddJob)
	app.Get("/job", routes.ListJob)

	// serving ui build
	app.Use("", filesystem.New(filesystem.Config{
		Root:       http.FS(embedBuildDir),
		PathPrefix: "ui/build",
	}))

	// cors
	app.Use(cors.New())
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	log.Fatal(app.Listen("localhost:3000"))
}
