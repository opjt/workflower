package controller

import (
	"gom/app/lib"
	"gom/app/service"

	"github.com/gin-gonic/gin"
)

type WebhookController struct {
	service service.WebhookService
	logger  lib.Logger
}

func NewWebhookController(userService service.WebhookService, logger lib.Logger) WebhookController {
	return WebhookController{
		service: userService,
		logger:  logger,
	}
}

func (w WebhookController) Gitlab(c *gin.Context) {
	err := w.service.Gitlab(c)
	if err != nil {
		w.logger.Error(err)
	}
	c.JSON(200, gin.H{"data": "success"})
}
