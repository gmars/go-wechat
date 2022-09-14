package miniapp

import (
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/miniapp/dev"
	"github.com/gmars/go-wechat/miniapp/kf_message"
	"github.com/gmars/go-wechat/miniapp/message_management"
	"github.com/gmars/go-wechat/miniapp/ocr"
	"github.com/gmars/go-wechat/miniapp/qrcode_link"
	"github.com/gmars/go-wechat/miniapp/security"
	"github.com/gmars/go-wechat/miniapp/user_info"
	"github.com/gmars/go-wechat/official/openapi"
)

// Code2Session 小程序登录
func (c *Client) Code2Session(jsCode string) (*user_info.LoginRes, error) {
	var (
		res     user_info.LoginRes
		request = core.NewApiRequest(nil)
	)
	_, err := request.JsonGet("/sns/jscode2session", map[string]string{
		"appid":      c.appId,
		"secret":     c.appSecret,
		"js_code":    jsCode,
		"grant_type": "authorization_code",
	}, &res)
	return &res, err
}

// OpenApi open api管理
func (c *Client) OpenApi() *openapi.OpenApi {
	return openapi.NewOpenApi(c.accessToken)
}

// UserInfo 用户信息
func (c *Client) UserInfo() *user_info.UserInfo {
	return user_info.NewUserInfo(c.accessToken)
}

// QrcodeLink 小程序码与小程序链接
func (c *Client) QrcodeLink() *qrcode_link.QrcodeLink {
	return qrcode_link.NewQrcodeLink(c.accessToken)
}

// MessageManagement 消息相关
func (c *Client) MessageManagement() *message_management.MessageManagement {
	return message_management.NewMessageManagement(c.accessToken)
}

// KfMessage 小程序客服
func (c *Client) KfMessage() *kf_message.KfMessage {
	return kf_message.NewKfMessage(c.accessToken)
}

// SecurityCenter 小程序安全
func (c *Client) SecurityCenter() *security.Security {
	return security.NewSecurityCenter(c.accessToken)
}

// Dev 运维
func (c *Client) Dev() *dev.Dev {
	return dev.NewDev(c.accessToken)
}

// ImageOcr 图片Ai处理及ocr
func (c *Client) ImageOcr() *ocr.Ocr {
	return ocr.NewOcr(c.accessToken)
}
