package controller

import (
	service "norht-gowind/src/modules/echo/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllEchos(c *fiber.Ctx) error {
	echos, err := service.GetAllEchos()

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "Error innesperado"})
	}

	return c.JSON(fiber.Map{
		"message": "echo!",
		"echos":   echos,
	})
}
