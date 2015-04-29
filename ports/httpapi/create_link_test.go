package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLink(t *testing.T) {
	port := NewPort(newApp())

	ts := httptest.NewServer(port.router)
	defer ts.Close()

	res, err := http.PostForm(ts.URL, map[string][]string{"url": []string{"http://valid.url"}})
	if err != nil {
		t.Errorf("Expected POST / not to return an error. Instead, it returned '%s'", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Error("Expected POST / to return a 201 CREATED status")
	}

	defer res.Body.Close()
	if n, _ := res.Body.Read(make([]byte, 10)); n == 0 {
		t.Error("Expected POST / to respond with a link stub")
	}
}

func TestCreateLinkMissingUrl(t *testing.T) {
	port := NewPort(newApp())

	ts := httptest.NewServer(port.router)
	defer ts.Close()

	res, err := http.PostForm(ts.URL, nil)
	if err != nil {
		t.Errorf("Expected POST / not to return an error. Instead, it returned '%s'", err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Error("Expected POST / with missing URL to return a 400 BAD REQUEST status")
	}

	defer res.Body.Close()
	if n, _ := res.Body.Read(make([]byte, 10)); n == 0 {
		t.Error("Expected POST / with missing URL to respond with an error message")
	}
}
