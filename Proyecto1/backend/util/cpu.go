package util

import (
	"encoding/json"
	"log"
	"os/exec"
	"sopes1/proyecto1/db"
	"sopes1/proyecto1/models"
)

func GetCPU(saveOnDB bool) (models.Cpu, error) {
	cmd := exec.Command("sudo", "cat", "/proc/cpu_so1_1s2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// marshal to Ram struct
	var cpuKernel models.CpuFromKernel
	err = json.Unmarshal(out, &cpuKernel)
	if err != nil {
		log.Fatalf("json.Unmarshal() failed with %s\n", err)
	}

	var cpu = ConvertCpuFromKernelToCpu(cpuKernel)

	if saveOnDB {
		// save to db
		dbClient, err := db.GetDB()
		if err != nil {
			log.Fatalf("db.GetDB() failed with %s\n", err)
			return cpu, err
		}
		// defer dbClient.Close()

		processesJson, err := json.Marshal(cpu.Processes)
		if err != nil {
			log.Fatalf("Failed to marshal processes: %v", err)
		}

		// insert to db
		_, err = dbClient.Exec(
			"INSERT INTO CPU (total_cpu, percentage_use, date_time, processes) VALUES (?, ?, ?, ?)",
			cpu.TotalCPU, cpu.PercentCPU, cpu.Date, string(processesJson))

		return cpu, err
	}

	return cpu, err

}

func GetHistoryCPU() ([]models.HistoryCpu, error) {
	dbClient, err := db.GetDB()
	if err != nil {
		log.Fatalf("db.GetDB() failed with %s\n", err)
		return []models.HistoryCpu{}, err
	}

	rows, err := dbClient.Query("SELECT percentage_use, date_time FROM CPU")
	if err != nil {
		log.Fatalf("db.Query() failed with %s\n", err)
		return []models.HistoryCpu{}, err
	}
	defer rows.Close()

	var historyCpu []models.HistoryCpu

	for rows.Next() {
		var hc models.HistoryCpu
		err = rows.Scan(&hc.Percentage, &hc.Date)
		if err != nil {
			log.Fatalf("rows.Scan() failed with %s\n", err)
			return []models.HistoryCpu{}, err
		}
		historyCpu = append(historyCpu, hc)
	}

	return historyCpu, nil
}
