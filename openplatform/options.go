package openplatform

import (
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/official/message"
	"github.com/gmars/go-wechat/openplatform/component"
	"github.com/gmars/go-wechat/util"
)

type Cache struct {
	Cache core.Cache
}

func WithCache(cache core.Cache) *Cache {
	return &Cache{Cache: cache}
}

func (a *Cache) Apply(c *Client) error {
	c.cache = a.Cache
	return nil
}

type AutoAccessTokenConfig struct {
	AppId     string
	AppSecret string
	Cache     core.Cache
}

// WithAutoAccessTokenConfig api调用基本配置
func WithAutoAccessTokenConfig(appId, appSecret string, cache core.Cache) *AutoAccessTokenConfig {
	return &AutoAccessTokenConfig{
		AppId:     appId,
		AppSecret: appSecret,
		Cache:     cache,
	}
}

func (a *AutoAccessTokenConfig) Apply(c *Client) error {
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
	c.cache = cache
	c.componentAccessToken = component.NewComponentAccessToken(a.AppId, a.AppSecret, cache)
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
