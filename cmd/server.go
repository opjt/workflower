package cmd

import (
	"workflower/app/api/middleware"
	"workflower/app/api/router"
	"workflower/app/core"
	"workflower/app/lib"

	"github.com/spf13/cobra"
)

type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run(c *cobra.Command, args []string) CommandRunner {
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

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
