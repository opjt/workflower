package service

import (
	"errors"
	"workflower/app/lib"
	"workflower/app/pkg/swit"
	"workflower/app/pkg/swit/dto"
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

func (s OauthService) SwitCallback(code string) (swit.TokenStore, error) {

	token, err := s.switApi.GetToken(code)
	if err != nil {
		return swit.TokenStore{}, errors.New("accessToken not found")
	}

	return token, nil
}

func (s OauthService) SwitTest() (dto.SwitChannel, error) {

	var channel dto.SwitChannel
	channel, err := s.switApi.GetChannelInfo(lib.NewEnv().Swit.ChannelId)
	if err != nil {
		return channel, err
	}

	return channel, nil
}
