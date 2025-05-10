package commands

import (
	"software-engineering-2/internal/model"
	_map "software-engineering-2/internal/storage/map"
	"software-engineering-2/internal/usecase/map/general/command"
)

const (
	AddCity    = "ADD-CITY"
	UpdateCity = "UPDATE-CITY"
	DeleteCity = "DELETE-CITY"
	AddRoad    = "ADD-ROAD"
	UpdateRoad = "UPDATE-ROAD"
	DeleteRoad = "DELETE-ROAD"
)

type AddCityCommand struct {
	Action   string      `json:"action"`
	MapName  string      `json:"map_name"`
	CityName string      `json:"city_name"`
	X        int         `json:"x"`
	Y        int         `json:"y"`
	Memento  *MapMemento `json:"memento"`
	storage  _map.Storage
}

func NewAddCityCommand(mapName, cityName string, x, y int, storage _map.Storage) command.Command {
	return &AddCityCommand{
		Action:   AddCity,
		MapName:  mapName,
		CityName: cityName,
		X:        x,
		Y:        y,
		storage:  storage,
	}
}

func (c *AddCityCommand) Execute() error {
	m, err := c.storage.GetMapByName(c.MapName)
	if err != nil {
		return err
	}

	c.Memento = NewMapMemento(m)

	city := &model.City{
		Name: c.CityName,
		X:    c.X,
		Y:    c.Y,
	}

	return c.storage.AddCity(c.MapName, city)
}

func (c *AddCityCommand) Undo() error {
	if c.Memento == nil {
		return nil
	}

	if err := c.storage.UpdateMap(c.MapName, c.Memento.GetState()); err != nil {
		return err
	}

	return nil
}

func (c *AddCityCommand) SetStorage(s _map.Storage) {
	c.storage = s
}

type UpdateCityCommand struct {
	Action      string      `json:"action"`
	MapName     string      `json:"map_name"`
	OldCityName string      `json:"old_city_name"`
	NewCityName string      `json:"new_city_name"`
	Memento     *MapMemento `json:"memento"`
	storage     _map.Storage
}

func NewUpdateCityCommand(mapName, oldCityName, newCityName string, storage _map.Storage) command.Command {
	return &UpdateCityCommand{
		Action:      UpdateCity,
		MapName:     mapName,
		OldCityName: oldCityName,
		NewCityName: newCityName,
		storage:     storage,
	}
}

func (c *UpdateCityCommand) Execute() error {
	m, err := c.storage.GetMapByName(c.MapName)
	if err != nil {
		return err
	}

	c.Memento = NewMapMemento(m)

	return c.storage.UpdateCityName(c.MapName, c.OldCityName, c.NewCityName)
}

func (c *UpdateCityCommand) Undo() error {
	if c.Memento == nil {
		return nil
	}

	if err := c.storage.UpdateMap(c.MapName, c.Memento.GetState()); err != nil {
		return err
	}

	return nil
}

func (c *UpdateCityCommand) SetStorage(s _map.Storage) {
	c.storage = s
}

type DeleteCityCommand struct {
	Action   string      `json:"action"`
	MapName  string      `json:"map_name"`
	CityName string      `json:"city_name"`
	Memento  *MapMemento `json:"memento"`
	storage  _map.Storage
}

func NewDeleteCityCommand(mapName, cityName string, storage _map.Storage) command.Command {
	return &DeleteCityCommand{
		Action:   DeleteCity,
		MapName:  mapName,
		CityName: cityName,
		storage:  storage,
	}
}

func (c *DeleteCityCommand) Execute() error {
	m, err := c.storage.GetMapByName(c.MapName)
	if err != nil {
		return err
	}

	c.Memento = NewMapMemento(m)

	return c.storage.DeleteCity(c.MapName, c.CityName)
}

func (c *DeleteCityCommand) Undo() error {
	if c.Memento == nil {
		return nil
	}

	if err := c.storage.UpdateMap(c.MapName, c.Memento.GetState()); err != nil {
		return err
	}

	return nil
}

func (c *DeleteCityCommand) SetStorage(s _map.Storage) {
	c.storage = s
}

type AddRoadCommand struct {
	Action   string      `json:"action"`
	MapName  string      `json:"map_name"`
	FromCity string      `json:"from_city"`
	ToCity   string      `json:"to_city"`
	Cost     int         `json:"cost"`
	Memento  *MapMemento `json:"memento"`
	storage  _map.Storage
}

func NewAddRoadCommand(mapName, fromCity, toCity string, cost int, storage _map.Storage) command.Command {
	return &AddRoadCommand{
		Action:   AddRoad,
		MapName:  mapName,
		FromCity: fromCity,
		ToCity:   toCity,
		Cost:     cost,
		storage:  storage,
	}
}

func (c *AddRoadCommand) Execute() error {
	m, err := c.storage.GetMapByName(c.MapName)
	if err != nil {
		return err
	}

	c.Memento = NewMapMemento(m)

	road := &model.Road{
		FromCity: c.FromCity,
		ToCity:   c.ToCity,
		Cost:     c.Cost,
	}

	return c.storage.AddRoad(c.MapName, road)
}

func (c *AddRoadCommand) Undo() error {
	if c.Memento == nil {
		return nil
	}

	if err := c.storage.UpdateMap(c.MapName, c.Memento.GetState()); err != nil {
		return err
	}

	return nil
}

func (c *AddRoadCommand) SetStorage(s _map.Storage) {
	c.storage = s
}

type UpdateRoadCommand struct {
	Action   string      `json:"action"`
	MapName  string      `json:"map_name"`
	FromCity string      `json:"from_city"`
	ToCity   string      `json:"to_city"`
	Cost     int         `json:"cost"`
	Memento  *MapMemento `json:"memento"`
	storage  _map.Storage
}

func NewUpdateRoadCommand(mapName, fromCity, toCity string, cost int, storage _map.Storage) command.Command {
	return &UpdateRoadCommand{
		Action:   UpdateRoad,
		MapName:  mapName,
		FromCity: fromCity,
		ToCity:   toCity,
		Cost:     cost,
		storage:  storage,
	}
}

func (c *UpdateRoadCommand) Execute() error {
	m, err := c.storage.GetMapByName(c.MapName)
	if err != nil {
		return err
	}

	c.Memento = NewMapMemento(m)

	return c.storage.UpdateRoadCost(c.MapName, c.FromCity, c.ToCity, c.Cost)
}

func (c *UpdateRoadCommand) Undo() error {
	if c.Memento == nil {
		return nil
	}

	if err := c.storage.UpdateMap(c.MapName, c.Memento.GetState()); err != nil {
		return err
	}

	return nil
}

func (c *UpdateRoadCommand) SetStorage(s _map.Storage) {
	c.storage = s
}

type DeleteRoadCommand struct {
	Action   string      `json:"action"`
	MapName  string      `json:"map_name"`
	FromCity string      `json:"from_city"`
	ToCity   string      `json:"to_city"`
	Memento  *MapMemento `json:"memento"`
	storage  _map.Storage
}

func NewDeleteRoadCommand(mapName, fromCity, toCity string, storage _map.Storage) command.Command {
	return &DeleteRoadCommand{
		Action:   DeleteRoad,
		MapName:  mapName,
		FromCity: fromCity,
		ToCity:   toCity,
		storage:  storage,
	}
}

func (c *DeleteRoadCommand) Execute() error {
	m, err := c.storage.GetMapByName(c.MapName)
	if err != nil {
		return err
	}

	c.Memento = NewMapMemento(m)

	return c.storage.DeleteRoad(c.MapName, c.FromCity, c.ToCity)
}

func (c *DeleteRoadCommand) Undo() error {
	if c.Memento == nil {
		return nil
	}

	if err := c.storage.UpdateMap(c.MapName, c.Memento.GetState()); err != nil {
		return err
	}

	return nil
}

func (c *DeleteRoadCommand) SetStorage(s _map.Storage) {
	c.storage = s
}
