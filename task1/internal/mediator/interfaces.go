package mediator

import (
	"software-engineering/internal/observer"
	"software-engineering/internal/storage"
)

type Mediator interface {
	Subscribe(eventType storage.EventType, observer observer.Observer)
	Unsubscribe(eventType storage.EventType, observer observer.Observer)
	Publish(event storage.Event)
}
