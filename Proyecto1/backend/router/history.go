package router

import (
	"sopes1/proyecto1/util"

	"github.com/gofiber/fiber/v2"
)

func History(c *fiber.Ctx) error {
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
}
