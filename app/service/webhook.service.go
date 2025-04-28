package service

import (
	"gom/app/lib"
	"gom/app/pkg/swit"

	"github.com/gin-gonic/gin"
)

type WebhookService struct {
	logger  lib.Logger
	switApi *swit.SwitGateway
}

func NewWebhookService(logger lib.Logger, switApi *swit.SwitGateway) WebhookService {
	return WebhookService{
		logger:  logger,
		switApi: switApi,
	}
}

func (s WebhookService) Gitlab(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "This is a test response",
	})
}
