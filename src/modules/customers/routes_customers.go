package routes

import (
	controller "norht-gowind/src/modules/customers/controllers"

	"github.com/gofiber/fiber/v2"
)

func LoadCustomerRoutes(router fiber.Router) {
	r := router.Group("/customers")
	r.Get("/", controller.GetAllCustomers)
}
