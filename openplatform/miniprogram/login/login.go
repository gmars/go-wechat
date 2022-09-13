package login

import (
	"context"
	"go-wechat/core"
)

type Login struct {
	request              *core.ApiRequest
	componentAccessToken core.AccessToken
}

func NewMiniProgramLogin(componentAccessToken core.AccessToken) *Login {
	return &Login{
		request:              core.NewApiRequest(componentAccessToken),
		componentAccessToken: componentAccessToken,
	}
}

func (s *Login) ThirdPartyCode2Session(authorizerAppId, jsCode string) (*ThirdPartyCode2SessionRes, error) {
	var res ThirdPartyCode2SessionRes
	_, err := s.request.JsonGet("/sns/component/jscode2session", map[string]string{
		"appid":           authorizerAppId,
		"grant_type":      "authorization_code",
		"component_appid": s.componentAccessToken.GetCurrentAppid(context.Background()),
		"js_code":         jsCode,
	}, &res)
	return &res, err
}
