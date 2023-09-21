package controllers

import (
	"github.com/gofiber/fiber/v2"
	"fiber/models"

)

func ObterClientes(c *fiber.Ctx) error {

	cliente := models.Cliente{
		ID: 1,
		Nombre: "João",
		Apellidos: "Silva",
		Email: "jao@correo.com",
		CreatedAt: "2021-09-01 00:00:00",
	}

	return c.JSON(fiber.Map{
		"cliente": cliente,
	})

}