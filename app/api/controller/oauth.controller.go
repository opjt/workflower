package controller

import (
	"fmt"
	"net/http"
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
	env := lib.NewEnv()
	clientID := env.Swit.ClientId

	redirectURI := fmt.Sprintf("%s/api/v1/oauth/callback", env.Server.Url)

	scope := "task:write channel:write message:write app:install"

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
	}
	c.JSON(200, gin.H{"token": token})
}
