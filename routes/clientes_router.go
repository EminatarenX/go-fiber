package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber/controllers"
	
)


func ClientesRouter (app *fiber.App) {
	client := app.Group("/clientes")

	client.Get("/", controllers.ObterClientes)
	
	
}