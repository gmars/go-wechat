package authorization

type PreAuthCodeRes struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

type FuncScopeCategory struct {
	Id   int    `json:"id"`
	Type int    `json:"type"`
	Name string `json:"name"`
	Desc string `json:"desC"`
}

type AuthorizationInfo struct {
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	FuncInfo               []struct {
		FuncScopeCategory FuncScopeCategory `json:"funcscope_category"`
	} `json:"func_info"`
}

type MiniProgramInfo struct {
	Network struct {
		RequestDomain   []string `json:"RequestDomain"`
		WsRequestDomain []string `json:"WsRequestDomain"`
		UploadDomain    []string `json:"UploadDomain"`
		DownloadDomain  []string `json:"DownloadDomain"`
		UDPDomain       []string `json:"UDPDomain"`
		TCPDomain       []string `json:"TCPDomain"`
	} `json:"network"`
	Categories []struct {
		First  string `json:"first"`
		Second string `json:"second"`
	} `json:"categories"`
	VisitStatus int `json:"visit_status"`
}

type AuthorizerInfo struct {
	NickName        string `json:"nick_name"`
	HeadImg         string `json:"head_img"`
	ServiceTypeInfo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"service_type_info"`
	VerifyTypeInfo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"verify_type_info"`
	UserName     string `json:"user_name"`
	Alias        string `json:"alias"`
	QrcodeUrl    string `json:"qrcode_url"`
	BusinessInfo struct {
		OpenStore int `json:"open_store"`
		OpenScan  int `json:"open_scan"`
		OpenPay   int `json:"open_pay"`
		OpenCard  int `json:"open_card"`
		OpenShake int `json:"open_shake"`
	} `json:"business_info"`
	PrincipalName   string          `json:"principal_name"`
	Signature       string          `json:"signature"`
	MiniProgramInfo MiniProgramInfo `json:"MiniProgramInfo"`
	RegisterType    int             `json:"register_type"`
	AccountStatus   int             `json:"account_status"`
	BasicConfig     struct {
		IsPhoneConfigured bool `json:"is_phone_configured"`
		IsEmailConfigured bool `json:"is_email_configured"`
	} `json:"basic_config"`
}

type AuthQueryRes struct {
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}

type AuthorizerInfoList struct {
	TotalCount int `json:"total_count"`
	List       []struct {
		AuthorizerAppid string `json:"authorizer_appid"`
		RefreshToken    string `json:"refresh_token"`
		AuthTime        int    `json:"auth_time"`
	} `json:"list"`
}

type AuthorizerOption struct {
	AuthorizerAppid string `json:"authorizer_appid"`
	OptionName      string `json:"option_name"`
	OptionValue     string `json:"option_value"`
}

type AuthorizerRes struct {
	AuthorizerInfo    AuthorizerInfo    `json:"authorizer_info"`
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}
