package main

import (
	"log"
	"sopes1/proyecto1/db"
	"sopes1/proyecto1/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// open db
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		panic(err)
	}

	// create tables
	tables := []string{
		`CREATE TABLE IF NOT EXISTS RAM (
			id INT AUTO_INCREMENT PRIMARY KEY,
			total VARCHAR(10) NOT NULL,
			free VARCHAR(10) NOT NULL,
			used_ram VARCHAR(10) NOT NULL,
			percentage_use VARCHAR(10) NOT NULL,
			date_time DATETIME NOT NULL
	)`,
		`CREATE TABLE IF NOT EXISTS CPU (
			id INT AUTO_INCREMENT PRIMARY KEY,
			total_cpu VARCHAR(10) NOT NULL,
			percentage_use VARCHAR(10) NOT NULL,
			date_time DATETIME NOT NULL
	)`,
		`CREATE TABLE IF NOT EXISTS STATUS_PROCESS (
			id INT AUTO_INCREMENT PRIMARY KEY,
			pid INT NOT NULL,
			status VARCHAR(20) NOT NULL,
			created_at DATETIME NOT NULL
	)`,
	}

	for _, table := range tables {
		_, err = dbClient.Exec(table)
		if err != nil {
			log.Fatalf("Failed to create table: %s\n", err)
			panic(err)
		}
	}
	log.Println("Tables created")

	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New())

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
