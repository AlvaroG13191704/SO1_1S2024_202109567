package router

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func StartProcess(c *fiber.Ctx) error {
	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al iniciar el proceso",
		})
	}

	process := cmd
	return c.Status(200).JSON(fiber.Map{
		"message": "Proceso iniciado",
		"pid":     process.Process.Pid,
	})
}

func StopProcess(c *fiber.Ctx) error {
	pid := c.Query("pid")

	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "PID debe ser un número entero",
		})
	}

	cmd := exec.Command("kill", "-SIGSTOP", strconv.Itoa(pidInt))
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al detener el proceso",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Proceso detenido",
		"pid":     pid,
	})
}

func ResumeProcess(c *fiber.Ctx) error {
	pid := c.Query("pid")

	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "PID debe ser un número entero",
		})
	}

	cmd := exec.Command("kill", "-SIGCONT", strconv.Itoa(pidInt))
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al reanudar el proceso",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Proceso reanudado",
		"pid":     pid,
	})
}

func KillProcess(c *fiber.Ctx) error {
	pid := c.Query("pid")

	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "PID debe ser un número entero",
		})
	}

	cmd := exec.Command("kill", "-9", strconv.Itoa(pidInt))
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al matar el proceso",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Proceso matado",
		"pid":     pid,
	})
}
