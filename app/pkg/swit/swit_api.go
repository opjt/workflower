package swit

import (
	"encoding/json"
	"fmt"
	"workflower/app/pkg/swit/dto"
)

func (switApi *SwitGateway) SendChannel(channelID, messageContent string) error {
	var value json.RawMessage = json.RawMessage(messageContent)

	body := map[string]any{
		"channel_id": channelID,
		"content":    " ",
		"body_type":  "plain",
		"attachments": []map[string]any{
			{
				"attachment_type": "custom",
				"values":          []json.RawMessage{value},
			},
		},
	}

	url := "https://openapi.swit.io/v1/api/message.create"
	if _, err := switApi.ApiCall("POST", url, body); err != nil {
		return fmt.Errorf("failed to send message to Swit: %w", err)
	}

	return nil
}

func (switApi *SwitGateway) GetChannelInfo(channelID string) (dto.SwitChannel, error) {

	var result struct {
		Data struct {
			Channel dto.SwitChannel `json:"channel"`
		} `json:"data"`
	}

	var channel dto.SwitChannel
	url := "https://openapi.swit.io/v1/api/channel.info"
	body := map[string]interface{}{
		"id": channelID,
	}
	responseBytes, err := switApi.ApiCall("GET", url, body)
	switApi.logger.Info(string(responseBytes))
	if err != nil {
		switApi.logger.Error(err)
		return channel, err
	}

	if err := json.Unmarshal(responseBytes, &result); err != nil {
		return channel, err
	}

	return result.Data.Channel, nil
}
