package basic_info

type WxVerifyInfo struct {
	QualificationVerify   bool `json:"qualification_verify"`
	NamingVerify          bool `json:"naming_verify"`
	AnnualReview          bool `json:"annual_review"`
	AnnualReviewBeginTime int  `json:"annual_review_begin_time"`
	AnnualReviewEndTime   int  `json:"annual_review_end_time"`
}

type SignatureInfo struct {
	Signature       string `json:"signature"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

type HeadImageInfo struct {
	HeadImageUrl    string `json:"head_image_url"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

type NicknameInfo struct {
	Nickname        string `json:"nickname"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

type AccountBasicInfoRes struct {
	Appid             string        `json:"appid"`
	AccountType       int           `json:"account_type"`
	PrincipalType     int           `json:"principal_type"`
	PrincipalName     string        `json:"principal_name"`
	RealNameStatus    int           `json:"realname_status"`
	WxVerifyInfo      WxVerifyInfo  `json:"wx_verify_info"`
	SignatureInfo     SignatureInfo `json:"signature_info"`
	HeadImageInfo     HeadImageInfo `json:"head_image_info"`
	Nickname          string        `json:"nickname"`
	RegisteredCountry int           `json:"registered_country"`
	NicknameInfo      NicknameInfo  `json:"nickname_info"`
	Credential        string        `json:"credential"`
	CustomerType      int           `json:"customer_type"`
}

type HaveOpenRes struct {
	HaveOpen bool `json:"have_open"`
}

type CheckNickNameRes struct {
	HitCondition bool   `json:"hit_condition"`
	Wording      string `json:"wording"`
}

type SetNickNameParams struct {
	NickName          string `json:"nick_name"`
	IdCard            string `json:"id_card"`
	License           string `json:"license"`
	NamingOtherStuff1 string `json:"naming_other_stuff_1"`
	NamingOtherStuff2 string `json:"naming_other_stuff_2"`
	NamingOtherStuff3 string `json:"naming_other_stuff_3"`
	NamingOtherStuff4 string `json:"naming_other_stuff_4"`
	NamingOtherStuff5 string `json:"naming_other_stuff_5"`
}

type SetNickNameRes struct {
	Wording string `json:"wording"`
	AuditId int    `json:"audit_id"`
}

type GetNickNameStatusRes struct {
	Nickname   string `json:"nickname"`
	AuditStat  int    `json:"audit_stat"`
	FailReason string `json:"fail_reason"`
	CreateTime int    `json:"create_time"`
	AuditTime  int    `json:"audit_time"`
}

type GetSearchStatusRes struct {
	Status int `json:"status"`
}

type GetFetchDataSetting struct {
	IsPreFetchOpen     bool   `json:"is_pre_fetch_open"`
	PreFetchType       int    `json:"pre_fetch_type"`
	PreFetchUrl        string `json:"pre_fetch_url"`
	PreEnv             string `json:"pre_env"`
	PreFunctionName    string `json:"pre_function_name"`
	IsPeriodFetchOpen  bool   `json:"is_period_fetch_open"`
	PeriodFetchType    int    `json:"period_fetch_type"`
	PeriodFetchUrl     string `json:"period_fetch_url"`
	PeriodEnv          string `json:"period_env"`
	PeriodFunctionName string `json:"period_function_name"`
}

type GetFetchDataSettingParams struct {
	Action string `json:"action"`
	GetFetchDataSetting
}

type SetHeadImageParams struct {
	HeadImgMediaId string `json:"head_img_media_id"`
	X1             string `json:"x1"`
	Y1             string `json:"y1"`
	X2             string `json:"x2"`
	Y2             string `json:"y2"`
}
