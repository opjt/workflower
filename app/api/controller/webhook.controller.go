package controller

import (
	"workflower/app/lib"
	"workflower/app/service"

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

func (w WebhookController) SwitEvent(c *gin.Context) {
	rawBody, err := c.GetRawData()
	if err != nil {
		w.logger.Error(err)
		c.JSON(500, gin.H{"error": "failed to read request body"})
		return
	}
	w.logger.Info(string(rawBody))
	// Content-Type 설정 (예: JSON)
	c.Data(200, "application/json", rawBody)
}
