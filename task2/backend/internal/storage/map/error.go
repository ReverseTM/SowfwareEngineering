package _map

import "errors"

var MapNotFoundError = errors.New("map not found")

var MapAlreadyExistsError = errors.New("map already exists")

var CityNotFoundError = errors.New("city not found")

var CityAlreadyExistsError = errors.New("city already exists")

var CityWithSameNameAlreadyExistsError = errors.New("city with same name already exists")

var RoadAlreadyExistsError = errors.New("road already exists")

var RoadNotFoundError = errors.New("road not found")
