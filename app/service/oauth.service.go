package service

import (
	"fmt"
	"gom/app/lib"
	"gom/app/pkg/swit"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OauthService struct {
	logger  lib.Logger
	switApi *swit.SwitGateway
}

func NewOauthService(logger lib.Logger, switApi *swit.SwitGateway) OauthService {
	return OauthService{
		logger:  logger,
		switApi: switApi,
	}
}

func (s OauthService) Oauth(c *gin.Context) {
	env := lib.NewEnv()
	clientID := env.ClientId

	redirectURI := fmt.Sprintf("%s/api/v1/oauth/callback", env.ServerUrl)

	scope := "task:write channel:write message:write app:install"

	oauthURL := fmt.Sprintf(
		"https://openapi.swit.io/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
		clientID, redirectURI, scope,
	)

	c.Redirect(http.StatusFound, oauthURL)
}

func (s OauthService) Callback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found"})
		return
	}

	// 여기서 code를 저장하거나, 바로 token 요청 보내기
	fmt.Println("Authorization Code:", code)

	_, err := s.switApi.GetToken(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accessToken not found"})
		return
	}
	c.JSON(200, gin.H{"message": "good response"})
}
