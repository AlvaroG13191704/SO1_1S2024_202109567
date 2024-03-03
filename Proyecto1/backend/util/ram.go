package util

import (
	"encoding/json"
	"log"
	"os/exec"
	"sopes1/proyecto1/db"
	"sopes1/proyecto1/models"
)

func GetRAM() (models.Ram, error) {
	cmd := exec.Command("sudo", "cat", "/proc/ram_so1_1s2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// marshal to Ram struct
	var cmdRam models.RamFromKernel
	err = json.Unmarshal(out, &cmdRam)
	if err != nil {
		log.Fatalf("json.Unmarshal() failed with %s\n", err)
	}

	var ram models.Ram = ConvertRamFromKernelToRam(cmdRam)

	// save to db
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return ram, err
	}
	// defer dbClient.Close()

	// insert to db
	_, err = dbClient.Exec(
		"INSERT INTO RAM (total, free, used_ram, percentage_use, date_time) VALUES (?, ?, ?, ?, ?)",
		ram.TotalRam, ram.FreeRam, ram.UsedRam, ram.Percent, ram.Date)

	// log.Println("ram - ", ram)
	return ram, err
}

func GetHistoryRam() ([]models.HistoryRam, error) {
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return []models.HistoryRam{}, err
	}

	rows, err := dbClient.Query("SELECT percentage_use, date_time FROM RAM")
	if err != nil {
		log.Fatalf("db.Query() failed with %s\n", err)
		return []models.HistoryRam{}, err
	}
	defer rows.Close()

	var historyRam []models.HistoryRam
	for rows.Next() {
		var hr models.HistoryRam
		err = rows.Scan(&hr.Percent, &hr.Date)
		if err != nil {
			log.Fatalf("rows.Scan() failed with %s\n", err)
			return []models.HistoryRam{}, err
		}
		historyRam = append(historyRam, hr)
	}

	return historyRam, nil
}
