// Package application provides a way to create and interact with a Sheath instance.
// It is also responsible for defining relevant application-level interfaces
package application

import (
	"github.com/larribas/sheath/application/domain"
)

// App defines the use cases available to a Sheath application, along with the topology
// of a certain instance (the implementation of choice for each defined interface)
type App struct {
	LinkRepository domain.LinkRepository
	Notifier       Notifier
	Validator      Validator
}

// NewApp instantiates a Sheath application, receiving the particular implementations that will run for each interface
func NewApp(l domain.LinkRepository, n Notifier, v Validator) *App {
	return &App{
		LinkRepository: l,
		Notifier:       n,
		Validator:      v,
	}
}
