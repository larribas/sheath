package protocol

import (
	"testing"

	"github.com/MedBrain/sheath/application/domain"
)

func TestValidateEmptyScheme(t *testing.T) {
	vEmpty := NewValidator("")
	vHTTP := NewValidator("http")
	link, _ := domain.NewLink("valid.url")

	if err := vEmpty.Validate(link); err != nil {
		t.Error("Expected scheme validator (for \"\") to validate an empty scheme")
	}

	if _, ok := vHTTP.Validate(link).(ErrProtocolIsNotAllowed); !ok {
		t.Error("Expected scheme validator (for http) NOT to validate an empty scheme")
	}
}

func TestValidateScheme(t *testing.T) {
	v := NewValidator("http", "https")
	linkHTTP, _ := domain.NewLink("http://valid.url")
	linkFoo, _ := domain.NewLink("foo://invalid.url")

	if err := v.Validate(linkHTTP); err != nil {
		t.Error("Expected scheme validator (for http, https) to validate an http scheme")
	}

	if _, ok := v.Validate(linkFoo).(ErrProtocolIsNotAllowed); !ok {
		t.Error("Expected scheme validator (for http, https) NOT to validate a foo scheme")
	}
}
