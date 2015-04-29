package httpapi

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetrieveLink(t *testing.T) {
	app := newApp()
	// Assume a pre-existing link
	originalURL := "http://www.some-url.com"
	req, _ := http.NewRequest("POST", originalURL, nil)
	stub, _ := app.CreateLink(originalURL, req)
	port := NewPort(app)

	ts := httptest.NewServer(port.router)
	defer ts.Close()

	// Modify the default client so that it doesn't follow redirects
	http.DefaultClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("")
		},
	}
	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, stub))

	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("Expected GET /:stub to return a 307 TEMPORARY REDIRECT status, %d returned instead", res.StatusCode)
	}

	if res.Header.Get("Location") != originalURL {
		t.Errorf("Expected GET /:stub to respond with an appropriate Location header")
	}
}

func TestRetrieveInexistentLink(t *testing.T) {
	port := NewPort(newApp())

	ts := httptest.NewServer(port.router)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/%s", ts.URL, "inexistent-link"))
	if err != nil {
		t.Errorf("Expected GET /:stub not to return an error. Instead, it returned '%s'", err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Error("Expected GET /:stub to return a 404 NOT FOUND status")
	}
}
