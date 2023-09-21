package main

import(
	"github.com/gofiber/fiber/v2"
	"fiber/routes"
)


func main() {

	app := fiber.New()

	routes.ClientesRouter(app)

	app.Listen(":4000")
}

