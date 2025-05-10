package _map

import "software-engineering-2/internal/model"

type Storage interface {
	GetMapByName(mapName string) (*model.RoadMap, error)
	GetAllMapNames() ([]string, error)
	GetAllCities(mapName string) ([]*model.City, error)
	GetAllRoads(mapName string) ([]*model.Road, error)

	AddMap(m *model.RoadMap) error
	AddCity(mapName string, city *model.City) error
	AddRoad(mapName string, road *model.Road) error

	UpdateMap(mapName string, roadMap *model.RoadMap) error
	UpdateCityName(mapName string, oldCityName, newCityName string) error
	UpdateRoadCost(mapName string, fromCity, toCity string, cost int) error

	DeleteMap(mapName string) error
	DeleteCity(mapName string, cityName string) error
	DeleteRoad(mapName string, fromCity, toCity string) error
}
