package router

import (
	"gom/app/core"
	"gom/app/lib"
	"gom/app/service"
)

type OauthRoutes struct {
	logger       lib.Logger
	engine       core.Engine
	oauthService service.OauthService
}

func NewOauthRoutes(
	logger lib.Logger,
	engine core.Engine,
	oauthService service.OauthService,
) OauthRoutes {
	return OauthRoutes{
		engine:       engine,
		logger:       logger,
		oauthService: oauthService,
	}
}

func (r OauthRoutes) Setup() {
	oauthRoutes := r.engine.ApiGroup.Group("/oauth")
	{
		oauthRoutes.GET("", r.oauthService.Oauth)
		oauthRoutes.GET("/callback", r.oauthService.Callback)

	}
}
