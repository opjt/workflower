package api

import (
	"workflower/app/api/controller"
	"workflower/app/api/middleware"
	"workflower/app/api/router"

	"go.uber.org/fx"
)

var Module = fx.Options(
	router.Module,
	middleware.Module,
	controller.Module,
)
