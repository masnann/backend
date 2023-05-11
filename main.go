package main

import (
	"backend/migration"
	"backend/route"
	"backend/utils/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":8080")

}
