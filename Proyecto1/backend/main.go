package main

import (
	"log"
	"sopes1/proyecto1/models"
	"sopes1/proyecto1/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/real-time", func(c *fiber.Ctx) error {

		ramCh := make(chan models.Ram, 1)
		cpuCh := make(chan models.Cpu, 1)
		errCh := make(chan error, 1)

		go func() {
			ram, err := util.GetRAM()
			if err != nil {
				errCh <- err
				return
			}
			ramCh <- ram
		}()

		go func() {
			cpu, err := util.GetCPU(true)
			if err != nil {
				errCh <- err
				return
			}
			cpuCh <- cpu
		}()

		var ram models.Ram
		var cpu models.Cpu
		for i := 0; i < 2; i++ {
			select {
			case ram = <-ramCh:
			case cpu = <-cpuCh:
			case err := <-errCh:
				log.Fatalf("Error fetching data: %s\n", err)
				return c.JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}
		// read from kernel
		return c.JSON(fiber.Map{
			"ram": ram,
			"cpu": map[string]string{
				"percentage": cpu.PercentCPU,
			},
		})
	})

	app.Get("/get-history", func(c *fiber.Ctx) error {
		ramHistory, err := util.GetHistoryRam()
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		cpuHistory, err := util.GetHistoryCPU()
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"ram": ramHistory,
			"cpu": cpuHistory,
		})
	})

	app.Get("/get-processes", func(c *fiber.Ctx) error {
		cpu, err := util.GetCPU(false)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"processes": cpu.Processes,
		})
	})

	app.Listen(":8080")
}
