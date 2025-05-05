package pkg

import (
	"workflower/app/pkg/swit"
	"workflower/app/pkg/webhook/gitlab"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(swit.NewSwitGateway),
	fx.Provide(gitlab.NewGitlabHandler),
)
