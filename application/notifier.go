package application

import (
	"github.com/larribas/sheath/application/domain"
)

// A Notifier is an output port to notify external applications of event domains that sprung from Sheath.
// It allows for the creation of a series of applications that use Sheath's usage information
// (analytical back offices, dashboards, etc.)
type Notifier interface {
	NotifyEvent(event *domain.Event)
}
