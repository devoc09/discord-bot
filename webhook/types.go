package webhook

type Config struct {
	UserName   string `json:"username"`
	Abatar_url string `json:"avatar_url"`
	WebhookUrl string `json:"webhook_url"`
}

type Message struct {
	Username   string `json:"username"`
	Content    string `json:"content"`
	Avatar_url string `json:"avatar_url"`
}
