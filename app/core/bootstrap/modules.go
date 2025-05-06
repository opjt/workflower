package bootstrap

import (
	"workflower/app/api"
	"workflower/app/core"
	"workflower/app/lib"
	"workflower/app/pkg"
	"workflower/app/service"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	lib.Module,
	service.Module,
	pkg.Module,
	fx.Provide(core.NewEngine),
)

var CmdModule = fx.Options(
	lib.Module,
	pkg.Module,
)
