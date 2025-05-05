package gitlab

import (
	"encoding/json"
	"errors"
	"gom/app/lib"
	"gom/app/pkg/swit"
)

type GitlabHandler struct {
	logger  lib.Logger
	switApi *swit.SwitGateway
}

func NewGitlabHandler(logger lib.Logger, switApi *swit.SwitGateway) GitlabHandler {
	return GitlabHandler{logger: logger, switApi: switApi}
}

func (g GitlabHandler) HandleMergeRequest(bodyBytes []byte) error {
	var webhook MergeRequestWebhookDTO
	if err := json.Unmarshal(bodyBytes, &webhook); err != nil {
		g.logger.Error(err)
		return errors.New("Failed to parse MR payload")
	}

	message, err := BuildSwitMRMessage(webhook)
	if err != nil {
		g.logger.Error(err)
		return errors.New("Failed to convert Swit Content")
	}

	g.logger.Info("MR event:", webhook)
	if err := g.switApi.SendChannel(lib.NewEnv().Swit.ChannelId, message); err != nil {
		g.logger.Error(err)
		return err
	}

	return nil
}

func (g *GitlabHandler) HandlePushRequest(bodyBytes []byte) error {
	// 여기서 Push 이벤트 처리 로직을 추가할 수 있습니다.
	// 예시로는 webhookparser.ParseGitlabPushRequest와 비슷한 로직을 처리
	return nil
}
