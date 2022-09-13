package thirdparty_management

import "go-wechat/core"

type ThirdPartyManagement struct {
	request *core.ApiRequest
}

func NewThirdPartyManagement(componentAccessToken core.AccessToken) *ThirdPartyManagement {
	return &ThirdPartyManagement{request: core.NewApiRequest(componentAccessToken)}
}

// GetTemplatedRaftList 获取草稿箱列表
func (s *ThirdPartyManagement) GetTemplatedRaftList() (*TemplatedDraftList, error) {
	var res TemplatedDraftList
	_, err := s.request.JsonGet("/wxa/gettemplatedraftlist", nil, &res)
	return &res, err
}

// AddToTemplate 将草稿添加到模板库
func (s *ThirdPartyManagement) AddToTemplate(draftId, templateType int) error {
	_, err := s.request.JsonPost("/wxa/addtotemplate", nil, map[string]int{
		"draft_id":      draftId,
		"template_type": templateType,
	}, nil)
	return err
}

// GetTemplateList 获取模板列表
func (s *ThirdPartyManagement) GetTemplateList() ([]TemplateItem, error) {
	var res TemplateListRes
	_, err := s.request.JsonGet("/wxa/gettemplatelist", nil, &res)
	return res.TemplateList, err
}

// DeleteTemplate 删除代码模板
func (s *ThirdPartyManagement) DeleteTemplate(templateId int) error {
	_, err := s.request.JsonPost("/wxa/deletetemplate", nil, map[string]int{
		"template_id": templateId,
	}, nil)
	return err
}

// GetServerDomain 返回测试版和全网发布版的“小程序服务器域名”值
func (s *ThirdPartyManagement) GetServerDomain() (*ServerDomainRes, error) {
	return s.modifyThirdPartyServerDomain("get", "", false)
}

// SetServerDomain 覆盖小程序服务器域名
func (s *ThirdPartyManagement) SetServerDomain(wxaServerDomain string, isModifyPublishedTogether bool) (*ServerDomainRes, error) {
	return s.modifyThirdPartyServerDomain("set", wxaServerDomain, isModifyPublishedTogether)
}

// AddServerDomain 添加小程序服务器域名
func (s *ThirdPartyManagement) AddServerDomain(wxaServerDomain string, isModifyPublishedTogether bool) (*ServerDomainRes, error) {
	return s.modifyThirdPartyServerDomain("add", wxaServerDomain, isModifyPublishedTogether)
}

// DeleteServerDomain 删除小程序服务器域名
func (s *ThirdPartyManagement) DeleteServerDomain(wxaServerDomain string, isModifyPublishedTogether bool) (*ServerDomainRes, error) {
	return s.modifyThirdPartyServerDomain("delete", wxaServerDomain, isModifyPublishedTogether)
}

// GetJumpDomainConfirmFile 获取第三方平台业务域名校验文件
func (s *ThirdPartyManagement) GetJumpDomainConfirmFile() (*ConfirmFile, error) {
	var res ConfirmFile
	_, err := s.request.JsonPost("/cgi-bin/component/get_domain_confirmfile", nil, nil, &res)
	return &res, err
}

// GetJumpDomain 返回测试版和全网发布版的业务域名
func (s *ThirdPartyManagement) GetJumpDomain() (*JumpDomainRes, error) {
	return s.modifyThirdPartyJumpDomain("get", "", false)
}

// SetJumpDomain 覆盖小程序业务域名
func (s *ThirdPartyManagement) SetJumpDomain(wxaServerDomain string, isModifyPublishedTogether bool) (*JumpDomainRes, error) {
	return s.modifyThirdPartyJumpDomain("set", wxaServerDomain, isModifyPublishedTogether)
}

// AddJumpDomain 添加小程序业务域名
func (s *ThirdPartyManagement) AddJumpDomain(wxaServerDomain string, isModifyPublishedTogether bool) (*JumpDomainRes, error) {
	return s.modifyThirdPartyJumpDomain("add", wxaServerDomain, isModifyPublishedTogether)
}

// DeleteJumpDomain 删除小程序业务域名
func (s *ThirdPartyManagement) DeleteJumpDomain(wxaServerDomain string, isModifyPublishedTogether bool) (*JumpDomainRes, error) {
	return s.modifyThirdPartyJumpDomain("delete", wxaServerDomain, isModifyPublishedTogether)
}

// GetJumpDomainConfirmFile 获取第三方平台业务域名校验文件
func (s *ThirdPartyManagement) modifyThirdPartyJumpDomain(action, wxaJumpH5Domain string, isModifyPublishedTogether bool) (*JumpDomainRes, error) {
	var res JumpDomainRes
	_, err := s.request.JsonPost("/cgi-bin/component/modify_wxa_jump_domain", nil, map[string]interface{}{
		"action":                       action,
		"wxa_jump_h5_domain":           wxaJumpH5Domain,
		"is_modify_published_together": isModifyPublishedTogether,
	}, &res)
	return &res, err
}

func (s *ThirdPartyManagement) modifyThirdPartyServerDomain(action, wxaServerDomain string, isModifyPublishedTogether bool) (*ServerDomainRes, error) {
	var res ServerDomainRes
	_, err := s.request.JsonPost("/cgi-bin/component/modify_wxa_server_domain", nil, map[string]interface{}{
		"action":                       action,
		"wxa_server_domain":            wxaServerDomain,
		"is_modify_published_together": isModifyPublishedTogether,
	}, &res)
	return &res, err
}
