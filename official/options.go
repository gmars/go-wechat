package official

import (
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/official/accesstoken"
	"github.com/gmars/go-wechat/official/message"
	"github.com/gmars/go-wechat/util"
)

type BaseConfig struct {
	AppId     string
	AppSecret string
	Cache     core.Cache
}

// WithBaseAutoAccessTokenConfig api调用基本配置
func WithBaseAutoAccessTokenConfig(appId, appSecret string, cache core.Cache) *BaseConfig {
	return &BaseConfig{
		AppId:     appId,
		AppSecret: appSecret,
		Cache:     cache,
	}
}

func (a *BaseConfig) Apply(c *Client) error {
	var (
		err   error
		cache = a.Cache
	)

	if a.Cache == nil {
		cache, err = util.NewFileCache("")
		if err != nil {
			return err
		}
	} else {
		cache = a.Cache
	}
	c.accessToken = accesstoken.NewAccessToken(a.AppId, a.AppSecret, cache)
	return nil
}

type MessageConfig struct {
	AppId  string
	Token  string
	AesKey string
}

// WithMessageHandlerConfig 配置message handler
func WithMessageHandlerConfig(appId, token, aesKey string) *MessageConfig {
	return &MessageConfig{
		AppId:  appId,
		Token:  token,
		AesKey: aesKey,
	}
}

func (a *MessageConfig) Apply(c *Client) error {
	c.messageHandler = message.NewMessageHandler(a.AesKey, a.Token, a.AppId)
	return nil
}

type AccessTokenConfig struct {
	AccessToken core.AccessToken
}

// WithAccessToken 配置access token
// 该方法可跳过基础配置
func WithAccessToken(accessToken core.AccessToken) *AccessTokenConfig {
	return &AccessTokenConfig{
		AccessToken: accessToken,
	}
}

func (a *AccessTokenConfig) Apply(c *Client) error {
	c.accessToken = a.AccessToken
	return nil
}
