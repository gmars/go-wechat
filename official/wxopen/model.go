package wxopen

type SaveJumpQrCodeParams struct {
	Prefix string `json:"prefix"`
	Appid  string `json:"appid"`
	Path   string `json:"path"`
	IsEdit uint32 `json:"is_edit"`
}

type GetJumpQrCodeParams struct {
	Appid      string   `json:"appid"`
	GetType    uint     `json:"get_type"`
	PrefixList []string `json:"prefix_list"`
	PageNum    uint     `json:"page_num"`
	PageSize   uint     `json:"page_size"`
}

type JumpQrCodeListRes struct {
	RuleList []struct {
		Prefix string `json:"prefix"`
		State  uint   `json:"state"`
		Path   string `json:"path"`
	} `json:"rule_list"`
	QrcodeJumpOpen     uint `json:"qrcodejump_open"`
	ListSize           uint `json:"list_size"`
	QrcodeJumpPubQuota uint `json:"qrcodejump_pub_quota"`
	TotalCount         uint `json:"total_count"`
}

type WxAmpLinkItem struct {
	Status              int    `json:"status"`
	Username            string `json:"username"`
	Appid               string `json:"appid"`
	Source              string `json:"source"`
	Nickname            string `json:"nickname"`
	Selected            int    `json:"selected"`
	NearbyDisplayStatus int    `json:"nearby_display_status"`
	Released            int    `json:"released"`
	HeadImgUrl          string `json:"headimg_url"`
	FuncInfos           []struct {
		Status int    `json:"status"`
		Id     int    `json:"id"`
		Name   string `json:"name"`
	} `json:"func_infos"`
	CopyVerifyStatus int    `json:"copy_verify_status"`
	Email            string `json:"email"`
}

type WxAmpLinkRes struct {
	WxOpens struct {
		Items []WxAmpLinkItem `json:"items"`
	} `json:"wxopens"`
}
