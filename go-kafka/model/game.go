package model

type Game struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Genre    string  `json:"genre"`
	Platform string  `json:"platform"`
}
