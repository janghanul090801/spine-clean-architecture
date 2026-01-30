package bootstrap

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/ent"
)

type Application struct {
	Env    *Env
	Client *ent.Client
	App    *fiber.App
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()

	var err error
	app.Client, err = NewClient(app.Env)
	if err != nil {
		panic(err)
	}

	if err = app.Client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	app.App = fiber.New(fiber.Config{
		AppName:      "Fiber Ent Clean Architecture",
		ServerHeader: "Fiber",
	})

	// Use global middlewares.
	app.App.Use(cors.New())
	app.App.Use(compress.New())
	app.App.Use(etag.New())
	app.App.Use(favicon.New())
	app.App.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))
	app.App.Use(logger.New())
	app.App.Use(recover.New())
	app.App.Use(requestid.New())

	return *app
}

func (app *Application) CloseDBConnection() {
	app.Client.Close()
}
