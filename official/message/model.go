package message

type ReplyMsg struct {
	ReplayType string
	TypeValue  string
	Data       interface{}
}

type NewsItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicUrl      string `json:"pic_url"`
	Url         string `json:"url"`
}

type EncryptMsg struct {
	ToUserName string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
}

type TypeParse struct {
	MsgType  string `xml:"MsgType"`
	Event    string `xml:"Event"`
	InfoType string `xml:"InfoType"`
}

type BaseMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MsgId        int64  `xml:"MsgId"`
	MsgDataId    uint32 `xml:"MsgDataId"`
	Idx          uint32 `xml:"Content"`
}

type TextMsg struct {
	BaseMsg
	Content      string `xml:"Content"`
	BizMsgMenuId string `xml:"bizmsgmenuid"`
}

type ImageMsg struct {
	BaseMsg
	PicUrl  string `xml:"PicUrl"`
	MediaId string `xml:"MediaId"`
}

type VoiceMsg struct {
	BaseMsg
	Format      string `xml:"Format"`
	MediaId     string `xml:"MediaId"`
	Recognition string `xml:"Recognition"`
}

type VideoMsg struct {
	BaseMsg
	ThumbMediaId string `xml:"ThumbMediaId"`
	MediaId      string `xml:"MediaId"`
}

type ShortVideoMsg VideoMsg

type LocationMsg struct {
	BaseMsg
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     uint32  `xml:"Scale"`
	Label     uint32  `xml:"Label"`
}

type LinkMsg struct {
	BaseMsg
	Title       float64 `xml:"Title"`
	Description float64 `xml:"Description"`
	Url         uint32  `xml:"Url"`
}

type BaseEventMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
}

type SubscribeEvent struct {
	BaseEventMsg
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

type ScanEvent struct {
	BaseEventMsg
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

type LocationEvent struct {
	BaseEventMsg
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision float64 `xml:"Precision"`
}

type ClickEvent struct {
	BaseEventMsg
	EventKey string `xml:"EventKey"`
}

type ViewEvent struct {
	BaseEventMsg
	EventKey string `xml:"EventKey"`
	MenuId   int    `xml:"MenuId"`
}

type ScanCodePushEvent struct {
	BaseEventMsg
	EventKey     string `xml:"EventKey"`
	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`
}

type ScanCodeWaitMsgEvent ScanCodePushEvent

type PicSysPhotoEvent struct {
	BaseEventMsg
	EventKey     string `xml:"EventKey"`
	SendPicsInfo struct {
		Count   int `xml:"Count"`
		PicList struct {
			Item struct {
				PicMd5Sum []string `xml:"PicMd5Sum"`
			} `xml:"item"`
		} `xml:"PicList"`
	} `xml:"SendPicsInfo"`
}

type PicPhotoOrAlbumEvent PicSysPhotoEvent

type PicWeixinEvent PicSysPhotoEvent

type LocationSelectEvent struct {
	BaseEventMsg
	EventKey         string `xml:"EventKey"`
	SendLocationInfo struct {
		LocationX string `xml:"Location_X"`
		LocationY string `xml:"Location_Y"`
		Scale     string `xml:"Scale"`
		Label     string `xml:"Label"`
	} `xml:"SendLocationInfo"`
}

type ViewMiniProgramEvent struct {
	BaseEventMsg
	EventKey string `xml:"EventKey"`
	MenuId   int    `xml:"MenuId"`
}

type TemplateSendJobFinishEvent struct {
	BaseEventMsg
	MsgID  int64  `xml:"MsgID"`
	Status string `xml:"Status"`
}

type MassSendJobFinishEvent struct {
	BaseEventMsg
	MsgID                int64  `xml:"MsgID"`
	Status               string `xml:"Status"`
	TotalCount           int    `xml:"TotalCount"`
	FilterCount          int    `xml:"FilterCount"`
	SentCount            int    `xml:"SentCount"`
	ErrorCount           int    `xml:"ErrorCount"`
	CopyrightCheckResult struct {
		Count      int `xml:"Count"`
		ResultList struct {
			Item []struct {
				ArticleIdx            int    `xml:"ArticleIdx"`
				UserDeclareState      int    `xml:"UserDeclareState"`
				AuditState            int    `xml:"AuditState"`
				OriginalArticleUrl    string `xml:"OriginalArticleUrl"`
				OriginalArticleType   int    `xml:"OriginalArticleType"`
				CanReprint            int    `xml:"CanReprint"`
				NeedReplaceContent    int    `xml:"NeedReplaceContent"`
				NeedShowReprintSource int    `xml:"NeedShowReprintSource"`
			} `xml:"item"`
		} `xml:"ResultList"`
		CheckState int `xml:"CheckState"`
	} `xml:"CopyrightCheckResult"`
	ArticleUrlResult struct {
		Count      int `xml:"Count"`
		ResultList struct {
			Item []struct {
				ArticleIdx int    `xml:"ArticleIdx"`
				ArticleUrl string `xml:"ArticleUrl"`
			} `xml:"item"`
		} `xml:"ResultList"`
	} `xml:"ArticleUrlResult"`
}

type SubscribeMsgPopupEvent struct {
	BaseEventMsg
	SubscribeMsgPopupEvent struct {
		List []struct {
			TemplateId            string `xml:"TemplateId"`
			SubscribeStatusString string `xml:"SubscribeStatusString"`
			PopupScene            int    `xml:"PopupScene"`
		} `xml:"List"`
	} `xml:"SubscribeMsgPopupEvent"`
}

type SubscribeMsgChangeEvent struct {
	BaseEventMsg
	SubscribeMsgChangeEvent struct {
		List []struct {
			TemplateId            string `xml:"TemplateId"`
			SubscribeStatusString string `xml:"SubscribeStatusString"`
		} `xml:"List"`
	} `xml:"SubscribeMsgChangeEvent"`
}

type SubscribeMsgSendEvent struct {
	BaseEventMsg
	SubscribeMsgSentEvent struct {
		List []struct {
			TemplateId  string `xml:"TemplateId"`
			ErrorStatus string `xml:"ErrorStatus"`
			MsgID       int64  `xml:"MsgID"`
			ErrorCode   string `xml:"ErrorCode"`
		} `xml:"List"`
	} `xml:"SubscribeMsgSentEvent"`
}

type PublishJobFinishEvent struct {
	BaseEventMsg
	PublishEventInfo struct {
		PublishId     int64  `xml:"publish_id"`
		PublishStatus int    `xml:"publish_status"`
		ArticleId     string `xml:"article_id"`
		ArticleDetail struct {
			Count int `xml:"count"`
			Item  []struct {
				Idx        int    `xml:"idx"`
				ArticleUrl string `xml:"article_url"`
			} `xml:"item"`
		} `xml:"article_detail"`
		FailIdx []int `xml:"fail_idx"`
	} `xml:"PublishEventInfo"`
}

type QualificationVerifySuccessEvent struct {
	BaseEventMsg
	ExpiredTime int64 `xml:"ExpiredTime"`
}

type QualificationVerifyFailEvent struct {
	BaseEventMsg
	FailTime   int64  `xml:"FailTime"`
	FailReason string `xml:"FailReason"`
}

type NamingVerifySuccessEvent struct {
	BaseEventMsg
	ExpiredTime int64 `xml:"ExpiredTime"`
}

type NamingVerifyFailEvent struct {
	BaseEventMsg
	FailTime   int64  `xml:"FailTime"`
	FailReason string `xml:"FailReason"`
}

type AnnualRenewEvent struct {
	BaseEventMsg
	ExpiredTime int64 `xml:"ExpiredTime"`
}

type VerifyExpiredEvent struct {
	BaseEventMsg
	ExpiredTime int64 `xml:"ExpiredTime"`
}

type PoiCheckNotifyEvent struct {
	BaseEventMsg
	UniqId string `xml:"UniqId"`
	PoiId  string `xml:"PoiId"`
	Result string `xml:"Result"`
	Msg    string `xml:"msg"`
}

type UserAuthorizeInvoiceEvent struct {
	BaseEventMsg
	SuccOrderId    string `xml:"SuccOrderId"`
	FailOrderId    string `xml:"FailOrderId"`
	AuthorizeAppId string `xml:"AuthorizeAppId "`
	Source         string `xml:"Source"`
}

type UpdateInvoiceStatusEvent struct {
	BaseEventMsg
	Status string `xml:"Status"`
	CardId string `xml:"CardId"`
	Code   string `xml:"Code"`
}

type SubmitInvoiceTitleEvent struct {
	BaseEventMsg
	Title     string `xml:"title"`
	Phone     string `xml:"phone"`
	TaxNo     string `xml:"tax_no"`
	Addr      string `xml:"addr"`
	BankType  string `xml:"bank_type"`
	BankNo    string `xml:"bank_no"`
	Attach    string `xml:"attach"`
	TitleType string `xml:"title_type"`
}

type UserInfoModifiedEvent struct {
	BaseEventMsg
	OpenID     string `xml:"OpenID"`
	AppID      string `xml:"AppID"`
	RevokeInfo string `xml:"RevokeInfo"`
}

type WxaNicknameAuditEvent struct {
	BaseEventMsg
	Ret      int    `xml:"ret"`
	Nickname string `xml:"nickname"`
	Reason   string `xml:"reason"`
}

type InfoTypeBase struct {
	AppId      string `xml:"AppId"`
	CreateTime int64  `xml:"CreateTime"`
	InfoType   string `xml:"InfoType"`
}

type ComponentVerifyTicketInfo struct {
	InfoTypeBase
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket"`
}

type AuthorizedInfo struct {
	InfoTypeBase
	AuthorizerAppid              string `xml:"AuthorizerAppid"`
	AuthorizationCode            string `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime int64  `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string `xml:"PreAuthCode"`
}

type UnauthorizedInfo struct {
	InfoTypeBase
	AuthorizerAppid string `xml:"AuthorizerAppid"`
}

type UpdateAuthorizedInfo struct {
	InfoTypeBase
	AuthorizerAppid              string `xml:"AuthorizerAppid"`
	AuthorizationCode            string `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime int64  `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string `xml:"PreAuthCode"`
}

type NotifyThirdFasterRegisterInfo struct {
	InfoTypeBase
	Appid    string `xml:"appid"`
	Status   int    `xml:"status"`
	AuthCode string `xml:"auth_code"`
	Msg      string `xml:"msg"`
	Info     struct {
		WxUser         string `xml:"wxuser"`
		IdName         string `xml:"idname"`
		ComponentPhone string `xml:"component_phone"`
	} `xml:"info"`
}

type NotifyThirdFastVerifyBetaAppInfo struct {
	InfoTypeBase
	Appid    string `xml:"appid"`
	Status   int    `xml:"status"`
	AuthCode string `xml:"auth_code"`
	Msg      string `xml:"msg"`
	Info     struct {
		Name               string `xml:"name"`
		Code               string `xml:"code"`
		CodeType           int    `xml:"code_type"`
		LegalPersonaWechat string `xml:"legal_persona_wechat"`
		LegalPersonaName   string `xml:"legal_persona_name"`
		ComponentPhone     string `xml:"component_phone"`
		UniqueId           string `xml:"unique_id"`
	} `xml:"info"`
}

const (
	EventSubscribe                  = "subscribe"
	EventScan                       = "SCAN"
	EventLocation                   = "LOCATION"
	EventClick                      = "CLICK"
	EventView                       = "VIEW"
	EventScanCodePush               = "scancode_push"
	EventScanCodeWaitMsg            = "scancode_waitmsg"
	EventPicSysPhoto                = "pic_sysphoto"
	EventPicPhotoOrAlbum            = "pic_photo_or_album"
	EventPicWeixin                  = "pic_weixin"
	EventLocationSelect             = "location_select"
	EventViewMiniProgram            = "view_miniprogram"
	EventTemplateSendJobFinish      = "TEMPLATESENDJOBFINISH"
	EventMassSendJobFinish          = "MASSSENDJOBFINISH"
	EventSubscribeMsgPopup          = "subscribe_msg_popup_event"
	EventSubscribeMsgChange         = "subscribe_msg_change_event"
	EventSubscribeMsgSent           = "subscribe_msg_sent_event"
	EventPublishJobFinish           = "PUBLISHJOBFINISH"
	EventQualificationVerifySuccess = "qualification_verify_success"
	EventQualificationVerifyFail    = "qualification_verify_fail"
	EventNamingVerifySuccess        = "naming_verify_success"
	EventNamingVerifyFail           = "naming_verify_fail"
	EventAnnualRenew                = "annual_renew"
	EventVerifyExpired              = "verify_expired"
	EventPoiCheckNotify             = "poi_check_notify"
	EventUserAuthorizeInvoice       = "user_authorize_invoice"
	EventUpdateInvoiceStatus        = "update_invoice_status"
	EventSubmitInvoiceTitle         = "submit_invoice_title"
	EventUserInfoModified           = "user_info_modified"

	MsgText       = "text"
	MsgImage      = "image"
	MsgVoice      = "voice"
	MsgVideo      = "video"
	MsgShortVideo = "shortvideo"
	MsgLocation   = "location"
	MsgLink       = "link"
	MsgEvent      = "event"

	InfoTypeComponentVerifyTicket        = "component_verify_ticket"
	InfoTypeAuthorized                   = "authorized"
	InfoTypeUnauthorized                 = "unauthorized"
	InfoTypeUpdateAuthorized             = "updateauthorized"
	InfoTypeNotifyThirdFasterRegister    = "notify_third_fasteregister"
	InfoTypeNotifyThirdFastVerifyBetaApp = "notify_third_fastverifybetaapp"

	ReplyTypeInfo     = "info"
	ReplyTypeStandard = "standard"
	ReplyTypeEvent    = "event"
)
