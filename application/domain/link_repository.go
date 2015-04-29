package domain

import (
	"fmt"
)

// A LinkRepository is responsible for the storage and retrieval of links.
type LinkRepository interface {
	Store(*Link) error
	Find(string) (*Link, error)
}

// ErrLinkNotFound is returned when the repository does not contain a Link identified by a certain stub
type ErrLinkNotFound string

// ErrLinkAlreadyExists is returned when the repository tries to save a Link twice (i.e. two Links with the same stub)
type ErrLinkAlreadyExists string

func (stub ErrLinkNotFound) Error() string {
	return fmt.Sprintf("There is no link identified by stub '%s'", stub)
}

func (stub ErrLinkAlreadyExists) Error() string {
	return fmt.Sprintf("There is already a link identified by stub '%s'", stub)
}
