package component

import (
	"context"
	"errors"
	"github.com/gmars/go-wechat/core"
)

type AccessToken struct {
	AppId     string
	AppSecret string
	Cache     core.Cache
}

func NewComponentAccessToken(appId, appSecret string, cache core.Cache) *AccessToken {
	return &AccessToken{
		AppId:     appId,
		AppSecret: appSecret,
		Cache:     cache,
	}
}

func (c *AccessToken) GetAccessToken(ctx context.Context) (string, error) {
	token, _ := c.Cache.GetData(ctx, c.GetCacheKey(ctx))
	if token == "" {
		return c.RefreshAccessToken(ctx)
	}

	return token, nil

}

// RefreshAccessToken 刷新开放平台refresh access token
func (c *AccessToken) RefreshAccessToken(ctx context.Context) (string, error) {

	val, err, _ := gsf.Do(c.GetCacheKey(ctx), func() (interface{}, error) {
		return c.fetchComponentToken(ctx)
	})

	if err != nil {
		return "", err
	}

	token, ok := val.(*ComAccessToken)
	if !ok {
		return "", errors.New("解析token结果出错")
	}

	err = c.Cache.SetData(ctx, c.GetCacheKey(ctx), token.ComponentAccessToken, int64(token.ExpiresIn-600))
	if err != nil {
		return "", err
	}

	return token.ComponentAccessToken, nil
}

func (c *AccessToken) GetCurrentAppid(ctx context.Context) string {
	return c.AppId
}

func (c *AccessToken) GetCacheKey(ctx context.Context) string {
	return componentAccessTokenCacheKeyPrefix + c.GetCurrentAppid(ctx)
}

// GetComponentTicket 获取验证票据
func (c *AccessToken) getComponentTicket(ctx context.Context) (string, error) {
	var (
		cacheKey = getComponentTicketCacheKey(c.GetCurrentAppid(context.Background()))
	)
	return c.Cache.GetData(ctx, cacheKey)
}

func (c *AccessToken) fetchComponentToken(ctx context.Context) (*ComAccessToken, error) {
	var res ComAccessToken
	ticket, _ := c.getComponentTicket(ctx)
	if ticket == "" {
		return nil, errors.New("component ticket not found")
	}

	req := core.NewApiRequest(nil)
	_, err := req.JsonPost("/cgi-bin/component/api_component_token", nil, map[string]string{
		"component_appid":         c.AppId,
		"component_appsecret":     c.AppSecret,
		"component_verify_ticket": ticket,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
