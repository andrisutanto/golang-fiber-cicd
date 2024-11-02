// main.go
package main

import (
	"golang-fiber-cicd/database"
	"golang-fiber-cicd/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/users", handlers.GetUsers)
	app.Get("/api/users/:id", handlers.GetUser)
	app.Post("/api/users", handlers.CreateUser)
	app.Put("/api/users/:id", handlers.UpdateUser)
	app.Delete("/api/users/:id", handlers.DeleteUser)
}

func main() {
	app := fiber.New()

	// Connect to the database
	database.Connect()

	// Setup routes
	setupRoutes(app)

	app.Listen(":3000")
}
