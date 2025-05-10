package command

import _map "software-engineering-2/internal/storage/map"

type Command interface {
	Execute() error
	Undo() error
	SetStorage(s _map.Storage)
}
