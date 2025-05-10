package logger

import (
	"fmt"
	"software-engineering/internal/observer"
	"software-engineering/internal/storage"
)

var _ observer.Observer = (*AuditLogger)(nil)

const messageTemplate = "Audit event received: type=%q, table=%q, old_value=%+v, new_value=%+v\n"

type AuditLogger struct{}

func NewAuditLogger() observer.Observer {
	return &AuditLogger{}
}

func (l *AuditLogger) Notify(event storage.Event) {
	fmt.Printf(
		messageTemplate,
		event.Type,
		event.Table,
		event.OldValue,
		event.NewValue,
	)
}
