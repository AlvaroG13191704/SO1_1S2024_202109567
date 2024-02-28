package models

type Ram struct {
	TotalRam    string  `json:"totalRam"`
	FreeRam     string  `json:"freeRam"`
	UsedRam     string  `json:"usedRam"`
	PercentUsed string  `json:"percentUsed"`
	Date        *string `json:"date"`
}
