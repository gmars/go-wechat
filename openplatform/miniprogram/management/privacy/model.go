package privacy

type UploadPrivacySettingParams struct {
	OwnerSetting struct {
		ContactEmail         string `json:"contact_email"`
		ContactPhone         string `json:"contact_phone"`
		ContactQq            string `json:"contact_qq"`
		ContactWeixin        string `json:"contact_weixin"`
		ExtFileMediaId       string `json:"ext_file_media_id"`
		NoticeMethod         string `json:"notice_method"`
		StoreExpireTimestamp string `json:"store_expire_timestamp"`
	} `json:"owner_setting"`
	SettingList []struct {
		PrivacyKey  string `json:"privacy_key"`
		PrivacyText string `json:"privacy_text"`
	} `json:"setting_list"`
	SdkPrivacyInfoList []struct {
		SdkName    string `json:"sdk_name"`
		SdkBizName string `json:"sdk_biz_name"`
		SdkList    []struct {
			PrivacyKey  string `json:"privacy_key"`
			PrivacyText string `json:"privacy_text"`
		} `json:"sdk_list"`
	} `json:"sdk_privacy_info_list"`
	PrivacyVer int `json:"privacy_ver"`
}

type GetPrivacySettingRes struct {
	CodeExist   int      `json:"code_exist"`
	PrivacyList []string `json:"privacy_list"`
	SettingList []struct {
		PrivacyKey   string `json:"privacy_key"`
		PrivacyText  string `json:"privacy_text"`
		PrivacyLabel string `json:"privacy_label"`
	} `json:"setting_list"`
	UpdateTime   int `json:"update_time"`
	OwnerSetting struct {
		ContactPhone         string `json:"contact_phone"`
		ContactEmail         string `json:"contact_email"`
		ContactQq            string `json:"contact_qq"`
		ContactWeixin        string `json:"contact_weixin"`
		StoreExpireTimestamp string `json:"store_expire_timestamp"`
		ExtFileMediaId       string `json:"ext_file_media_id"`
		NoticeMethod         string `json:"notice_method"`
	} `json:"owner_setting"`
	PrivacyDesc struct {
		PrivacyDescList []struct {
			PrivacyKey  string `json:"privacy_key"`
			PrivacyDesc string `json:"privacy_desc"`
		} `json:"privacy_desc_list"`
	} `json:"privacy_desc"`
	SdkPrivacyInfoList []struct {
		SdkName    string `json:"sdk_name"`
		SdkBizName string `json:"sdk_biz_name"`
		SdkList    []struct {
			PrivacyKey   string `json:"privacy_key"`
			PrivacyText  string `json:"privacy_text"`
			PrivacyLabel string `json:"privacy_label"`
		} `json:"sdk_list"`
	} `json:"sdk_privacy_info_list"`
}

type UploadPrivacySettingRes struct {
	ExtFileMediaId string `json:"ext_file_media_id"`
}

type ApplyPrivacyInterfaceParams struct {
	ApiName   string   `json:"api_name"`
	Content   string   `json:"content"`
	PicList   []string `json:"pic_list"`
	VideoList []string `json:"video_list"`
	UrlList   []string `json:"url_list"`
}

type ApplyPrivacyInterfaceRes struct {
	AuditId int `json:"audit_id"`
}

type InterfaceItem struct {
	ApiName    string `json:"api_name"`
	ApiChName  string `json:"api_ch_name"`
	ApiDesc    string `json:"api_desc"`
	Status     int    `json:"status"`
	ApiLink    string `json:"api_link"`
	GroupName  string `json:"group_name"`
	AuditId    int    `json:"audit_id,omitempty"`
	FailReason string `json:"fail_reason,omitempty"`
}

type GetPrivacyInterfaceRes struct {
	InterfaceList []InterfaceItem `json:"interface_list"`
}
