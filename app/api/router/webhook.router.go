package router

import (
	"gom/app/core"
	"gom/app/lib"
	"gom/app/service"
)

type WebhookRoutes struct {
	logger         lib.Logger
	engine         core.Engine
	webhookService service.WebhookService
}

func NewWebhookRoutes(
	logger lib.Logger,
	engine core.Engine,
	webhookService service.WebhookService,

) WebhookRoutes {
	return WebhookRoutes{
		engine:         engine,
		logger:         logger,
		webhookService: webhookService,
	}
}

func (r WebhookRoutes) Setup() {
	gitlabRoutes := r.engine.ApiGroup.Group("/webhook/gitlab")
	{
		gitlabRoutes.POST("", r.webhookService.Gitlab)

	}
}
