package gitlab

type MergeRequestWebhookDTO struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

type Project struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	WebURL string `json:"web_url"`
}

type ObjectAttributes struct {
	ID           int        `json:"id"`
	TargetBranch string     `json:"target_branch"`
	SourceBranch string     `json:"source_branch"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	State        string     `json:"state"`
	URL          string     `json:"url"`
	LastCommit   LastCommit `json:"last_commit"`
	UpdatedAt    string     `json:"updated_at"`
}

type LastCommit struct {
	ID        string       `json:"id"`
	Message   string       `json:"message"`
	Timestamp string       `json:"timestamp"`
	URL       string       `json:"url"`
	Author    CommitAuthor `json:"author"`
}

type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
