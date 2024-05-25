package main

import (
	"norht-gowind/src/api/routes"
	config_global "norht-gowind/src/common/configs"
	"norht-gowind/src/common/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config_global.LoadConfig()

	app := fiber.New()

	database.Connect()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
