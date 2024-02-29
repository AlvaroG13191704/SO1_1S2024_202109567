package models

type RamFromKernel struct {
	TotalRam    int64 `json:"totalRam"`
	FreeRam     int64 `json:"freeRam"`
	UsedRam     int64 `json:"usedRam"`
	PercentUsed int64 `json:"percentUsed"`
	// Date        *string `json:"date"`
}

type Ram struct {
	TotalRam string `json:"totalRam"`
	FreeRam  string `json:"freeRam"`
	UsedRam  string `json:"usedRam"`
	Percent  string `json:"percent"`
	Date     string `json:"date"`
}

type HistoryRam struct {
	Date    string `json:"date"`
	Percent string `json:"percent"`
}
