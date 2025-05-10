package commands

import "software-engineering-2/internal/model"

type MapMemento struct {
	RoadMap *model.RoadMap `json:"state"`
}

func NewMapMemento(m *model.RoadMap) *MapMemento {
	copiedMap := &model.RoadMap{
		Name:   m.Name,
		Cities: make(map[string]*model.City),
		Roads:  make([]*model.Road, len(m.Roads)),
	}

	for name, city := range m.Cities {
		copiedMap.Cities[name] = &model.City{
			Name: city.Name,
			X:    city.X,
			Y:    city.Y,
		}
	}

	copy(copiedMap.Roads, m.Roads)

	return &MapMemento{RoadMap: copiedMap}
}

func (m *MapMemento) GetState() *model.RoadMap {
	return m.RoadMap
}
