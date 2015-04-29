package httpapi

import (
	"github.com/larribas/sheath/application"
	"github.com/larribas/sheath/notifiers/null"
	"github.com/larribas/sheath/storage/inmemory"
	"github.com/larribas/sheath/validators/collection"
	"github.com/larribas/sheath/validators/protocol"
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
