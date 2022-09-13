package security

const (
	SecVersion                  = 2
	MediaTypeImage              = 2
	MediaTypeVoice              = 1
	SceneProfile                = 1
	SceneComment                = 2
	SceneForum                  = 3
	SceneSocialLog              = 4
	SceneRegister               = 0
	SceneCheatingInTheMarketing = 1
)

type MsgSecCheckParams struct {
	Openid    string `json:"openid"`
	Scene     int    `json:"scene"`
	Version   int    `json:"version"`
	Content   string `json:"content"`
	Title     string `json:"title"`
	Signature string `json:"signature"`
	Nickname  string `json:"nickname"`
}

type MsgSecCheckDetail struct {
	Strategy string `json:"strategy"`
	ErrCode  int    `json:"errcode"`
	Suggest  string `json:"suggest"`
	Label    int    `json:"label"`
	Prob     int    `json:"prob,omitempty"`
	Level    int    `json:"level,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
}

type MsgSecCheckRes struct {
	Result struct {
		Suggest string `json:"suggest"`
		Label   int    `json:"label"`
	} `json:"result"`
	Detail  []MsgSecCheckDetail `json:"detail"`
	TraceId string              `json:"trace_id"`
}

type MediaCheckAsyncParams struct {
	Openid    string `json:"openid"`
	Scene     int    `json:"scene"`
	Version   int    `json:"version"`
	MediaUrl  string `json:"media_url"`
	MediaType int    `json:"media_type"`
}

type MediaCheckAsyncRes struct {
	TraceId string `json:"trace_id"`
}

type GetUserRiskRankParams struct {
	Appid        string `json:"appid"`
	Openid       string `json:"openid"`
	Scene        int    `json:"scene"`
	MobileNo     string `json:"mobile_no"`
	BankCardNo   string `json:"bank_card_no"`
	CertNo       string `json:"cert_no"`
	ClientIp     string `json:"client_ip"`
	EmailAddress string `json:"email_address"`
	ExtendedInfo string `json:"extended_info"`
	IsTest       bool   `json:"is_test"`
}

type GetUserRiskRankRes struct {
	RiskRank int `json:"risk_rank"`
	UnoinId  int `json:"unoin_id"`
}
