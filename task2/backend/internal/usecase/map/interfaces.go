package _map

import (
	"software-engineering-2/internal/delivery/map/common"
	"software-engineering-2/internal/model"
)

type UseCase interface {
	GetAllMapNames() ([]string, error)
	GetAllCities(mapName string) ([]*model.City, error)
	GetAllRoads(mapName string) ([]*model.Road, error)

	AddMap(mapName string) error
	AddCity(mapName string, request common.CityCreateRequest) error
	AddRoad(mapName string, request common.RoadCreateRequest) error

	UpdateCityName(mapName string, oldCityName, newCityName string) error
	UpdateRoadCost(mapName string, request common.RoadUpdateRequest) error

	DeleteMap(mapName string) error
	DeleteCity(mapName, cityName string) error
	DeleteRoad(mapName string, request common.RoadDeleteRequest) error

	Undo(mapName string) error
	Redo(mapName string) error

	Download(mapName string) (*MapData, error)
	Upload(mapData *MapData) error
}
