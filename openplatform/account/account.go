package account

import "github.com/gmars/go-wechat/core"

type Account struct {
	request *core.ApiRequest
}

func NewOpenAccount(authorizerAccessToken core.AccessToken) *Account {
	return &Account{request: core.NewApiRequest(authorizerAccessToken)}
}

// CreateOpenApiBind 创建开放平台账号并和appid对应的小程序/公众号绑定
func (a *Account) CreateOpenApiBind(appid string) (string, error) {
	var res CreateOpenAccountBindRes
	_, err := a.request.JsonPost("/cgi-bin/open/create", nil, map[string]string{
		"appid": appid,
	}, &res)
	return res.OpenAppid, err
}

// Bind 绑定开放平台账号到appid对应的小程序/公众号绑定
func (a *Account) Bind(openAppid, appId string) error {
	_, err := a.request.JsonPost("/cgi-bin/open/bind", nil, map[string]string{
		"appid":      appId,
		"open_appid": openAppid,
	}, nil)
	return err
}

// Unbind 解除绑定开放平台账号到appid对应的小程序/公众号绑定
func (a *Account) Unbind(openAppid, appId string) error {
	_, err := a.request.JsonPost("/cgi-bin/open/unbind", nil, map[string]string{
		"appid":      appId,
		"open_appid": openAppid,
	}, nil)
	return err
}

// GetOpenAccount 获取开放平台账号appid
func (a *Account) GetOpenAccount(appid string) (string, error) {
	var res CreateOpenAccountBindRes
	_, err := a.request.JsonPost("/cgi-bin/open/get", nil, map[string]string{
		"appid": appid,
	}, &res)
	return res.OpenAppid, err
}
