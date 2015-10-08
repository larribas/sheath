package main // import "github.com/MedBrain/sheath"

import (
	"github.com/MedBrain/sheath/application"
	"github.com/MedBrain/sheath/notifiers/null"
	"github.com/MedBrain/sheath/ports/httpapi"
	"github.com/MedBrain/sheath/storage/redis"
	"github.com/MedBrain/sheath/validators/collection"
	"github.com/MedBrain/sheath/validators/protocol"
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
