package models

type Log struct {
	Value     string
	CreatedAt string
}

type Data struct {
	Name  string `json:"name"`
	Album string `json:"album"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}
