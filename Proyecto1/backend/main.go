package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"sopes1/proyecto1/db"
	"sopes1/proyecto1/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/real-time", func(c *fiber.Ctx) error {

		ram, err := getRAM()
		if err != nil {
			log.Fatalf("getRAM() failed with %s\n", err)
			return c.JSON(fiber.Map{
				"error": "getRAM() failed",
			})
		}

		cpu, err := getCPU()
		if err != nil {
			log.Fatalf("getCPU() failed with %s\n", err)
			return c.JSON(fiber.Map{
				"error": "getCPU() failed",
			})
		}

		// read from kernel
		return c.JSON(fiber.Map{
			"ram": ram,
			"cpu": cpu,
		})
	})

	app.Listen(":8080")
}

func getCPU() (models.Cpu, error) {
	cmd := exec.Command("sudo", "cat", "/proc/cpu_so1_1s2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// marshal to Ram struct
	var cpu models.Cpu
	err = json.Unmarshal(out, &cpu)
	if err != nil {
		log.Fatalf("json.Unmarshal() failed with %s\n", err)
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	cpu.Date = &date

	// save to db
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return cpu, err
	}

	// insert to db
	_, err = dbClient.Exec(
		"INSERT INTO CPU (total_cpu, percentage_use, date_time, processes) VALUES ($1, $2, $3, $4)",
		cpu.TotalCPU, cpu.PercentCPU, cpu.Date, cpu.Processes)

	return cpu, err

}

func getRAM() (models.Ram, error) {
	cmd := exec.Command("sudo", "cat", "/proc/ram_so1_1s2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// marshal to Ram struct
	var ram models.Ram
	err = json.Unmarshal(out, &ram)
	if err != nil {
		log.Fatalf("json.Unmarshal() failed with %s\n", err)
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	ram.Date = &date

	// save to db
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return ram, err
	}

	// insert to db
	_, err = dbClient.Exec(
		"INSERT INTO RAM (total, free, used_ram, percentage_use, date_time) VALUES ($1, $2, $3, $4, $5)",
		ram.TotalRam, ram.FreeRam, ram.UsedRam, ram.PercentUsed, ram.Date)

	return ram, err
}
