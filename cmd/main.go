package main

import (
	"github.com/NARUBROWN/spine/pkg/boot"
	"github.com/janghanul090801/spine-clean-architecture/api/controller"
	"github.com/janghanul090801/spine-clean-architecture/api/route"
	"github.com/janghanul090801/spine-clean-architecture/bootstrap"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/infra/model"
	"github.com/janghanul090801/spine-clean-architecture/infra/repository"
	"github.com/janghanul090801/spine-clean-architecture/interceptor"
	"github.com/janghanul090801/spine-clean-architecture/internal/logger"
	"github.com/janghanul090801/spine-clean-architecture/usecase"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"os"
	"time"
)

func main() {

	app := bootstrap.App()

	db := app.DB

	defer app.CloseDBConnection()

	getLogger := logger.GetLogger()
	getLogger.Info("Starting Spine Server")

	db.RegisterModel(
		(*model.UserModel)(nil),
		(*model.TaskModel)(nil),
	)

	app.App.Constructor(
		// DB
		func() *bun.DB { return db },

		// etc.
		func() *zap.Logger { return getLogger },
		func() time.Duration { return time.Duration(config.E.ContextTimeout) * time.Second },

		// Repository
		repository.NewUserRepository,
		repository.NewTaskRepository,

		// Usecase
		usecase.NewTaskUsecase,
		usecase.NewSignupUsecase,
		usecase.NewLoginUsecase,
		usecase.NewProfileUsecase,
		usecase.NewRefreshTokenUsecase,

		// Controller
		controller.NewTaskController,
		controller.NewProfileController,
		controller.NewSignupController,
		controller.NewLoginController,
		controller.NewRefreshTokenController,

		// Interceptor
		interceptor.NewTxInterceptor,
		interceptor.NewAuthInterceptor,
	)

	app.App.Interceptor(
		interceptor.NewCORSInterceptor(),
		interceptor.NewRateLimitInterceptor(),
		// interceptor.NewLoggingInterceptor(),
		interceptor.NewErrorInterceptor(),
	)

	route.NewLoginRouter(app.App)
	route.NewSignupRouter(app.App)
	route.NewProfileRouter(app.App)
	route.NewRefreshTokenRouter(app.App)
	route.NewTaskRouter(app.App)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	getLogger.Info("Server starting", zap.String("port", port))
	app.App.Run(boot.Options{
		Address:                ":" + port,
		EnableGracefulShutdown: true,
		HTTP:                   &boot.HTTPOptions{},
	})
}
