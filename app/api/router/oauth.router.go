package router

import (
	"workflower/app/api/controller"
	"workflower/app/core"
	"workflower/app/lib"
)

type OauthRoutes struct {
	logger          lib.Logger
	engine          core.Engine
	oauthController controller.OauthController
}

func NewOauthRoutes(
	logger lib.Logger,
	engine core.Engine,
	oauthController controller.OauthController,
) OauthRoutes {
	return OauthRoutes{
		engine:          engine,
		logger:          logger,
		oauthController: oauthController,
	}
}

func (r OauthRoutes) Setup() {
	oauthRoutes := r.engine.ApiGroup.Group("/oauth")
	{
		oauthRoutes.GET("", r.oauthController.Swit)
		oauthRoutes.GET("/callback", r.oauthController.SwitCallback)
		oauthRoutes.GET("/test", r.oauthController.SwitTest)

	}
}
