package routes

import (
	routes_customers "norht-gowind/src/modules/customers"
	routes "norht-gowind/src/modules/echo"
	routes_employees "norht-gowind/src/modules/employees"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	router := app.Group("/api/v1")
	routes.LoadEchoRoutes(router)
	routes_customers.LoadCustomerRoutes(router)
	routes_employees.LoadEmployeesRoutes(router)
}
