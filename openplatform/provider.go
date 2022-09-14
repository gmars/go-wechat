package openplatform

import (
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/official/message"
	"github.com/gmars/go-wechat/openplatform/account"
	"github.com/gmars/go-wechat/openplatform/authorization"
	"github.com/gmars/go-wechat/openplatform/authorizer_token"
	"github.com/gmars/go-wechat/openplatform/component"
	"github.com/gmars/go-wechat/openplatform/miniprogram/login"
	"github.com/gmars/go-wechat/openplatform/miniprogram/management"
	"github.com/gmars/go-wechat/openplatform/thirdparty_management"
)

// Authorization 授权账号管理
func (c *Client) Authorization() *authorization.Authorization {
	return authorization.NewAuthorization(c.componentAccessToken)
}

// MessageHandler 消息处理器
func (c *Client) MessageHandler() *message.Message {
	return c.messageHandler
}

// NotifyHandler 通知消息处理
func (c *Client) NotifyHandler() *component.Notify {
	return component.NewAuthorizationNotify(c.messageHandler, c.cache)
}

// AuthorizerAccessToken 授权方接口调用凭证管理 实现了AccessToken接口
func (c *Client) AuthorizerAccessToken(appId string, refreshToken string, callBack func(string, *authorizer_token.AuthorizerTokenRes)) *authorizer_token.AuthorizerToken {
	return authorizer_token.NewAuthorizerToken(appId, refreshToken, c.componentAccessToken, c.cache, callBack)
}

// OpenAccount 开放平台账号管理，绑定等操作
func (c *Client) OpenAccount(authorizerToken core.AccessToken) *account.Account {
	return account.NewOpenAccount(authorizerToken)
}

// ThirdPartyManagement 第三方平台管理[小程序相关]
func (c *Client) ThirdPartyManagement() *thirdparty_management.ThirdPartyManagement {
	return thirdparty_management.NewThirdPartyManagement(c.componentAccessToken)
}

// MiniProgramLogin 代小程序实现登录
func (c *Client) MiniProgramLogin() *login.Login {
	return login.NewMiniProgramLogin(c.componentAccessToken)
}

// MiniProgramManagement 代商家管理小程序
func (c *Client) MiniProgramManagement(authorizerToken core.AccessToken) *management.Management {
	return management.NewManagement(authorizerToken)
}
