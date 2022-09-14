package basic_info

import "github.com/gmars/go-wechat/core"

type BasicInfo struct {
	request *core.ApiRequest
}

func NewBasicInfoManager(authorizerAccessToken core.AccessToken) *BasicInfo {
	return &BasicInfo{request: core.NewApiRequest(authorizerAccessToken)}
}

// GetAccountBasicInfo 获取基本信息
func (s *BasicInfo) GetAccountBasicInfo() (*AccountBasicInfoRes, error) {
	var res AccountBasicInfoRes
	_, err := s.request.JsonPost("/cgi-bin/account/getaccountbasicinfo", nil, nil, &res)
	return &res, err
}

// GetBindOpenAccount 查询绑定的开放平台帐号
func (s *BasicInfo) GetBindOpenAccount() (bool, error) {
	var res HaveOpenRes
	_, err := s.request.JsonGet("/cgi-bin/open/have", nil, &res)
	return res.HaveOpen, err
}

// CheckNickName 小程序名称检测
func (s *BasicInfo) CheckNickName(nickName string) (*CheckNickNameRes, error) {
	var res CheckNickNameRes
	_, err := s.request.JsonPost("/cgi-bin/wxverify/checkwxverifynickname", nil, map[string]string{
		"nick_name": nickName,
	}, &res)
	return &res, err
}

// SetNickName 设置小程序名称
func (s *BasicInfo) SetNickName(params *SetNickNameParams) (*SetNickNameRes, error) {
	var res SetNickNameRes
	_, err := s.request.JsonPost("/wxa/setnickname", nil, params, &res)
	return &res, err
}

// GetNickNameStatus 查询小程序名称审核状态
func (s *BasicInfo) GetNickNameStatus(auditId int) (*GetNickNameStatusRes, error) {
	var res GetNickNameStatusRes
	_, err := s.request.JsonPost("/wxa/api_wxa_querynickname", nil, map[string]int{
		"audit_id": auditId,
	}, &res)
	return &res, err
}

// SetSignature 设置小程序介绍
func (s *BasicInfo) SetSignature(signature string) error {
	var res GetNickNameStatusRes
	_, err := s.request.JsonPost("/cgi-bin/account/modifysignature", nil, map[string]string{
		"signature": signature,
	}, &res)
	return err
}

// GetSearchStatus 获取搜索状态
func (s *BasicInfo) GetSearchStatus() (bool, error) {
	var res GetSearchStatusRes
	_, err := s.request.JsonGet("/wxa/getwxasearchstatus", nil, &res)
	return res.Status == 0, err
}

// SetSearchStatus 设置搜索状态
func (s *BasicInfo) SetSearchStatus(status int) error {
	var res GetNickNameStatusRes
	_, err := s.request.JsonPost("/cgi-bin/account/changewxasearchstatus", nil, map[string]int{
		"status": status,
	}, &res)
	return err
}

// GetFetchDataSetting 数据预拉取和数据周期性更新
func (s *BasicInfo) GetFetchDataSetting(params *GetFetchDataSettingParams) (*GetFetchDataSetting, error) {
	var res GetFetchDataSetting
	_, err := s.request.JsonPost("/cgi-bin/account/fetchdatasetting", nil, params, &res)
	return &res, err
}

// SetHeadImage 修改头像
func (s *BasicInfo) SetHeadImage(params *SetHeadImageParams) error {
	_, err := s.request.JsonPost("/cgi-bin/account/fetchdatasetting", nil, params, nil)
	return err
}
