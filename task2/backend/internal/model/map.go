package model

type RoadMap struct {
	Name   string           `json:"name"`
	Cities map[string]*City `json:"cities"`
	Roads  []*Road          `json:"roads"`
}
