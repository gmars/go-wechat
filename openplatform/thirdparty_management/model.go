package thirdparty_management

type DraftItem struct {
	CreateTime  int    `json:"create_time"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
	DraftId     int    `json:"draft_id"`
}

type TemplatedDraftList struct {
	DraftList []DraftItem `json:"draft_list"`
}

type CategoryList struct {
	Address     string `json:"address"`
	Tag         string `json:"tag"`
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
	ThirdClass  string `json:"third_class"`
	Title       string `json:"title"`
	FirstId     int    `json:"first_id"`
	SecondId    int    `json:"second_id"`
	ThirdId     int    `json:"third_id"`
}

type TemplateItem struct {
	CreateTime             int          `json:"create_time"`
	UserVersion            string       `json:"user_version"`
	UserDesc               string       `json:"user_desc"`
	TemplateId             int          `json:"template_id"`
	SourceMiniProgramAppid string       `json:"source_miniprogram_appid"`
	SourceMiniProgram      string       `json:"source_miniprogram"`
	Developer              string       `json:"developer"`
	TemplateType           int          `json:"template_type"`
	CategoryList           CategoryList `json:"category_list"`
	AuditScene             int          `json:"audit_scene"`
	AuditStatus            int          `json:"audit_status"`
	Reason                 string       `json:"reason"`
}

type TemplateListRes struct {
	TemplateList []TemplateItem `json:"template_list"`
}

type ServerDomainRes struct {
	PublishedWxaServerDomain string `json:"published_wxa_server_domain"`
	TestingWxaServerDomain   string `json:"testing_wxa_server_domain"`
	InvalidWxaServerDomain   string `json:"invalid_wxa_server_domain"`
}

type ConfirmFile struct {
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

type JumpDomainRes struct {
	PublishedWxaJumpH5Domain string `json:"published_wxa_jump_h5_domain"`
	TestingWxaJumpH5Domain   string `json:"testing_wxa_jump_h5_domain"`
	InvalidWxaJumpH5Domain   string `json:"invalid_wxa_jump_h5_domain"`
}
