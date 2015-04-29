package domain

import (
	"fmt"
	"net/url"

	"github.com/dchest/uniuri"
)

// A Link entity represents an (stub, originalURL) pair, identified by stub
type Link struct {
	Stub     string
	Original *url.URL
}

// NewLink creates a Link from a URL string. It generates the Link's unique identity.
// If the URL string is not valid, it returns an error
func NewLink(rawURL string) (*Link, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return &Link{}, ErrURLNotValid(rawURL)
	}

	return &Link{Original: u, Stub: uniuri.New()}, nil
}

// ErrURLNotValid is returned whenever a URL string has an invalid format
type ErrURLNotValid string

func (u ErrURLNotValid) Error() string {
	return fmt.Sprintf("The url '%s' has an invalid format", u)
}
