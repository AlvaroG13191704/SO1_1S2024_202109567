package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Initialize default config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	app.Get("/data", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name": "Alvaro Norberto García Meza",
			"uid":  "202109567",
		})
	})

	app.Listen(":8080")
}
