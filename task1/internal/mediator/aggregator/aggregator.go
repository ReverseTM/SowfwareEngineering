package aggregator

import (
	"slices"
	"software-engineering/internal/mediator"
	"software-engineering/internal/observer"
	"software-engineering/internal/storage"
)

var _ mediator.Mediator = (*EventAggregator)(nil)

type EventAggregator struct {
	observers map[storage.EventType][]observer.Observer
}

func NewEventAggregator() mediator.Mediator {
	return &EventAggregator{
		observers: make(map[storage.EventType][]observer.Observer),
	}
}

func (agg *EventAggregator) Subscribe(eventType storage.EventType, obs observer.Observer) {
	agg.observers[eventType] = append(agg.observers[eventType], obs)
}

func (agg *EventAggregator) Unsubscribe(eventType storage.EventType, obs observer.Observer) {
	if _, ok := agg.observers[eventType]; !ok {
		return
	}

	observers := agg.observers[eventType]
	agg.observers[eventType] = slices.DeleteFunc(observers, func(o observer.Observer) bool {
		return o == obs
	})
}

func (agg *EventAggregator) Publish(event storage.Event) {
	for _, obs := range agg.observers[event.Type] {
		obs.Notify(event)
	}
}
