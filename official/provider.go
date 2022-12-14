package official

import (
	"github.com/gmars/go-wechat/official/account"
	"github.com/gmars/go-wechat/official/assetmanage"
	"github.com/gmars/go-wechat/official/customerservice"
	"github.com/gmars/go-wechat/official/draft"
	"github.com/gmars/go-wechat/official/menu"
	"github.com/gmars/go-wechat/official/message"
	"github.com/gmars/go-wechat/official/openapi"
	"github.com/gmars/go-wechat/official/subscription"
	"github.com/gmars/go-wechat/official/template_msg"
	"github.com/gmars/go-wechat/official/user"
	"github.com/gmars/go-wechat/official/wxopen"
	account2 "github.com/gmars/go-wechat/openplatform/account"
)

// MessageHandler 消息处理器
func (c *Client) MessageHandler() *message.Message {
	return c.messageHandler
}

// Account 账号管理[生成带参二维码 key托管等]
func (c *Client) Account() *account.Account {
	return account.NewAccount(c.accessToken)
}

// AssentManage 素材管理
func (c *Client) AssentManage() *assetmanage.AssetsManage {
	return assetmanage.NewAssetsManage(c.accessToken)
}

// CustomerService 客服消息
func (c *Client) CustomerService() *customerservice.CustomerService {
	return customerservice.NewCustomerService(c.accessToken)
}

// Draft 草稿箱
func (c *Client) Draft() *draft.Draft {
	return draft.NewDraft(c.accessToken)
}

// Menu 菜单管理
func (c *Client) Menu() *menu.Menu {
	return menu.NewMenu(c.accessToken)
}

// OpenApi 开放api管理
func (c *Client) OpenApi() *openapi.OpenApi {
	return openapi.NewOpenApi(c.accessToken)
}

// Subscription 订阅通知
func (c *Client) Subscription() *subscription.Subscription {
	return subscription.NewSubscription(c.accessToken)
}

// User 用户管理
func (c *Client) User() *user.User {
	return user.NewUserManagement(c.accessToken)
}

// WxOpen 服务号二维码打开小程序设置
func (c *Client) WxOpen() *wxopen.WxOpen {
	return wxopen.NewWxOpen(c.accessToken)
}

// TemplateMessage 模板消息
func (c *Client) TemplateMessage() *template_msg.TemplateMessage {
	return template_msg.NewTemplateMessage(c.accessToken)
}

// OpenPlatform 开放平台账号操作
func (c *Client) OpenPlatform() *account2.Account {
	return account2.NewOpenAccount(c.accessToken)
}
