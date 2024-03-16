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

	cpu := models.Cpu{
		TotalCPU:   fmt.Sprintf("%d", cpuFromKernel.TotalCPU),
		PercentCPU: fmt.Sprintf("%.2f", float64(cpuFromKernel.PercentCPU)/1000000),
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		Processes:  cpuFromKernel.Processes,
	}
	return cpu
}
