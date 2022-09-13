package security

import "go-wechat/core"

type Security struct {
	request *core.ApiRequest
}

func NewSecurityCenter(accessToken core.AccessToken) *Security {
	return &Security{request: core.NewApiRequest(accessToken)}
}

// MsgSecCheck 文本内容安全识别
func (s *Security) MsgSecCheck(params *MsgSecCheckParams) (*MsgSecCheckRes, error) {
	var res MsgSecCheckRes
	params.Version = SecVersion
	_, err := s.request.JsonPost("/wxa/msg_sec_check", nil, params, &res)
	return &res, err
}

// MediaCheckAsync 音视频内容安全识别
func (s *Security) MediaCheckAsync(params *MediaCheckAsyncParams) (string, error) {
	var res MediaCheckAsyncRes
	params.Version = SecVersion
	_, err := s.request.JsonPost("/wxa/media_check_async", nil, params, &res)
	return res.TraceId, err
}

// GetUserRiskRank 获取用户安全等级
func (s *Security) GetUserRiskRank(params *GetUserRiskRankParams) (*GetUserRiskRankRes, error) {
	var res GetUserRiskRankRes
	_, err := s.request.JsonPost("/wxa/getuserriskrank", nil, params, &res)
	return &res, err
}
