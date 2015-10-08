package httpapi

import (
	"github.com/MedBrain/sheath/application"
	"github.com/MedBrain/sheath/notifiers/null"
	"github.com/MedBrain/sheath/storage/inmemory"
	"github.com/MedBrain/sheath/validators/collection"
	"github.com/MedBrain/sheath/validators/protocol"
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
