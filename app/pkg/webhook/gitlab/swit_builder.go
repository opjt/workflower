package gitlab

import (
	"bytes"
	"gom/app/lib"
	switTemplate "gom/app/pkg/swit/template"
	"text/template"
)

type MRTemplateData struct {
	AuthorName      string
	AuthorUsername  string
	AppID           string
	MRTitle         string
	AuthorAvatarURL string
	MRURL           string
	ProjectName     string
	SourceBranch    string
	TargetBranch    string
	MRState         string
	CreatedAt       string
}

func BuildSwitMRMessage(dto MergeRequestWebhookDTO) (string, error) {
	appID := lib.NewEnv().Swit.AppId
	tmpl, err := template.New("switMR").Parse(switTemplate.MRTemplate)
	if err != nil {
		return "", err
	}

	updatedAt := dto.ObjectAttributes.UpdatedAt
	timeStr := updatedAt[11:16] // "13:25"

	data := MRTemplateData{
		AuthorName:      dto.User.Name,
		AuthorUsername:  dto.User.Username,
		AppID:           appID,
		MRTitle:         dto.ObjectAttributes.Title,
		AuthorAvatarURL: dto.User.AvatarURL,
		MRURL:           dto.ObjectAttributes.URL,
		ProjectName:     dto.Project.Name,
		SourceBranch:    dto.ObjectAttributes.SourceBranch,
		TargetBranch:    dto.ObjectAttributes.TargetBranch,
		MRState:         dto.ObjectAttributes.State,
		CreatedAt:       timeStr,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
