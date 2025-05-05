package pkg

import (
	"gom/app/pkg/swit"
	"gom/app/pkg/webhook/gitlab"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(swit.NewSwitGateway),
	fx.Provide(gitlab.NewGitlabHandler),
)
