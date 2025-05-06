package controller

import (
	"fmt"
	"net/http"
	"strings"
	"workflower/app/lib"
	"workflower/app/service"

	"github.com/gin-gonic/gin"
)

type OauthController struct {
	service service.OauthService
	logger  lib.Logger
}

func NewOauthController(oauthService service.OauthService, logger lib.Logger) OauthController {
	return OauthController{
		service: oauthService,
		logger:  logger,
	}
}

func (o OauthController) Swit(c *gin.Context) {
	isApp := c.Query("type")

	env := lib.NewEnv()
	clientID := env.Swit.ClientId

	// TODO: 현재 경로에서 /callback 붙이는 형식으로 전환 필요
	redirectURI := fmt.Sprintf("%s/api/v1/oauth/callback", env.Server.Url)

	scopes := []string{
		"task:write",
		"channel:write",
		"message:write",
		"workspace:read",
		"project:read",
		"channel:read",
		"subscriptions:read",
		"subscriptions:write",
		"channels.messages:read",
	}
	if isApp != "" {
		scopes = append(scopes, "app:install")
	}

	scope := strings.Join(scopes, " ")

	oauthURL := fmt.Sprintf(
		"https://openapi.swit.io/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
		clientID, redirectURI, scope,
	)

	c.Redirect(http.StatusFound, oauthURL)
}

func (o OauthController) SwitCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found"})
		return
	}
	token, err := o.service.SwitCallback(code)
	if err != nil {
		o.logger.Error(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

func (o OauthController) SwitTest(c *gin.Context) {

	result, err := o.service.SwitTest()
	if err != nil {
		o.logger.Error(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": result})
}
