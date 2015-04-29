// Package collection provides a way to submit a link through multiple validators and know which of them passed
package collection

import (
	"bytes"

	"fmt"
	"github.com/larribas/sheath/application"
	"github.com/larribas/sheath/application/domain"
)

// Collection maps validator names to particular validators with the purpose of executing them altogether
// and grouping all the errors returned by individual validators
type Collection map[string]application.Validator

// New creates a new Collection from the supplied map of validator names onto individual validator instances
func New(validators map[string]application.Validator) Collection {
	return Collection(validators)
}

// Validate submits the link to a validation process by all the individual validators that compose the collection,
// and groups all the errors they return, if any (nil otherwise)
func (coll Collection) Validate(link *domain.Link) error {
	errors := make(ErrFailedValidationsWithinCollection)

	for name, validator := range coll {
		err := validator.Validate(link)
		if err != nil {
			errors[name] = err
		}
	}

	if len(errors) == 0 {
		return nil
	}

	return errors
}

// ErrFailedValidationsWithinCollection contains a map of validator names and errors
// they have returned on the collection's Validate execution
type ErrFailedValidationsWithinCollection map[string]error

func (errors ErrFailedValidationsWithinCollection) Error() string {
	if errors == nil {
		return ""
	}

	var buffer bytes.Buffer
	buffer.WriteString("The following errors occured during validation:\n")

	for name, err := range errors {
		buffer.WriteString(fmt.Sprintf("[%s] %s\n", name, err.Error()))
	}

	return buffer.String()
}
