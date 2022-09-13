package dev

type GetDomainInfoRes struct {
	RequestDomain   []string `json:"requestdomain"`
	WsRequestDomain []string `json:"wsrequestdomain"`
	UploadDomain    []string `json:"uploaddomain"`
	DownloadDomain  []string `json:"downloaddomain"`
	UdpDomain       []string `json:"udpdomain"`
	BizDomain       []string `json:"bizdomain"`
}

type GetPerformanceParams struct {
	CostTimeType     int    `json:"cost_time_type"`
	DefaultStartTime int    `json:"default_start_time"`
	DefaultEndTime   int    `json:"default_end_time"`
	Device           string `json:"device"`
	NetworkType      string `json:"networktype"`
	Scene            string `json:"scene"`
	IsDownloadCode   string `json:"is_download_code"`
}

type TimeData struct {
	RefDate      string `json:"ref_date"`
	CostTimeType int    `json:"cost_time_type"`
	CostTime     int    `json:"cost_time"`
}

type GetPerformanceRes struct {
	DefaultTimeData struct {
		List []TimeData `json:"list"`
	} `json:"default_time_data"`
	CompareTimeData struct {
		List []TimeData `json:"list"`
	} `json:"compare_time_data"`
}

type GetSceneListRes struct {
	Scene []struct {
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	} `json:"scene"`
}

type GetVersionListRes struct {
	CvList []struct {
		Type              int      `json:"type"`
		ClientVersionList []string `json:"client_version_list"`
	} `json:"cvlist"`
}

type RealtimeLogSearchParams struct {
	Date      string `json:"date"`
	BeginTime int    `json:"begintime"`
	EndTime   int    `json:"endtime"`
	Start     int    `json:"start"`
	Limit     int    `json:"limit"`
	TraceId   string `json:"traceId"`
	Url       string `json:"url"`
	Id        string `json:"id"`
	FilterMsg string `json:"filter_msg"`
	Level     int    `json:"level"`
}

type RealtimeLogItem struct {
	Level          int    `json:"level"`
	Platform       int    `json:"platform"`
	LibraryVersion string `json:"libraryVersion"`
	ClientVersion  string `json:"clientVersion"`
	Id             string `json:"id"`
	Timestamp      int    `json:"timestamp"`
	Msg            []struct {
		Time  int      `json:"time"`
		Msg   []string `json:"msg"`
		Level int      `json:"level"`
	} `json:"msg"`
	Url       string `json:"url"`
	TraceId   string `json:"traceid"`
	FilterMsg string `json:"filterMsg"`
}

type RealtimeLogSearchRes struct {
	Data struct {
		List  []RealtimeLogItem `json:"list"`
		Total int               `json:"total"`
	} `json:"data"`
}

type GetFeedbackParams struct {
	Type int `json:"begintime"`
	Page int `json:"endtime"`
	Num  int `json:"start"`
}

type GetFeedbackItem struct {
	RecordId   int      `json:"record_id"`
	CreateTime int      `json:"create_time"`
	Content    string   `json:"content"`
	Phone      int64    `json:"phone"`
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	HeadUrl    string   `json:"head_url"`
	Type       int      `json:"type"`
	MediaIds   []string `json:"mediaIds"`
	SystemInfo string   `json:"systemInfo"`
}
type GetFeedbackRes struct {
	List     []GetFeedbackItem `json:"list"`
	TotalNum int               `json:"total_num"`
}

type GetJsErrDetailParams struct {
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	ErrorMsgMd5   string `json:"errorMsgMd5"`
	ErrorStackMd5 string `json:"errorStackMd5"`
	AppVersion    string `json:"appVersion"`
	SdkVersion    string `json:"sdkVersion"`
	OsName        string `json:"osName"`
	ClientVersion string `json:"clientVersion"`
	Openid        string `json:"openid"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	Desc          string `json:"desc"`
}

type GetJsErrDetailItem struct {
	Count         string `json:"Count"`
	SdkVersion    string `json:"sdkVersion"`
	ClientVersion string `json:"ClientVersion"`
	ErrorStackMd5 string `json:"errorStackMd5"`
	TimeStamp     string `json:"TimeStamp"`
	AppVersion    string `json:"appVersion"`
	ErrorMsgMd5   string `json:"errorMsgMd5"`
	ErrorMsg      string `json:"errorMsg"`
	ErrorStack    string `json:"errorStack"`
	Ds            string `json:"Ds"`
	OsName        string `json:"OsName"`
	OpenId        string `json:"openId"`
	PluginVersion string `json:"pluginversion"`
	AppId         string `json:"appId"`
	DeviceModel   string `json:"DeviceModel"`
	Source        string `json:"source"`
	Route         string `json:"route"`
	Uin           string `json:"Uin"`
	Nickname      string `json:"nickname"`
}

type GetJsErrDetailRes struct {
	Success    bool                 `json:"success"`
	Openid     string               `json:"openid"`
	Data       []GetJsErrDetailItem `json:"data"`
	TotalCount int                  `json:"totalCount"`
}

type GetJsErrListParams struct {
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	ErrType    string `json:"errType"`
	AppVersion string `json:"appVersion"`
	Openid     string `json:"openid"`
	Keyword    string `json:"keyword"`
	OrderBy    string `json:"orderby"`
	Desc       string `json:"desc"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
}

type GetJsErrListData struct {
	ErrorMsgMd5   string `json:"errorMsgMd5"`
	ErrorMsg      string `json:"errorMsg"`
	Uv            int    `json:"uv"`
	Pv            int    `json:"pv"`
	ErrorStackMd5 string `json:"errorStackMd5"`
	ErrorStack    string `json:"errorStack"`
	PvPercent     string `json:"pvPercent"`
	UvPercent     string `json:"uvPercent"`
}

type GetJsErrListRes struct {
	Success    bool               `json:"success"`
	Openid     string             `json:"openid"`
	Data       []GetJsErrListData `json:"data"`
	TotalCount int                `json:"totalCount"`
}

type GetGrayReleasePlanRes struct {
	GrayReleasePlan struct {
		Status                  int  `json:"status"`
		CreateTimestamp         int  `json:"create_timestamp"`
		GrayPercentage          int  `json:"gray_percentage"`
		SupportExperiencerFirst bool `json:"support_experiencer_first"`
		SupportDebugerFirst     bool `json:"support_debuger_first"`
	} `json:"gray_release_plan"`
}

// ServiceInfo 必服务标签
type ServiceInfo struct {
	ID    uint   `json:"id"`
	Type  uint8  `json:"type"`
	Name  string `json:"name"`
	AppID string `json:"appid"`
	Path  string `json:"path"`
}

// KFInfo 客服信息
type KFInfo struct {
	OpenKF    bool   `json:"open_kf"`
	KFHeading string `json:"kf_headimg"`
	KFName    string `json:"kf_name"`
}

// PicList 门店图片
type PicList struct {
	List []string `json:"list"`
}

// ServiceInfos 服务标签列表
type ServiceInfos struct {
	ServiceInfos []ServiceInfo `json:"service_infos"`
}

type AddNearbyPoiParams struct {
	PicList           PicList      `json:"pic_list"`           // 门店图片，最多9张，最少1张，上传门店图片如门店外景、环境设施、商品服务等，图片将展示在微信客户端的门店页。图片链接通过文档https://mpoi.weixin.qq.com/wiki?t=resource/res_main&id=mp1444738729中的《上传图文消息内的图片获取URL》接口获取。必填，文件格式为bmp、png、jpeg、jpg或gif，大小不超过5M pic_list是字符串，内容是一个json！
	ServiceInfos      ServiceInfos `json:"service_infos"`      // 必服务标签列表 选填，需要填写服务标签ID、APPID、对应服务落地页的path路径，详细字段格式见下方示例
	StoreName         string       `json:"store_name"`         // 门店名字 必填，门店名称需按照所选地理位置自动拉取腾讯地图门店名称，不可修改，如需修改请重现选择地图地点或重新创建地点
	Hour              string       `json:"hour"`               // 营业时间，格式11:11-12:12 必填
	Credential        string       `json:"credential"`         // 资质号 必填, 15位营业执照注册号或9位组织机构代码
	Address           string       `json:"address"`            // 地址 必填
	CompanyName       string       `json:"company_name"`       // 主体名字 必填
	QualificationList string       `json:"qualification_list"` // 证明材料 必填 如果company_name和该小程序主体不一致，需要填qualification_list，详细规则见附近的小程序使用指南-如何证明门店的经营主体跟公众号或小程序帐号主体相关http://kf.qq.com/faq/170401MbUnim17040122m2qY.html
	KFInfo            KFInfo       `json:"kf_info"`            // 客服信息 选填，可自定义服务头像与昵称，具体填写字段见下方示例kf_info pic_list是字符串，内容是一个json！
	PoiID             string       `json:"poi_id"`             // 如果创建新的门店，poi_id字段为空 如果更新门店，poi_id参数则填对应门店的poi_id 选填
}

type AddNearbyPoiRes struct {
	Data struct {
		AuditID           string `json:"audit_id"`           //	审核单 ID
		PoiID             string `json:"poi_id"`             //	附近地点 ID
		RelatedCredential string `json:"related_credential"` //	经营资质证件号
	} `json:"data"`
}

// PositionList 地点列表
type PositionList struct {
	LeftApplyNum uint `json:"left_apply_num"` // 剩余可添加地点个数
	MaxApplyNum  uint `json:"max_apply_num"`  // 最大可添加地点个数
	Data         struct {
		List []struct {
			PoiID                string `json:"poi_id"`                // 附近地点 ID
			QualificationAddress string `json:"qualification_address"` // 资质证件地址
			QualificationNum     string `json:"qualification_num"`     // 资质证件证件号
			AuditStatus          int    `json:"audit_status"`          // 地点审核状态
			DisplayStatus        int    `json:"display_status"`        // 地点展示在附近状态
			RefuseReason         string `json:"refuse_reason"`         // 审核失败原因，audit_status=4 时返回
		} `json:"poi_list"` // 地址列表
	} `json:"-"`
	RawData string `json:"data"` // 地址列表的 JSON 格式字符串
}

type PositionListRes struct {
	Data PositionList `json:"data"`
}
