package main // import "github.com/larribas/sheath"

import (
	"github.com/larribas/sheath/application"
	"github.com/larribas/sheath/notifiers/null"
	"github.com/larribas/sheath/ports/httpapi"
	"github.com/larribas/sheath/storage/redis"
	"github.com/larribas/sheath/validators/collection"
	"github.com/larribas/sheath/validators/protocol"
)

// SheathApp defines and exposes the topology (the particular implementations running on each interface)
// of the Sheath application instance that will run when building and executing the program
func SheathApp() *application.App {
	return application.NewApp(
		redis.NewLinkRepository(),
		null.Notifier{},
		collection.New(
			map[string]application.Validator{
				"protocol": protocol.NewValidator("http", "https"),
			},
		),
	)
}

func main() {

	app := SheathApp()

	p := httpapi.NewPort(app)
	p.Expose()
}
