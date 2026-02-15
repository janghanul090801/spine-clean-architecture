package main

import (
	"os"
	"time"

	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/boot"
	"github.com/janghanul090801/spine-clean-architecture/api/controller"
	"github.com/janghanul090801/spine-clean-architecture/api/route"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/infra/database"
	"github.com/janghanul090801/spine-clean-architecture/infra/model"
	"github.com/janghanul090801/spine-clean-architecture/infra/repository"
	"github.com/janghanul090801/spine-clean-architecture/interceptor"
	"github.com/janghanul090801/spine-clean-architecture/internal/logger"
	"github.com/janghanul090801/spine-clean-architecture/usecase"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

func main() {
	config.NewEnv()

	app := spine.New()

	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.RegisterModel()

	getLogger := logger.GetLogger()
	getLogger.Info("Starting Spine Server")

	db.RegisterModel(
		(*model.UserModel)(nil),
		(*model.TaskModel)(nil),
	)

	app.Constructor(
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
		usecase.NewAuthUseCase,
		usecase.NewProfileUsecase,

		// Controller
		controller.NewTaskController,
		controller.NewProfileController,
		controller.NewAuthController,

		// Interceptor
		interceptor.NewTxInterceptor,
		interceptor.NewAuthInterceptor,
	)

	app.Interceptor(
		interceptor.NewCORSInterceptor(),
		interceptor.NewRateLimitInterceptor(),
		// interceptor.NewLoggingInterceptor(),
		interceptor.NewErrorInterceptor(),
	)

	route.NewProfileRouter(app)
	route.NewAuthRouter(app)
	route.NewTaskRouter(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	getLogger.Info("Server starting", zap.String("port", port))
	app.Run(boot.Options{
		Address:                ":" + port,
		EnableGracefulShutdown: true,
		HTTP:                   &boot.HTTPOptions{},
	})
}
