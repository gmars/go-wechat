package domain

import "go-wechat/core"

type Domain struct {
	request *core.ApiRequest
}

func NewDomain(authorizerAccessToken core.AccessToken) *Domain {
	return &Domain{request: core.NewApiRequest(authorizerAccessToken)}
}

// ModifyServerDomain 配置小程序服务器域名
func (s *Domain) ModifyServerDomain(params *ModifyServerDomainParams) (*ModifyServerDomainRes, error) {
	var res ModifyServerDomainRes
	_, err := s.request.JsonPost("/wxa/modify_domain", nil, params, &res)
	return &res, err
}

// ModifyJumpDomain 配置小程序业务域名
func (s *Domain) ModifyJumpDomain(action ActionType, domains []string) error {
	_, err := s.request.JsonPost("/wxa/setwebviewdomain", nil, map[string]interface{}{
		"action":        action,
		"webviewdomain": domains,
	}, nil)
	return err
}

// ModifyServerDomainDirectly 快速配置小程序服务器域名[不需要先配置到第三方平台]
func (s *Domain) ModifyServerDomainDirectly(params *ModifyServerDomainParams) error {
	_, err := s.request.JsonPost("/wxa/setwebviewdomain", nil, params, nil)
	return err
}

// GetJumpDomainConfirmFile 获取业务域名校验文件
func (s *Domain) GetJumpDomainConfirmFile() (*GetJumpDomainConfirmFileRes, error) {
	var res GetJumpDomainConfirmFileRes
	_, err := s.request.JsonPost("/wxa/get_webviewdomain_confirmfile", nil, nil, &res)
	return &res, err
}

// ModifyJumpDomainDirectly 快速配置小程序业务域名
func (s *Domain) ModifyJumpDomainDirectly(action ActionType, domains []string) ([]string, error) {
	var res ModifyJumpDomainDirectlyRes
	_, err := s.request.JsonPost("/wxa/setwebviewdomain", nil, map[string]interface{}{
		"action":        action,
		"webviewdomain": domains,
	}, &res)
	return res.WebViewDomain, err
}

// GetEffectiveServerDomain 获取发布后生效服务器域名列表
func (s *Domain) GetEffectiveServerDomain() (*GetEffectiveServerDomainRes, error) {
	var res GetEffectiveServerDomainRes
	_, err := s.request.JsonPost("/wxa/get_effective_domain", nil, nil, &res)
	return &res, err
}

// GetEffectiveJumpDomain 获取发布后生效业务域名列表
func (s *Domain) GetEffectiveJumpDomain() (*GetEffectiveJumpDomainRes, error) {
	var res GetEffectiveJumpDomainRes
	_, err := s.request.JsonPost("/wxa/get_effective_webviewdomain", nil, nil, &res)
	return &res, err
}

// GetPrefetchDomain 获取 DNS 预解析域名
func (s *Domain) GetPrefetchDomain() (*GetPrefetchDomainRes, error) {
	var res GetPrefetchDomainRes
	_, err := s.request.JsonGet("/wxa/get_prefetchdnsdomain", nil, &res)
	return &res, err
}

// SetPrefetchDomain 设置 DNS 预解析域名
func (s *Domain) SetPrefetchDomain() (*GetPrefetchDomainRes, error) {
	var res GetPrefetchDomainRes
	_, err := s.request.JsonGet("/wxa/get_prefetchdnsdomain", nil, &res)
	return &res, err
}
