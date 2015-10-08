package collection

import (
	"errors"
	"testing"

	"github.com/MedBrain/sheath/application"
	"github.com/MedBrain/sheath/application/domain"
)

type DumbValidator int

func (v DumbValidator) Validate(link *domain.Link) error {
	if link.Original.String() == "http://valid.url" {
		return nil
	}

	return errors.New("Invalid URL")
}

var collection = New(
	map[string]application.Validator{
		"dumb1": DumbValidator(1),
		"dumb2": DumbValidator(2),
	},
)

func TestCollectionReturnsNilErrorOnAValidLink(t *testing.T) {
	link, _ := domain.NewLink("http://valid.url")

	err := collection.Validate(link)
	if err != nil {
		t.Error("Expected valid link to return nil error when going through the whole collection")
	}
}

func TestCollectionReturnsMultipleErrors(t *testing.T) {
	link, _ := domain.NewLink("http://invalid.url")

	err := collection.Validate(link)
	if errors, ok := err.(ErrFailedValidationsWithinCollection); !ok {
		t.Error("Expected collection to return ErrFailedValidationsWithinCollection")
	} else if len(errors) != 2 {
		t.Error("Expected collection to register 2 errors")
	}
}
