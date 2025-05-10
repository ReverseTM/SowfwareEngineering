package model

type Road struct {
	FromCity string `json:"from_city"`
	ToCity   string `json:"to_city"`
	Cost     int    `json:"cost"`
}
