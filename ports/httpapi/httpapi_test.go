package httpapi

import (
	"github.com/medbook/sheath/application"
	"github.com/medbook/sheath/notifiers/null"
	"github.com/medbook/sheath/storage/inmemory"
	"github.com/medbook/sheath/validators/collection"
	"github.com/medbook/sheath/validators/protocol"
)

func newApp() *application.App {
	return application.NewApp(
		inmemory.NewLinkRepository(),
		&null.Notifier{},
		collection.New(
			map[string]application.Validator{
				"protocol": protocol.NewValidator("http", "https"),
			},
		),
	)
}
