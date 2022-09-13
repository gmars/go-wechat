package domain

type ActionType = string

const (
	ActionAdd    ActionType = "add"
	ActionDelete ActionType = "delete"
	ActionSet    ActionType = "set"
	ActionGet    ActionType = "get"
)

type ModifyServerDomainParams struct {
	Action          ActionType `json:"action"`
	RequestDomain   []string   `json:"requestdomain"`
	WsRequestDomain []string   `json:"wsrequestdomain"`
	UploadDomain    []string   `json:"uploaddomain"`
	DownloadDomain  []string   `json:"downloaddomain"`
	UdpDomain       []string   `json:"udpdomain"`
	TcpDomain       []string   `json:"tcpdomain"`
}

type ModifyServerDomainRes struct {
	RequestDomain          []string `json:"requestdomain"`
	WsRequestDomain        []string `json:"wsrequestdomain"`
	UploadDomain           []string `json:"uploaddomain"`
	DownloadDomain         []string `json:"downloaddomain"`
	UdpDomain              []string `json:"udpdomain"`
	TcpDomain              []string `json:"tcpdomain"`
	InvalidRequestDomain   []string `json:"invalid_requestdomain"`
	InvalidWsRequestDomain []string `json:"invalid_wsrequestdomain"`
	InvalidUploadDomain    []string `json:"invalid_uploaddomain"`
	InvalidDownloadDomain  []string `json:"invalid_downloaddomain"`
	InvalidUdpDomain       []string `json:"invalid_udpdomain"`
	InvalidTcpDomain       []string `json:"invalid_tcpdomain"`
	NoIcpDomain            []string `json:"no_icp_domain"`
}

type GetJumpDomainConfirmFileRes struct {
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

type ModifyJumpDomainDirectlyRes struct {
	WebViewDomain []string `json:"webviewdomain"`
}

type GetEffectiveServerDomainRes struct {
	MpDomain struct {
		RequestDomain   []string `json:"requestdomain"`
		WsRequestDomain []string `json:"wsrequestdomain"`
		UploadDomain    []string `json:"uploaddomain"`
		DownloadDomain  []string `json:"downloaddomain"`
		UdpDomain       []string `json:"udpdomain"`
		TcpDomain       []string `json:"tcpdomain"`
	} `json:"mp_domain"`
	ThirdDomain struct {
		RequestDomain   []string `json:"requestdomain"`
		WsRequestDomain []string `json:"wsrequestdomain"`
		UploadDomain    []string `json:"uploaddomain"`
		DownloadDomain  []string `json:"downloaddomain"`
		UdpDomain       []string `json:"udpdomain"`
		TcpDomain       []string `json:"tcpdomain"`
	} `json:"third_domain"`
	DirectDomain struct {
		RequestDomain   []string `json:"requestdomain"`
		WsRequestDomain []string `json:"wsrequestdomain"`
		UploadDomain    []string `json:"uploaddomain"`
		DownloadDomain  []string `json:"downloaddomain"`
		UdpDomain       []string `json:"udpdomain"`
		TcpDomain       []string `json:"tcpdomain"`
	} `json:"direct_domain"`
	EffectiveDomain struct {
		RequestDomain   []string `json:"requestdomain"`
		WsRequestDomain []string `json:"wsrequestdomain"`
		UploadDomain    []string `json:"uploaddomain"`
		DownloadDomain  []string `json:"downloaddomain"`
		UdpDomain       []string `json:"udpdomain"`
		TcpDomain       []string `json:"tcpdomain"`
	} `json:"effective_domain"`
}

type GetEffectiveJumpDomainRes struct {
	MpWebviewDomain        []string `json:"mp_webviewdomain"`
	ThirdWebviewDomain     []string `json:"third_webviewdomain"`
	DirectWebviewDomain    []string `json:"direct_webviewdomain"`
	EffectiveWebviewDomain []string `json:"effective_webviewdomain"`
}

type GetPrefetchDomainRes struct {
	PrefetchDnsDomain []struct {
		Url    string `json:"url"`
		Status int    `json:"status"`
	} `json:"prefetch_dns_domain"`
	SizeLimit int `json:"size_limit"`
}
