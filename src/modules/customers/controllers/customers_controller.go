package controller

import (
	service "norht-gowind/src/modules/customers/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllCustomers(c *fiber.Ctx) error {
	customers, err := service.GetAllCustomers()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error inesperado"})
	}

	return c.JSON(fiber.Map{
		"message":   "Lista de clientes obtenida exitosamente",
		"customers": customers,
	})
}
