package general

import (
	"software-engineering-2/internal/model"
	_map "software-engineering-2/internal/storage/map"
)

type Storage struct {
	maps map[string]*model.RoadMap
}

func NewStorage() _map.Storage {
	return &Storage{
		maps: make(map[string]*model.RoadMap),
	}
}

func (s *Storage) GetMapByName(mapName string) (*model.RoadMap, error) {
	m, ok := s.maps[mapName]
	if !ok {
		return nil, _map.MapNotFoundError
	}

	return m, nil
}

func (s *Storage) UpdateMap(mapName string, roadMap *model.RoadMap) error {
	if _, ok := s.maps[mapName]; !ok {
		return _map.MapNotFoundError
	}

	s.maps[mapName] = roadMap

	return nil
}

func (s *Storage) GetAllMapNames() ([]string, error) {
	mapNames := make([]string, 0, len(s.maps))

	for name, _ := range s.maps {
		mapNames = append(mapNames, name)
	}

	return mapNames, nil
}

func (s *Storage) GetAllCities(mapName string) ([]*model.City, error) {
	m, ok := s.maps[mapName]
	if !ok {
		return nil, _map.MapAlreadyExistsError
	}

	cities := make([]*model.City, 0, len(m.Cities))
	for _, city := range m.Cities {
		cities = append(cities, city)
	}

	return cities, nil
}

func (s *Storage) GetAllRoads(mapName string) ([]*model.Road, error) {
	m, ok := s.maps[mapName]
	if !ok {
		return nil, _map.MapAlreadyExistsError
	}

	return m.Roads, nil
}

func (s *Storage) AddMap(m *model.RoadMap) error {
	if _, ok := s.maps[m.Name]; ok {
		return _map.MapAlreadyExistsError
	}

	s.maps[m.Name] = m

	return nil
}

func (s *Storage) AddCity(mapName string, city *model.City) error {
	m, ok := s.maps[mapName]
	if !ok {
		return _map.MapNotFoundError
	}

	if _, ok = m.Cities[city.Name]; ok {
		return _map.CityAlreadyExistsError
	}

	m.Cities[city.Name] = city

	return nil
}

func (s *Storage) AddRoad(mapName string, road *model.Road) error {
	m, ok := s.maps[mapName]
	if !ok {
		return _map.MapNotFoundError
	}

	if _, ok = m.Cities[road.FromCity]; !ok {
		return _map.CityNotFoundError
	}

	if _, ok = m.Cities[road.ToCity]; !ok {
		return _map.CityNotFoundError
	}

	index := -1
	for i, r := range m.Roads {
		if (r.FromCity == road.FromCity && r.ToCity == road.ToCity) ||
			(r.FromCity == road.ToCity && r.ToCity == road.FromCity) {
			index = i
			break
		}
	}

	if index != -1 {
		return _map.RoadAlreadyExistsError
	}

	m.Roads = append(m.Roads, road)

	return nil
}

func (s *Storage) UpdateCityName(mapName string, oldCityName, newCityName string) error {
	m, ok := s.maps[mapName]
	if !ok {
		return _map.MapNotFoundError
	}

	city, ok := m.Cities[oldCityName]
	if !ok {
		return _map.CityNotFoundError
	}

	if _, ok = m.Cities[newCityName]; ok {
		return _map.CityWithSameNameAlreadyExistsError
	}

	delete(m.Cities, oldCityName)

	city.Name = newCityName
	m.Cities[newCityName] = city

	return nil
}

func (s *Storage) UpdateRoadCost(mapName string, fromCity, toCity string, cost int) error {
	m, ok := s.maps[mapName]
	if !ok {
		return _map.MapNotFoundError
	}

	var road *model.Road
	for _, r := range m.Roads {
		if (r.FromCity == fromCity && r.ToCity == toCity) ||
			(r.FromCity == toCity && r.ToCity == fromCity) {
			road = r
			break
		}
	}

	if road == nil {
		return _map.RoadNotFoundError
	}

	road.Cost = cost

	return nil
}

func (s *Storage) DeleteMap(mapName string) error {
	if _, ok := s.maps[mapName]; !ok {
		return _map.MapNotFoundError
	}

	delete(s.maps, mapName)

	return nil
}

func (s *Storage) DeleteCity(mapName string, cityName string) error {
	m, ok := s.maps[mapName]
	if !ok {
		return _map.MapNotFoundError
	}

	if _, ok = m.Cities[cityName]; !ok {
		return _map.CityNotFoundError
	}

	roads := make([]*model.Road, 0)
	for _, r := range m.Roads {
		if r.FromCity != cityName && r.ToCity != cityName {
			roads = append(roads, r)
		}
	}

	delete(m.Cities, cityName)
	m.Roads = roads

	return nil
}

func (s *Storage) DeleteRoad(mapName string, fromCity, toCity string) error {
	m, ok := s.maps[mapName]
	if !ok {
		return _map.MapNotFoundError
	}

	index := -1
	for i, r := range m.Roads {
		if (r.FromCity == fromCity && r.ToCity == toCity) ||
			(r.FromCity == toCity && r.ToCity == fromCity) {
			index = i
			break
		}
	}

	if index == -1 {
		return _map.RoadNotFoundError
	}

	m.Roads = append(m.Roads[:index], m.Roads[index+1:]...)

	return nil
}
