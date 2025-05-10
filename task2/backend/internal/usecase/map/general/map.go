package general

import (
	"fmt"
	"software-engineering-2/internal/delivery/map/common"
	"software-engineering-2/internal/model"
	_map "software-engineering-2/internal/storage/map"
	_map2 "software-engineering-2/internal/usecase/map"
	"software-engineering-2/internal/usecase/map/general/command"
	"software-engineering-2/internal/usecase/map/general/command/commands"
)

type UseCase struct {
	storage      _map.Storage
	undoStackMap map[string][]command.Command
	redoStackMap map[string][]command.Command
}

func NewUseCase(storage _map.Storage) _map2.UseCase {
	return &UseCase{
		storage:      storage,
		undoStackMap: make(map[string][]command.Command),
		redoStackMap: make(map[string][]command.Command),
	}
}

func (uc *UseCase) GetAllMapNames() ([]string, error) {
	return uc.storage.GetAllMapNames()
}

func (uc *UseCase) GetAllCities(mapName string) ([]*model.City, error) {
	if mapName == "" {
		return nil, _map2.ErrInvalidMapName
	}

	return uc.storage.GetAllCities(mapName)
}

func (uc *UseCase) GetAllRoads(mapName string) ([]*model.Road, error) {
	if mapName == "" {
		return nil, _map2.ErrInvalidMapName
	}

	return uc.storage.GetAllRoads(mapName)
}

func (uc *UseCase) AddMap(mapName string) error {
	if mapName == "" {
		return _map2.ErrInvalidMapName
	}

	roadMap := &model.RoadMap{
		Name:   mapName,
		Cities: make(map[string]*model.City),
		Roads:  make([]*model.Road, 0),
	}

	return uc.storage.AddMap(roadMap)
}

func (uc *UseCase) AddCity(mapName string, request common.CityCreateRequest) error {
	if request.CityName == "" {
		return _map2.ErrInvalidCityName
	}

	if request.X < 1 || request.Y < 1 {
		return _map2.ErrInvalidCoordinates
	}

	cmd := commands.NewAddCityCommand(mapName, request.CityName, request.X, request.Y, uc.storage)
	uc.undoStackMap[mapName] = append(uc.undoStackMap[mapName], cmd)

	return cmd.Execute()
}

func (uc *UseCase) AddRoad(mapName string, request common.RoadCreateRequest) error {
	if request.FromCity == "" || request.ToCity == "" {
		return _map2.ErrInvalidCityName
	}

	if request.Cost < 0 {
		return _map2.ErrInvalidCost
	}

	cmd := commands.NewAddRoadCommand(mapName, request.FromCity, request.ToCity, request.Cost, uc.storage)
	uc.undoStackMap[mapName] = append(uc.undoStackMap[mapName], cmd)

	return cmd.Execute()
}

func (uc *UseCase) UpdateCityName(mapName string, oldCityName, newCityName string) error {
	if oldCityName == "" || newCityName == "" {
		return _map2.ErrInvalidCityName
	}

	cmd := commands.NewUpdateCityCommand(mapName, oldCityName, newCityName, uc.storage)
	uc.undoStackMap[mapName] = append(uc.undoStackMap[mapName], cmd)

	return cmd.Execute()
}

func (uc *UseCase) UpdateRoadCost(mapName string, request common.RoadUpdateRequest) error {
	if request.FromCity == "" || request.ToCity == "" {
		return _map2.ErrInvalidCityName
	}

	if request.Cost < 0 {
		return _map2.ErrInvalidCost
	}

	cmd := commands.NewUpdateRoadCommand(mapName, request.FromCity, request.ToCity, request.Cost, uc.storage)
	uc.undoStackMap[mapName] = append(uc.undoStackMap[mapName], cmd)

	return cmd.Execute()
}

func (uc *UseCase) DeleteMap(mapName string) error {
	if mapName == "" {
		return _map2.ErrInvalidMapName
	}

	if err := uc.storage.DeleteMap(mapName); err != nil {
		return err
	}

	fmt.Printf("%+v\n", uc.undoStackMap[mapName])

	delete(uc.undoStackMap, mapName)
	delete(uc.redoStackMap, mapName)

	return nil
}

func (uc *UseCase) DeleteCity(mapName, cityName string) error {
	if cityName == "" {
		return _map2.ErrInvalidCityName
	}

	cmd := commands.NewDeleteCityCommand(mapName, cityName, uc.storage)
	uc.undoStackMap[mapName] = append(uc.undoStackMap[mapName], cmd)

	return cmd.Execute()
}

func (uc *UseCase) DeleteRoad(mapName string, request common.RoadDeleteRequest) error {
	if request.FromCity == "" || request.ToCity == "" {
		return _map2.ErrInvalidCityName
	}

	cmd := commands.NewDeleteRoadCommand(mapName, request.FromCity, request.ToCity, uc.storage)
	uc.undoStackMap[mapName] = append(uc.undoStackMap[mapName], cmd)

	return cmd.Execute()
}

func (uc *UseCase) Undo(mapName string) error {
	if mapName == "" {
		return _map2.ErrInvalidMapName
	}

	undoStack := uc.undoStackMap[mapName]
	undoStackLen := len(undoStack)

	if undoStackLen == 0 {
		return nil
	}

	cmd := undoStack[undoStackLen-1]
	uc.undoStackMap[mapName] = undoStack[:undoStackLen-1]

	if err := cmd.Undo(); err != nil {
		return err
	}

	uc.redoStackMap[mapName] = append(uc.redoStackMap[mapName], cmd)

	return nil
}

func (uc *UseCase) Redo(mapName string) error {
	if mapName == "" {
		return _map2.ErrInvalidMapName
	}

	redoStack := uc.redoStackMap[mapName]
	redoStackLen := len(redoStack)

	if redoStackLen == 0 {
		return nil
	}

	cmd := redoStack[redoStackLen-1]
	uc.redoStackMap[mapName] = redoStack[:redoStackLen-1]

	return cmd.Execute()
}

func (uc *UseCase) Download(mapName string) (*_map2.MapData, error) {
	if mapName == "" {
		return nil, _map2.ErrInvalidMapName
	}

	m, err := uc.storage.GetMapByName(mapName)
	if err != nil {
		return nil, err
	}

	undoStack := uc.undoStackMap[mapName]
	redoStack := uc.redoStackMap[mapName]

	history := &_map2.History{
		UndoCommands: make([]command.Command, 0, len(undoStack)),
		RedoCommands: make([]command.Command, 0, len(redoStack)),
	}

	for _, cmd := range undoStack {
		history.UndoCommands = append(history.UndoCommands, cmd)
	}

	for _, cmd := range redoStack {
		history.RedoCommands = append(history.RedoCommands, cmd)
	}

	mapData := &_map2.MapData{
		Map:     m,
		History: history,
	}

	return mapData, nil
}

func (uc *UseCase) Upload(mapData *_map2.MapData) error {
	if mapData.Map == nil {
		return _map2.ErrNoMapToAdd
	}

	if err := uc.storage.AddMap(mapData.Map); err != nil {
		return err
	}

	if mapData.History != nil {
		uc.undoStackMap[mapData.Map.Name] = mapData.History.UndoCommands
		for _, cmd := range uc.undoStackMap[mapData.Map.Name] {
			cmd.SetStorage(uc.storage)
		}

		uc.redoStackMap[mapData.Map.Name] = mapData.History.RedoCommands
		for _, cmd := range uc.redoStackMap[mapData.Map.Name] {
			cmd.SetStorage(uc.storage)
		}
	}

	return nil
}
