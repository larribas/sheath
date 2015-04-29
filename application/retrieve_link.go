package application

import (
	"github.com/medbook/sheath/application/domain"
	"net/http"
)

// RetrieveLink receives a link stub identifying a previously stored link, and the request issuing the use case,
// It returns the original URL corresponding to such stub, and an ErrNotFound error (if such is the case)
func (app *App) RetrieveLink(linkStub string, r *http.Request) (originalURL string, err error) {
	link, err := app.LinkRepository.Find(linkStub)
	if err != nil {
		return
	}

	app.Notifier.NotifyEvent(domain.NewLinkRetrievedEvent(link, r))
	originalURL = link.Original.String()
	return
}
