package webhook

type Config struct {
	UserName   string `json:"username"`
	Abatar_url string `json:"avatar_url"`
	WebhookUrl string `json:"webhook_url"`
}

type MinMessage struct {
	Username   string `json:"username"`
	Content    string `json:"content"`
	Avatar_url string `json:"avatar_url"`
}

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Embed struct {
	Color  uint64  `json:"color"`
	Fields []Field `json:"fields"`
}

// type EmbedMessage struct {
// 	MinMessage MinMessage
// 	Embeds     []Embed `json:"embeds"`
// }

type EmbedMessage struct {
	Username   string  `json:"username"`
	Content    string  `json:"content"`
	Avatar_url string  `json:"avatar_url"`
	Embeds     []Embed `json:"embeds"`
}
