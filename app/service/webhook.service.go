package service

import (
	"encoding/json"
	"errors"
	"workflower/app/lib"
	"workflower/app/pkg/webhook/gitlab"

	"github.com/gin-gonic/gin"
)

type WebhookService struct {
	logger        lib.Logger
	gitlabHandler gitlab.GitlabHandler
}

func NewWebhookService(logger lib.Logger, gitlabHandler gitlab.GitlabHandler) WebhookService {
	return WebhookService{
		logger:        logger,
		gitlabHandler: gitlabHandler,
	}
}

func (s WebhookService) Gitlab(c *gin.Context) error {
	// 1. raw body 복사해서 읽기
	var raw map[string]interface{}
	if err := c.ShouldBindJSON(&raw); err != nil {
		s.logger.Error(err)
		return errors.New("invalid payload")

	}

	objectKind, ok := raw["object_kind"].(string)
	if !ok {
		s.logger.Error("object_kind missing")
		return errors.New("object_kind missing")
	}
	// 다시 json으로 변환
	bodyBytes, _ := json.Marshal(raw)

	switch objectKind {
	case "merge_request":
		return s.gitlabHandler.HandleMergeRequest(bodyBytes)
	default:
		s.logger.Warnf("Unhandled event type: %s", objectKind)
		return nil
	}

}
