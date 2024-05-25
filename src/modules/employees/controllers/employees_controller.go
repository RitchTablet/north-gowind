package controllers

import (
	service "norht-gowind/src/modules/employees/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllEmployees(c *fiber.Ctx) error {
	employees, err := service.GetAllEmployees()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error inesperado"})
	}

	return c.JSON(fiber.Map{
		"message":   "Lista de empleados obtenida exitosamente",
		"employees": employees,
	})
}
