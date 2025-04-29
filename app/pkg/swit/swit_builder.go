package swit

import (
	"bytes"
	"gom/app/lib"
	switTemplate "gom/app/pkg/swit/template"
	"gom/app/pkg/webhook/gitlab"
	"text/template"
	"time"
)

type MRTemplateData struct {
	Title           string
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

func BuildSwitMRMessage(dto gitlab.MergeRequestWebhookDTO) (string, error) {
	appID := lib.NewEnv().AppId
	tmpl, err := template.New("switMR").Parse(switTemplate.MRTemplate)
	if err != nil {
		return "", err
	}

	createdAt, _ := time.Parse(time.RFC3339, dto.ObjectAttributes.UpdatedAt)
	timeStr := createdAt.Format("15:04")

	data := MRTemplateData{
		Title:           "Merge Request",
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
