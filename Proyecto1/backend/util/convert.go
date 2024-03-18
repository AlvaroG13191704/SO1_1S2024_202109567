package util

import (
	"fmt"
	"sopes1/proyecto1/models"
	"time"
)

func ConvertRamFromKernelToRam(ramFromKernel models.RamFromKernel) models.Ram {
	ram := models.Ram{
		TotalRam: fmt.Sprintf("%.2f", (float64(ramFromKernel.TotalRam) / (1024 * 1024 * 1024))),
		FreeRam:  fmt.Sprintf("%.2f", (float64(ramFromKernel.FreeRam) / (1024 * 1024 * 1024))),
		UsedRam:  fmt.Sprintf("%.2f", (float64(ramFromKernel.UsedRam) / (1024 * 1024 * 1024))),
		Percent:  fmt.Sprintf("%.2f", float64(ramFromKernel.PercentUsed)),
		Date:     time.Now().Format("2006-01-02 15:04:05"),
	}
	return ram
}

func ConvertCpuFromKernelToCpu(cpuFromKernel models.CpuFromKernel) models.Cpu {
	var iterateValueUntil float64 = float64(cpuFromKernel.PercentCPU)
	var percentCPU string

	// Iterate until the value is less than or equal to 1
	for iterateValueUntil > 1 {
		if iterateValueUntil < 1000 {
			iterateValueUntil /= 10
			break
		} else {
			iterateValueUntil /= 100
		}
	}

	percentCPU = fmt.Sprintf("%.2f", iterateValueUntil)

	cpu := models.Cpu{
		TotalCPU:   fmt.Sprintf("%d", cpuFromKernel.TotalCPU),
		PercentCPU: percentCPU,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		Processes:  cpuFromKernel.Processes,
	}
	return cpu
}
