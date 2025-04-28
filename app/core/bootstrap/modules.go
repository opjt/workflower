package bootstrap

import (
	"gom/app/api"
	"gom/app/core"
	"gom/app/lib"
	"gom/app/pkg"
	"gom/app/service"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	lib.Module,
	service.Module,
	pkg.Module,
	fx.Provide(core.NewEngine),
)
