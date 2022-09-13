package customerservice

type KfInfo struct {
	KfAccount        string `json:"kf_account"`
	KfHeadImgUrl     string `json:"kf_headimgurl"`
	KfId             int    `json:"kf_id"`
	KfNick           string `json:"kf_nick"`
	InviteWx         string `json:"invite_wx"`
	InviteExpireTime int    `json:"invite_expire_time"`
	InviteStatus     string `json:"invite_status"`
}

type KfListRes struct {
	KfList []KfInfo `json:"kf_list"`
}

type KfOnlineInfo struct {
	KfAccount    string `json:"kf_account"`
	Status       int    `json:"status"`
	KfId         string `json:"kf_id"`
	AcceptedCase int    `json:"accepted_case"`
}

type OnlineKfRes struct {
	KfOnlineList []KfOnlineInfo `json:"kf_online_list"`
}

type DialogStatus struct {
	CreateTime int    `json:"createtime"`
	KfAccount  string `json:"kf_account"`
}

type KfDialogItem struct {
	CreateTime int    `json:"createtime"`
	Openid     string `json:"openid"`
}

type KfDialogRecords struct {
	SessionList []KfDialogItem `json:"sessionlist"`
}

type WaitingRes struct {
	Count        int `json:"count"`
	WaitCaseList []struct {
		LatestTime int    `json:"latest_time"`
		Openid     string `json:"openid"`
	} `json:"waitcaselist"`
}

type ChartRecordsParams struct {
	StartTime int `json:"starttime"`
	EndTime   int `json:"endtime"`
	MsgId     int `json:"msgid"`
	Number    int `json:"number"`
}

type RecordItem struct {
	Openid   string `json:"openid"`
	Opercode int    `json:"opercode"`
	Text     string `json:"text"`
	Time     int    `json:"time"`
	Worker   string `json:"worker"`
}

type RecordsList struct {
	RecordList []RecordItem `json:"recordlist"`
	Number     int          `json:"number"`
	MsgId      int          `json:"msgid"`
}

type VideoMsgParams struct {
	MediaId      string `json:"media_id"`
	ThumbMediaId string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

type MusicMsgParams struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HqMusicUrl   string `json:"hqmusicurl"`
	ThumbMediaId string `json:"thumb_media_id"`
}

type LinkArticlesParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

type MenuMsgListItem struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}
