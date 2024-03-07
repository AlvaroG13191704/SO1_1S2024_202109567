package router

import (
	"sopes1/proyecto1/util"

	"github.com/gofiber/fiber/v2"
)

func GetProcesses(c *fiber.Ctx) error {
	cpu, err := util.GetCPU(false)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"processes": cpu.Processes,
	})
}
