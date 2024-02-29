package models

type CpuFromKernel struct {
	TotalCPU   int64 `json:"totalCPU"`
	PercentCPU int64 `json:"percentCPU"`
	// Date       *string   `json:"date"`
	Processes []Process `json:"processes"`
}

type Process struct {
	Pid      int64          `json:"pid"`
	Name     string         `json:"name"`
	User     int64          `json:"user"`
	Children []ProcessChild `json:"children"`
}

type ProcessChild struct {
	Pid       int64  `json:"pid"`
	Name      string `json:"name"`
	PidFather int64  `json:"pidFather"`
}

type Cpu struct {
	TotalCPU   string    `json:"totalCPU"`
	PercentCPU string    `json:"percentCPU"`
	Date       string    `json:"date"`
	Processes  []Process `json:"processes"`
}

type HistoryCpu struct {
	Date       string `json:"date"`
	Percentage string `json:"percentage"`
}
