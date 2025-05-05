package service

import (
	"gom/app/lib"
	"gom/app/pkg/swit"

	"github.com/gin-gonic/gin"
)

type TestService struct {
	logger  lib.Logger
	switApi *swit.SwitGateway
}

func NewTestService(logger lib.Logger, switApi *swit.SwitGateway) TestService {
	return TestService{
		logger:  logger,
		switApi: switApi,
	}
}

func (s TestService) Test(c *gin.Context) {
	url := "https://openapi.swit.io/v1/api/message.create" // 호출하려는 URL
	body := map[string]any{
		"body_type":  "plain",
		"channel_id": lib.NewEnv().Swit.ChannelId,
		"content":    "ping test",
	}

	// ApiCall 함수 호출
	err := s.switApi.ApiCall(url, body)
	if err != nil {
		s.logger.Error(err)
	}
	c.JSON(200, gin.H{
		"message": "This is a test response",
	})
}
