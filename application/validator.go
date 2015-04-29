package application

import (
	"github.com/larribas/sheath/application/domain"
)

// A Validator is a way to enforce a certain policy on the kind of URLs the application handles.
// It may be used to allow some protocols, blacklist certain hosts, etc.
type Validator interface {
	Validate(*domain.Link) error
}
