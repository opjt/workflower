package service

import (
	"errors"
	"workflower/app/lib"
	"workflower/app/pkg/swit"
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
