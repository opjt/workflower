package swit

import (
	"encoding/json"
	"workflower/app/pkg/swit/dto"
)

func (switApi *SwitGateway) GetWorkspaceList() ([]dto.SwitWorkspace, error) {
	var result struct {
		Data struct {
			Workspaces []dto.SwitWorkspace `json:"workspaces"`
		} `json:"data"`
	}

	url := "https://openapi.swit.io/v1/api/workspace.list"

	responseBytes, err := switApi.ApiCall("GET", url, nil)
	if err != nil {
		switApi.logger.Error(err)
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &result); err != nil {
		switApi.logger.Error(err)
		return nil, err
	}

	return result.Data.Workspaces, nil
}

func (switApi *SwitGateway) GetChannelList(workspaceId string) ([]dto.SwitChannel, error) {
	var result struct {
		Data struct {
			Channels []dto.SwitChannel `json:"channels"`
		} `json:"data"`
	}

	url := "https://openapi.swit.io/v1/api/channel.list"
	body := map[string]interface{}{
		"workspace_id": workspaceId,
	}

	responseBytes, err := switApi.ApiCall("GET", url, body)
	if err != nil {
		switApi.logger.Error(err)
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &result); err != nil {
		switApi.logger.Error(err)
		return nil, err
	}

	return result.Data.Channels, nil
}

func (switApi *SwitGateway) GetSubscriptionList() ([]dto.SwitSubscription, error) {
	var result struct {
		Subscriptions []dto.SwitSubscription `json:"items"`
	}

	url := "https://openapi.swit.io/v2/subscriptions"

	responseBytes, err := switApi.ApiCall("GET", url, nil)
	if err != nil {
		switApi.logger.Error(err)
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &result); err != nil {
		switApi.logger.Error(err)
		return nil, err
	}

	return result.Subscriptions, nil
}

func (switApi *SwitGateway) DeleteSubscription(subscriptionId string) error {

	url := "https://openapi.swit.io/v2/subscriptions/" + subscriptionId

	_, err := switApi.ApiCall("DELETE", url, nil)
	if err != nil {
		switApi.logger.Error(err)
		return err
	}

	return nil

}

func (switApi *SwitGateway) CreateSubscription(event_source, resource_type string) (dto.SwitSubscription, error) {
	var result dto.SwitSubscription
	url := "https://openapi.swit.io/v2/subscriptions"

	body := map[string]interface{}{
		"event_source":  event_source,
		"resource_type": resource_type,
	}

	responseBytes, err := switApi.ApiCall("POST", url, body)
	if err != nil {
		switApi.logger.Error(err)
		return dto.SwitSubscription{}, err
	}

	if err := json.Unmarshal(responseBytes, &result); err != nil {
		switApi.logger.Error(err)
		return dto.SwitSubscription{}, err
	}
	return result, nil

}
