package service

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewOauthService),
	fx.Provide(NewWebhookService),
)
