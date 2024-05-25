package routes

import (
	controller "norht-gowind/src/modules/employees/controllers"

	"github.com/gofiber/fiber/v2"
)

func LoadEmployeesRoutes(router fiber.Router) {
	r := router.Group("employee")
	r.Get("/", controller.GetAllEmployees)
}
