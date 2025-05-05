package router

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewOauthRoutes),
	fx.Provide(NewWebhookRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	oauthRoutes OauthRoutes,
	webhookRoutes WebhookRoutes,

) Routes {
	return Routes{
		oauthRoutes, webhookRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
