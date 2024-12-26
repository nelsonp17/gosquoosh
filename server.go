package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonp17/gosquoosh/routers"
)

func server(serverConfig ServerConfig) {
	app := fiber.New()

	// routes
	router := app.Group(serverConfig.ENDPOINT_EXPONE) // <- /api/v1/image-convert

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// MICROSERVICIOS
	routers.Init(router)

	// start server
	err := app.Listen(serverConfig.HOST + ":" + serverConfig.PORT)
	if err != nil {
		return
	}
}
