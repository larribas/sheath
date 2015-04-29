package domain

import (
	"testing"
)

func TestNewLinkWithInvalidUrlReturnsError(t *testing.T) {
	_, err := NewLink("%1")
	if _, ok := err.(ErrURLNotValid); !ok {
		t.Error("Expected NewLink with invalid url to return ErrUrlNotValid error")
	}
}

func TestNewLinkGrabsAllUrlComponents(t *testing.T) {
	url := "http://user:pass@host.com:80/path?key=value#fragment"
	_, err := NewLink(url)
	if err != nil {
		t.Error("Expected NewLink with valid url not to return an error")
	}
}
