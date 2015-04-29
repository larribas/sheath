package application

import (
	"log"
	"net/http"

	"github.com/medbook/sheath/application/domain"
)

// CreateLink receives a raw URL and the request issuing the use case, validates the link and stores it.
// It returns a stub to retrieve the original URL by, and a Validation or Storage error, if such is the case
func (app *App) CreateLink(rawURL string, r *http.Request) (linkStub string, err error) {
	link, err := domain.NewLink(rawURL)
	if err != nil {
		return
	}

	if err = app.Validator.Validate(link); err != nil {
		return
	}

	if err = app.LinkRepository.Store(link); err != nil {
		log.Printf("The following error occured while storing a link: %s\n", err.Error())
		return
	}

	app.Notifier.NotifyEvent(domain.NewLinkCreatedEvent(link, r))
	linkStub = link.Stub
	return
}
