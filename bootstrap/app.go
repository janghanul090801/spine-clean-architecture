package bootstrap

import (
	"github.com/NARUBROWN/spine"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/infra/database"
	"github.com/uptrace/bun"
)

type Application struct {
	DB  *bun.DB
	App spine.App
}

func App() Application {
	app := &Application{}
	config.NewEnv()

	var err error
	app.DB, err = database.NewDB()
	if err != nil {
		panic(err)
	}

	app.DB.RegisterModel()

	app.App = spine.New()

	return *app
}

func (app *Application) CloseDBConnection() {
	app.DB.Close()
}
