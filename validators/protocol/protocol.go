// Package protocol provides a validator for protocol schemes
package protocol

import (
	"fmt"

	"github.com/MedBrain/sheath/application/domain"
)

// Validator contains a whitelist of allowed protocol schemes, and validates a Link based on it
type Validator struct {
	whitelist map[string]bool
}

// NewValidator initializes a Validator with a whitelist of protocol schemes.
// N.B: An empty scheme is disallowed by default. It represents a relative URL,
// so redirects will be relative to the site requesting such link, and thus are to be used carefully
func NewValidator(validSchemes ...string) *Validator {
	v := &Validator{whitelist: make(map[string]bool)}

	for _, scheme := range validSchemes {
		v.whitelist[scheme] = true
	}

	return v
}

// Validate checks if the link's protocol scheme is allowed, and returns an error otherwise
func (v Validator) Validate(link *domain.Link) error {
	scheme := link.Original.Scheme
	if _, found := v.whitelist[scheme]; !found {
		return ErrProtocolIsNotAllowed(scheme)
	}

	return nil
}

// ErrProtocolIsNotAllowed is returned when the protocol for a certain link is not allowed
type ErrProtocolIsNotAllowed string

func (err ErrProtocolIsNotAllowed) Error() string {
	return fmt.Sprintf("Links with a '%s' protocol scheme are not supported", err)
}
