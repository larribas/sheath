package httpapi

import (
	"net/http"

	"github.com/larribas/sheath/application"
)

func createLink(app *application.App, w http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You must supply a 'url' to shorten"))
		return
	}

	linkStub, err := app.CreateLink(url, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(linkStub))
}
