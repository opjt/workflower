package bootstrap

import (
	"context"
	"gom/app/api/middleware"
	"gom/app/api/router"
	"gom/app/core"

	"gom/app/lib"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func RunServer(opt fx.Option) {
	logger := lib.GetLogger()
	opts := fx.Options(
		fx.WithLogger(func() fxevent.Logger {
			return logger.GetFxLogger()
		}),
		fx.Invoke(Run()),
	)
	ctx := context.Background()
	app := fx.New(opt, opts)
	err := app.Start(ctx)
	defer app.Stop(ctx)
	if err != nil {
		logger.Fatal(err)
	}
}

func Run() any {
	return func(
		middleware middleware.Middlewares,
		env lib.Env,
		engine core.Engine,
		route router.Routes,
		logger lib.Logger,

	) {
		middleware.Setup()
		route.Setup()

		logger.Info("Running server")

		if env.Server.Port == "" {
			_ = engine.Gin.Run()
		} else {
			_ = engine.Gin.Run(":" + env.Server.Port)
		}
	}
}
