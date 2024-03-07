package main

import (
	"sopes1/proyecto1/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// add cors
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type")
		return c.Next()
	})

	// create a fiber group
	api := app.Group("/api")
	api.Get("/real-time", router.RealTime)
	api.Get("/get-history", router.History)
	api.Get("/get-processes", router.GetProcesses)

	// proccess status
	api.Get("/process/start", router.StartProcess)
	api.Get("/process/stop", router.StopProcess)
	api.Get("/process/resume", router.ResumeProcess)
	api.Get("/process/kill", router.KillProcess)

	app.Listen(":8080")
}
