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
