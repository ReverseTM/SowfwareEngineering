package observer

import "software-engineering/internal/storage"

type Observer interface {
	Notify(event storage.Event)
}
