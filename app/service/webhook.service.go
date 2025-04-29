package service

import (
	"encoding/json"
	"gom/app/lib"
	"gom/app/pkg/swit"
	"gom/app/pkg/webhook/gitlab"
	"net/http"

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
	// 1. raw body 복사해서 읽기
	var raw map[string]interface{}
	if err := c.ShouldBindJSON(&raw); err != nil {
		s.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	objectKind, ok := raw["object_kind"].(string)
	if !ok {
		s.logger.Error("object_kind missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "object_kind missing"})
		return
	}
	// 다시 json으로 변환
	bodyBytes, _ := json.Marshal(raw)

	switch objectKind {
	case "merge_request":
		// MR DTO로 변환

		var webhook gitlab.MergeRequestWebhookDTO
		err := json.Unmarshal(bodyBytes, &webhook)

		if err != nil {
			s.logger.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse MR payload"})
			return
		}
		message, err := swit.BuildSwitMRMessage(webhook)
		if err != nil {
			s.logger.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to convert Swit Content"})
			return
		}
		s.logger.Info("MR event:", webhook)
		if err := s.switApi.SendChannel("25040800002914X1A5DZ", message); err != nil {
			s.logger.Error(err)
		}
	// case "push":
	// 	// Push DTO로 변환
	// 	parsed, err := webhookparser.ParseGitlabPushRequest(c)
	// 	if err != nil {
	// 		s.logger.Error(err)
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse Push payload"})
	// 		return
	// 	}
	// 	s.logger.Info("Push event:", parsed)
	// 	s.switHandler.SendSwitPush(parsed)
	default:
		s.logger.Warnf("Unhandled event type: %s", objectKind)
		c.JSON(http.StatusOK, gin.H{"message": "Event ignored"})
	}
}
