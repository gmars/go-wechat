package user_info

type LoginRes struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

type OpenIdRes struct {
	OpenPid string `json:"openpid"`
}

type CheckEncryptedMsg struct {
	Vaild      bool `json:"vaild"`
	CreateTime int  `json:"create_time"`
}

type PaymentAddition struct {
	TransactionId string `json:"transaction_id"`
	MchId         string `json:"mch_id"`
	OutTradeNo    string `json:"out_trade_no"`
}

type PaymentUnionRes struct {
	UnionId string `json:"unionid"`
}

type UserEncryptKeyInfoItem struct {
	EncryptKey string `json:"encrypt_key"`
	Version    int    `json:"version"`
	ExpireIn   int    `json:"expire_in"`
	Iv         string `json:"iv"`
	CreateTime int    `json:"create_time"`
}

type UserEncryptKeyRes struct {
	KeyInfoList []UserEncryptKeyInfoItem `json:"key_info_list"`
}

type PhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     int    `json:"countryCode"`
	Watermark       struct {
		Timestamp int    `json:"timestamp"`
		Appid     string `json:"appid"`
	} `json:"watermark"`
}

type PhoneInfoRes struct {
	PhoneInfo PhoneInfo `json:"phone_info"`
}
