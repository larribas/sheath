package main // import "github.com/medbook/sheath"

import (
	"github.com/medbook/sheath/application"
	"github.com/medbook/sheath/notifiers/null"
	"github.com/medbook/sheath/ports/httpapi"
	"github.com/medbook/sheath/storage/redis"
	"github.com/medbook/sheath/validators/collection"
	"github.com/medbook/sheath/validators/protocol"
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
