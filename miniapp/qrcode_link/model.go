package qrcode_link

const (
	EnvVersionRelease = "release"
	EnvVersionTrial   = "trial"
	EnvVersionDevelop = "develop"
)

type WxaQrcodeRes struct {
	ContentType string `json:"contentType"`
	Buffer      []byte `json:"buffer"`
}

type GetQrcodeParams struct {
	Width     int  `json:"width"`
	AutoColor bool `json:"auto_color"`
	IsHyaline bool `json:"is_hyaline"`
	LineColor struct {
		R int `json:"r"`
		G int `json:"g"`
		B int `json:"b"`
	} `json:"line_color"`
}

type GetUnlimitedQrcodeParams struct {
	GetQrcodeParams
	Path       string `json:"path"`
	CheckPath  bool   `json:"check_path"`
	EnvVersion string `json:"env_version"`
}

type SchemeData struct {
	SchemeInfo struct {
		Appid      string `json:"appid"`
		Path       string `json:"path"`
		Query      string `json:"query"`
		CreateTime int    `json:"create_time"`
		ExpireTime int    `json:"expire_time"`
		EnvVersion string `json:"env_version"`
	} `json:"scheme_info"`
	SchemeQuota struct {
		LongTimeUsed  int `json:"long_time_used"`
		LongTimeLimit int `json:"long_time_limit"`
	} `json:"scheme_quota"`
}

type GenerateSchemeParams struct {
	JumpWxa struct {
		Path       string `json:"path"`
		Query      string `json:"query"`
		EnvVersion string `json:"env_version"`
	} `json:"jump_wxa"`
	IsExpire       bool `json:"is_expire"`
	ExpireTime     int  `json:"expire_time"`
	ExpireType     int  `json:"expire_type"`
	ExpireInterval int  `json:"expire_interval"`
}

type GenerateNFCSchemeParams struct {
	JumpWxa struct {
		Path       string `json:"path"`
		Query      string `json:"query"`
		EnvVersion string `json:"env_version"`
	} `json:"jump_wxa"`
	Sn      string `json:"sn"`
	ModelId string `json:"model_id"`
}

type GenerateSchemeRes struct {
	OpenLink string `json:"openlink"`
}

type GenerateUrlLinkParams struct {
	Path           string `json:"path"`
	Query          string `json:"query"`
	IsExpire       bool   `json:"is_expire"`
	ExpireType     int    `json:"expire_type"`
	ExpireTime     int    `json:"expire_time"`
	ExpireInterval int    `json:"expire_interval"`
	EnvVersion     string `json:"env_version"`
	CloudBase      struct {
		Env           string `json:"env"`
		Domain        string `json:"domain"`
		Path          string `json:"path"`
		Query         string `json:"query"`
		ResourceAppid string `json:"resource_appid"`
	} `json:"cloud_base"`
}

type GenerateUrlLinkRes struct {
	UrlLink string `json:"url_link"`
}

type QueryUrlLinkRes struct {
	UrlLinkInfo struct {
		Appid      string `json:"appid"`
		Path       string `json:"path"`
		Query      string `json:"query"`
		CreateTime int    `json:"create_time"`
		ExpireTime int    `json:"expire_time"`
		EnvVersion string `json:"env_version"`
		CloudBase  struct {
			Env           string `json:"env"`
			Doamin        string `json:"doamin"`
			Path          string `json:"path"`
			Query         string `json:"query"`
			ResourceAppid string `json:"resource_appid"`
		} `json:"cloud_base"`
	} `json:"url_link_info"`
	UrlLinkQuota struct {
		LongTimeUsed  int `json:"long_time_used"`
		LongTimeLimit int `json:"long_time_limit"`
	} `json:"url_link_quota"`
	VisitOpenid string `json:"visit_openid"`
}

type GenerateShortLinkRes struct {
	Link string `json:"link"`
}
