package routes

import (
	controller "norht-gowind/src/modules/echo/controllers"

	"github.com/gofiber/fiber/v2"
)

func LoadEchoRoutes(router fiber.Router) {
	r := router.Group("echo")
	r.Get("/", controller.GetAllEchos)
}
