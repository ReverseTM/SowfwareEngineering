package _map

import "errors"

var ErrInvalidMapName = errors.New("invalid map name")

var ErrInvalidCityName = errors.New("invalid city name")

var ErrInvalidCoordinates = errors.New("invalid coordinates")

var ErrInvalidCost = errors.New("invalid cost")

var ErrNoMapToAdd = errors.New("no map to add")
