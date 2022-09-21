package website

import (
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/official/message"
	"net/url"
)

type WebSite struct {
	appId     string
	appSecret string
	baseUrl   string
	Message   *message.Message
}

func NewWebSite(appId, appSecret, token, aesKey, baseUrl string) *WebSite {
	webSite := &WebSite{
		appId:     appId,
		appSecret: appSecret,
		baseUrl:   baseUrl,
	}

	if token != "" && aesKey != "" {
		webSite.Message = message.NewMessageHandler(aesKey, token, appId)
	}
	return webSite
}

// GetLoginRedirectUrl 获取网站登录url
func (s *WebSite) GetLoginRedirectUrl(encodePath string, state string, lang LangType) (string, error) {
	snsUrl, err := url.Parse("https://open.weixin.qq.com/connect/qrconnect")
	if err != nil {
		return "", err
	}
	query := snsUrl.Query()
	query.Add("appid", s.appId)
	query.Add("redirect_uri", s.baseUrl+encodePath)
	query.Add("response_type", "code")
	query.Add("scope", "snsapi_login")
	query.Add("state", state)
	query.Add("lang", lang)
	snsUrl.RawQuery = query.Encode()
	return snsUrl.String(), nil
}

// GetLoginParams 获取非跳转登录的配置对象
// 页面引入http://res.wx.qq.com/connect/zh_CN/htmledition/js/wxLogin.js
func (s *WebSite) GetLoginParams(encodePath string, state string, style StyleType) (*LoginConfigRes, error) {
	return &LoginConfigRes{
		SelfRedirect: false,
		Id:           "",
		AppId:        s.appId,
		Scope:        "snsapi_login",
		RedirectUri:  s.baseUrl + encodePath,
		State:        state,
		Style:        style,
		Href:         "",
	}, nil
}

// GetAccessTokenWithCode 使用code换取access token
func (s *WebSite) GetAccessTokenWithCode(code string) (*AccessTokenRes, error) {
	var res AccessTokenRes
	request := core.NewApiRequest(nil)
	_, err := request.JsonGet("/sns/oauth2/access_token", map[string]string{
		"appid":      s.appId,
		"secret":     s.appSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}, &res)
	return &res, err
}

// GetUserInfo 获取用户信息
func (s *WebSite) GetUserInfo(accessToken, openId string, lang LangType) (*UserInfoRes, error) {
	var res UserInfoRes
	request := core.NewApiRequest(nil)
	_, err := request.JsonGet("/sns/userinfo", map[string]string{
		"access_token": accessToken,
		"openid":       openId,
		"lang":         lang,
	}, &res)
	return &res, err
}
