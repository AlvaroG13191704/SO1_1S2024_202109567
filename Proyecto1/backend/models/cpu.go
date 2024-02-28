package models

type Cpu struct {
	TotalCPU   string    `json:"totalCPU"`
	PercentCPU string    `json:"percentCPU"`
	Date       *string   `json:"date"`
	Processes  []Process `json:"processes"`
}

type Process struct {
	Pid      string         `json:"pid"`
	Name     string         `json:"name"`
	User     string         `json:"user"`
	Children []ProcessChild `json:"children"`
}

type ProcessChild struct {
	Pid       string `json:"pid"`
	Name      string `json:"name"`
	PidFather string `json:"pidFather"`
}
