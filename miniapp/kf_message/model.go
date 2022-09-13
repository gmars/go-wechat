package kf_message

const (
	TypingCommand       = "Typing"
	CancelTypingCommand = "CancelTyping"
)

type UploadRes struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
}

type GetTempMediaRes struct {
	Buffer []byte `json:"buffer"`
}

type LinkMsg struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ThumbUrl    string `json:"thumb_url"`
}

type MiniProgramPageMsg struct {
	Title        string `json:"title"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}
