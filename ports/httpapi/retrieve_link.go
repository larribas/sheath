package httpapi

import (
	"net/http"

	"github.com/MedBrain/sheath/application"
	"github.com/MedBrain/sheath/application/domain"
)

func retrieveLink(app *application.App, w http.ResponseWriter, r *http.Request) {
	linkStub := r.URL.Query().Get(":link_stub")

	originalURL, err := app.RetrieveLink(linkStub, r)
	if err != nil {
		switch err.(type) {
		case domain.ErrLinkNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	// TODO Compare with StatusMovedPermanently, StatusFound, and StatusSeeOther
	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}
