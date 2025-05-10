package storage

type EventType int

func (e EventType) String() string {
	switch e {
	case Insert:
		return "INSERT"
	case Update:
		return "UPDATE"
	case Delete:
		return "DELETE"
	default:
		return "UNKNOWN"
	}
}

const (
	Insert EventType = iota
	Update
	Delete
)

type Event struct {
	Type     EventType
	Table    string
	OldValue any
	NewValue any
}
