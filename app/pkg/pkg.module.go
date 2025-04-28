package pkg

import (
	"gom/app/pkg/swit"

	"go.uber.org/fx"
)

var Module = fx.Options(

	fx.Provide(swit.NewSwitGateway),
)
