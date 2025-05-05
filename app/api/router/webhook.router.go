package router

import (
	"gom/app/api/controller"
	"gom/app/core"
	"gom/app/lib"
)

type WebhookRoutes struct {
	logger            lib.Logger
	engine            core.Engine
	webhookController controller.WebhookController
}

func NewWebhookRoutes(
	logger lib.Logger,
	engine core.Engine,
	webhookController controller.WebhookController,

) WebhookRoutes {
	return WebhookRoutes{
		engine:            engine,
		logger:            logger,
		webhookController: webhookController,
	}
}

func (r WebhookRoutes) Setup() {
	gitlabRoutes := r.engine.ApiGroup.Group("/webhook/gitlab")
	{
		gitlabRoutes.POST("", r.webhookController.Gitlab)

	}
}
