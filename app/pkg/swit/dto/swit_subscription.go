package dto

type SwitSubscription struct {
	Id           string `json:"id"`
	EventSource  string `json:"event_source"`
	ResourceType string `json:"resource_type"`
}
