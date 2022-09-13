package account

type CreateQrCodeRes struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
}

type GenShortenRes struct {
	ShortKey string `json:"short_key"`
}

type FetchShortenRes struct {
	LongData      string `json:"long_data"`
	CreateTime    int    `json:"create_time"`
	ExpireSeconds int    `json:"expire_seconds"`
}
