package api

import (
	"gom/app/api/middleware"
	"gom/app/api/router"

	"go.uber.org/fx"
)

var Module = fx.Options(
	router.Module,
	middleware.Module,
)
