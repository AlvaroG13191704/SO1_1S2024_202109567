package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Ram struct {
	Total       string
	Free        string
	Used        string
	PorcentUsed string
}

func main() {
	CatProcFile("ram_202109567")
}

func CatProcFile(fileName string) {
	cmd := exec.Command("sudo", "cat", "/proc/"+fileName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// split the output by new line
	var currentRam Ram

	lines := strings.Split(string(out), "\n")
	for index, line := range lines {
		// split the line by :
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			// eliminate the spaces
			value := strings.TrimSpace(parts[1])

			// convert value to float32
			floatValue, err := strconv.ParseFloat(value, 32)
			if err != nil {
				log.Fatalf("Failed to convert string to float: %v", err)
			}

			if index == 0 {
				currentRam.Free = fmt.Sprintf("%.2f", (floatValue / (1024 * 1024)))

			} else if index == 1 {
				currentRam.Total = fmt.Sprintf("%.2f", (floatValue / (1024 * 1024)))

			} else if index == 2 {
				currentRam.Used = fmt.Sprintf("%.2f", (floatValue / (1024 * 1024)))

			} else if index == 3 {
				currentRam.PorcentUsed = fmt.Sprintf("%.2f", (floatValue))

			}
		}
	}

	fmt.Printf("Total: %v\n", currentRam)

}
