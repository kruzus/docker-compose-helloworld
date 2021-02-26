package main

import (
	"godock/data"

	"github.com/gofiber/fiber/v2"
)

func main() {

	jsonData := data.AllPlayers.NewList(&data.Bob, &data.Elise, &data.Maria, &data.Steve)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&jsonData)
	})

	app.Listen(":9000")
}
