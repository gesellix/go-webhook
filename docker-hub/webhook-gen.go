package dockerhub

type DockerHub struct {
	CallbackURL string `json:"callback_url"`
	PushData    struct {
		Images   []string `json:"images"`
		PushedAt int64    `json:"pushed_at"`
		Pusher   string   `json:"pusher"`
	} `json:"push_data"`
	Repository struct {
		CommentCount    int64  `json:"comment_count"`
		DateCreated     int64  `json:"date_created"`
		Description     string `json:"description"`
		Dockerfile      string `json:"dockerfile"`
		FullDescription string `json:"full_description"`
		IsOfficial      bool   `json:"is_official"`
		IsPrivate       bool   `json:"is_private"`
		IsTrusted       bool   `json:"is_trusted"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		Owner           string `json:"owner"`
		RepoName        string `json:"repo_name"`
		RepoURL         string `json:"repo_url"`
		StarCount       int64  `json:"star_count"`
		Status          string `json:"status"`
	} `json:"repository"`
}
