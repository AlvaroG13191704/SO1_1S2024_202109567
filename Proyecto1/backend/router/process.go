package router

import (
	"log"
	"os/exec"
	"sopes1/proyecto1/db"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func StartProcess(c *fiber.Ctx) error {
	log.Println("Starting process")

	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al iniciar el proceso",
		})
	}

	process := cmd

	// save to db
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al instanciar db",
		})
	}
	// insert to db
	_, err = dbClient.Exec(
		"INSERT INTO STATUS_PROCESS (pid, status, created_at) VALUES (?, ?, ?)",
		process.Process.Pid, "running", time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		log.Fatalf("Error al insertar proceso en db: %v", err)

	}

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

	// save to db
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al instanciar db",
		})
	}
	// update to db
	_, err = dbClient.Exec(
		"UPDATE STATUS_PROCESS SET status = ? WHERE pid = ?",
		"stopped", pidInt)

	if err != nil {
		log.Fatalf("Error al insertar proceso en db: %v", err)
	}

	log.Println("Process stopped")

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
	// save to db
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al instanciar db",
		})
	}
	// update to db
	_, err = dbClient.Exec(
		"UPDATE STATUS_PROCESS SET status = ? WHERE pid = ?",
		"running", pidInt)

	if err != nil {
		log.Fatalf("Error al insertar proceso en db: %v", err)
	}

	log.Println("Process resumed")

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

	// save to db
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Error al instanciar db",
		})
	}
	// update to db
	_, err = dbClient.Exec(
		"UPDATE STATUS_PROCESS SET status = ? WHERE pid = ?",
		"killed", pidInt)

	if err != nil {
		log.Fatalf("Error al insertar proceso en db: %v", err)
	}

	log.Println("Process killed")
	return c.Status(200).JSON(fiber.Map{
		"message": "Proceso matado",
		"pid":     pid,
	})
}
