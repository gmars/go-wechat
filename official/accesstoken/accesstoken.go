package accesstoken

import (
	"context"
	"errors"
	"fmt"
	"go-wechat/core"
	"go-wechat/util"
)

type AccessToken struct {
	AppId     string
	AppSecret string
	Cache     core.Cache
}

func NewAccessToken(appId, appSecret string, cache core.Cache) *AccessToken {
	return &AccessToken{
		AppId:     appId,
		AppSecret: appSecret,
		Cache:     cache,
	}
}

// GetAccessToken 获取access token
func (a *AccessToken) GetAccessToken(ctx context.Context) (string, error) {
	token, err := a.Cache.GetData(ctx, a.GetCacheKey(ctx))
	if err != nil && (errors.Is(err, util.IsNotExist) || errors.Is(err, util.IsExpires)) {
		return a.RefreshAccessToken(ctx)
	} else if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

// RefreshAccessToken 刷新access token缓存
func (a *AccessToken) RefreshAccessToken(ctx context.Context) (string, error) {
	val, err, _ := gsf.Do(a.GetCacheKey(ctx), func() (interface{}, error) {
		return a.fetchToken(ctx)
	})

	if err != nil {
		return "", err
	}
	token, ok := val.(*TokenRes)
	if !ok {
		return "", errors.New("解析token结果出错")
	}
	err = a.Cache.SetData(ctx, a.GetCacheKey(ctx), token.AccessToken, int64(token.ExpiresIn))
	if err != nil {
		fmt.Printf("official accesstoken cache is not valiable:%s", err.Error())
	}
	return token.AccessToken, nil
}

func (a *AccessToken) GetCurrentAppid(ctx context.Context) string {
	return a.AppId
}

func (a *AccessToken) GetCacheKey(ctx context.Context) string {
	return cachePrefix + a.GetCurrentAppid(ctx)
}

func (a *AccessToken) fetchToken(ctx context.Context) (*TokenRes, error) {
	var res TokenRes
	req := core.NewApiRequest(nil)
	_, err := req.JsonGet(accessTokenUrl, map[string]string{
		"grant_type": "client_credential",
		"appid":      a.GetCurrentAppid(ctx),
		"secret":     a.AppSecret,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
