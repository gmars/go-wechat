package authorization

import (
	"context"
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/util"
)

type Authorization struct {
	request     *core.ApiRequest
	accessToken core.AccessToken
}

func NewAuthorization(a core.AccessToken) *Authorization {
	return &Authorization{request: core.NewApiRequest(a), accessToken: a}
}

// CreatePreAuthCode 创建预授权码
func (a *Authorization) CreatePreAuthCode() (*PreAuthCodeRes, error) {
	var res PreAuthCodeRes
	_, err := a.request.JsonPost("/cgi-bin/component/api_create_preauthcode", nil, map[string]string{
		"component_appid": a.accessToken.GetCurrentAppid(context.Background()),
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryAuth 使用授权码获取授权信息
func (a *Authorization) QueryAuth(authorizationCode string) (*AuthorizationInfo, error) {
	var res AuthQueryRes
	_, err := a.request.JsonPost("/cgi-bin/component/api_query_auth", nil, map[string]string{
		"component_appid":    a.accessToken.GetCurrentAppid(context.Background()),
		"authorization_code": authorizationCode,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res.AuthorizationInfo, nil
}

// GetAuthorizerList 获取已授权的账号基本信息
func (a *Authorization) GetAuthorizerList(page, pageSize int) (*AuthorizerInfoList, error) {
	var res AuthorizerInfoList
	offset, count := util.PageCondition(page, pageSize, 500)
	_, err := a.request.JsonPost("/cgi-bin/component/api_get_authorizer_list", nil, map[string]interface{}{
		"component_appid": a.accessToken.GetCurrentAppid(context.Background()),
		"offset":          offset,
		"count":           count,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// AuthorizerInfo 获取授权方的基本信息
func (a *Authorization) AuthorizerInfo(authorizerAppid string) (*AuthorizerInfo, error) {
	var res AuthorizerInfo
	_, err := a.request.JsonPost("/cgi-bin/component/api_get_authorizer_info", nil, map[string]interface{}{
		"component_appid":  a.accessToken.GetCurrentAppid(context.Background()),
		"authorizer_appid": authorizerAppid,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// AuthorizerOption 获取授权方选项
func (a *Authorization) AuthorizerOption(authorizerAppid, optionName string) (*AuthorizerOption, error) {
	var res AuthorizerOption
	_, err := a.request.JsonPost("/cgi-bin/component/api_get_authorizer_option", nil, map[string]interface{}{
		"component_appid":  a.accessToken.GetCurrentAppid(context.Background()),
		"authorizer_appid": authorizerAppid,
		"option_name":      optionName,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetAuthorizerOption 设置授权方选项信息
func (a *Authorization) SetAuthorizerOption(authorizerAppid, optionName, optionValue string) error {
	_, err := a.request.JsonPost("/cgi-bin/component/api_set_authorizer_option", nil, map[string]interface{}{
		"component_appid":  a.accessToken.GetCurrentAppid(context.Background()),
		"authorizer_appid": authorizerAppid,
		"option_name":      optionName,
		"option_value":     optionValue,
	}, nil)
	return err
}
