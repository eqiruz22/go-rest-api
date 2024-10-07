package main

import (
	"fiber/backend/database"
	"fiber/backend/database/migration"
	"fiber/backend/routes"
	"fiber/backend/validation"

	_ "fiber/backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title My API
// @version 1.0
// @description This is a sample API.
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()
	validation.InitValidator()
	routes.RouteAuth(app)
	routes.RouteUser(app)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Listen(":8081")
}