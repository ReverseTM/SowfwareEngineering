package common

type CityCreateRequest struct {
	CityName string `json:"city_name"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

type RoadCreateRequest struct {
	FromCity string `json:"from_city"`
	ToCity   string `json:"to_city"`
	Cost     int    `json:"cost"`
}

type RoadUpdateRequest struct {
	FromCity string `json:"from_city"`
	ToCity   string `json:"to_city"`
	Cost     int    `json:"cost"`
}

type RoadDeleteRequest struct {
	FromCity string `json:"from_city"`
	ToCity   string `json:"to_city"`
}
