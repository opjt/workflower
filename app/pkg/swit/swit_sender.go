package swit

import (
	"encoding/json"
	"fmt"
)

func (switApi *SwitGateway) SendChannel(channelID, messageContent string) error {
	var value json.RawMessage = json.RawMessage(messageContent)

	body := map[string]interface{}{
		"channel_id": channelID,
		"content":    " ",
		"body_type":  "plain",
		"attachments": []map[string]interface{}{
			{
				"attachment_type": "custom",
				"values":          []json.RawMessage{value},
			},
		},
	}

	url := "https://openapi.swit.io/v1/api/message.create"
	if err := switApi.ApiCall(url, body); err != nil {
		return fmt.Errorf("failed to send message to Swit: %w", err)
	}

	return nil
}
