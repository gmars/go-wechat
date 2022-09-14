package miniapp

import (
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/official/accesstoken"
	"github.com/gmars/go-wechat/util"
)

type BaseConfig struct {
	AppId     string
	AppSecret string
	Cache     core.Cache
}

// WithBaseConfig 基本参数配置
func WithBaseConfig(appId, appSecret string, cache core.Cache) *BaseConfig {
	return &BaseConfig{
		AppId:     appId,
		AppSecret: appSecret,
		Cache:     cache,
	}
}

func (a *BaseConfig) Apply(c *Client) error {
	c.cache = a.Cache
	c.appId = a.AppId
	c.appSecret = a.AppSecret
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
	c.appId = a.AppId
	c.appSecret = a.AppSecret
	c.accessToken = accesstoken.NewAccessToken(a.AppId, a.AppSecret, cache)
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
