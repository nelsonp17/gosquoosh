package gosquoosh

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonp17/gosquoosh/routers"
)

func server(serverConfig ServerConfig) {
	app := fiber.New()

	// Servir el archivo robots.txt
	app.Static("/robots.txt", "./robots.txt")

	// routes
	router := app.Group(serverConfig.ENDPOINT_EXPONE) // <- /api/v1/image-convert

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	// app.Get("/robots.txt", func(c *fiber.Ctx) error {
	// 	return c.SendFile("./robots.txt")
	// })

	// MICROSERVICIOS
	routers.Init(router)

	// start server
	err := app.Listen(serverConfig.HOST + ":" + serverConfig.PORT)
	if err != nil {
		return
	}
}
