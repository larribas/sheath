package null

import (
	"github.com/MedBrain/sheath/application/domain"
)

// Notifier represents a notifier that does nothing with the events it receives
type Notifier struct{}

// NotifyEvent ignores the domain events it received
func (n Notifier) NotifyEvent(event *domain.Event) {}
