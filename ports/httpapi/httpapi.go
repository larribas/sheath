package httpapi

import (
	"fmt"
	"net/http"

	"github.com/MedBrain/sheath/application"

	"github.com/bmizerany/pat"
)

// Port configuration
const (
	address = ""
	port    = "1827"
)

// Port exposes an HTTP API to create links and retrieve them (being redirected).
// It's intended to be used from browsers and web application clients
type Port struct {
	router *pat.PatternServeMux
}

// NewPort receives a Sheath application instance, and returns a ready Port for such application
func NewPort(app *application.App) *Port {
	port := &Port{}
	port.router = pat.New()
	port.router.Get("/:link_stub", withApp(app, retrieveLink))
	port.router.Post("/", withApp(app, createLink))

	return port
}

// Expose starts the port by listening on the address and port specified, and directing such requests to the Port
func (p *Port) Expose() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), p.router)
}

type controllerFunc func(*application.App, http.ResponseWriter, *http.Request)

func withApp(app *application.App, f controllerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f(app, w, r)
	})
}
