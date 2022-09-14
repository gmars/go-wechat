package authorizer_token

import (
	"context"
	"errors"
	"fmt"
	"github.com/gmars/go-wechat/core"
)

type AuthorizerToken struct {
	request                    *core.ApiRequest
	appId                      string
	authorizerRefreshToken     string
	componentAccessToken       core.AccessToken
	refreshTokenChangeCallBack func(string, *AuthorizerTokenRes)
	cache                      core.Cache
}

func NewAuthorizerToken(appId, refreshToken string, componentAccessToken core.AccessToken, cache core.Cache, callBack func(string, *AuthorizerTokenRes)) *AuthorizerToken {
	return &AuthorizerToken{
		request:                    core.NewApiRequest(componentAccessToken),
		appId:                      appId,
		authorizerRefreshToken:     refreshToken,
		componentAccessToken:       componentAccessToken,
		refreshTokenChangeCallBack: callBack,
		cache:                      cache,
	}
}

func (a *AuthorizerToken) GetAccessToken(ctx context.Context) (string, error) {
	token, err := a.cache.GetData(ctx, a.GetCacheKey(ctx))
	if err == nil {
		return token, nil
	}
	return a.RefreshAccessToken(ctx)
}

func (a *AuthorizerToken) RefreshAccessToken(ctx context.Context) (string, error) {
	val, err, _ := gsf.Do(a.GetCacheKey(ctx), func() (interface{}, error) {
		return a.fetchAuthorizerAccessToken(ctx)
	})

	if err != nil {
		return "", err
	}

	token, ok := val.(*AuthorizerTokenRes)
	if !ok {
		return "", errors.New("解析token结果出错")
	}

	err = a.cache.SetData(ctx, a.GetCacheKey(ctx), token.AuthorizerAccessToken, int64(token.ExpiresIn-600))
	if err != nil {
		return "", err
	}

	return token.AuthorizerAccessToken, nil
}

func (a *AuthorizerToken) GetCurrentAppid(ctx context.Context) string {
	return a.appId
}

func (a *AuthorizerToken) GetCacheKey(ctx context.Context) string {
	return authorizerAccessTokenCacheKeyPrefix + a.appId
}

func (a *AuthorizerToken) fetchAuthorizerAccessToken(ctx context.Context) (*AuthorizerTokenRes, error) {
	var res AuthorizerTokenRes
	_, err := a.request.JsonPost("/cgi-bin/component/api_authorizer_token", nil, map[string]string{
		"component_appid":          a.componentAccessToken.GetCurrentAppid(ctx),
		"authorizer_appid":         a.appId,
		"authorizer_refresh_token": a.authorizerRefreshToken,
	}, &res)
	if err != nil {
		return nil, err
	}

	//存入缓存
	_ = a.cacheRefreshToken(ctx, &res)
	err = a.cache.SetData(ctx, authorizerRefreshTokenCacheKeyPrefix+a.appId, res.AuthorizerRefreshToken, int64(311040000))
	if err != nil {
		fmt.Printf("authorization refresh token 存入缓存出错，值为【%s】,错误:%s", res.AuthorizerRefreshToken, err.Error())
	}
	a.authorizerRefreshToken = res.AuthorizerRefreshToken
	//回调到程序
	a.refreshTokenChangeCallBack(a.appId, &res)

	return &res, nil
}

// 缓存refresh值
func (a *AuthorizerToken) cacheRefreshToken(ctx context.Context, res *AuthorizerTokenRes) error {
	err := a.cache.SetData(ctx, authorizerRefreshTokenCacheKeyPrefix+a.appId, res.AuthorizerRefreshToken, int64(311040000))
	if err != nil {
		fmt.Printf("authorization refresh token 存入缓存出错，值为【%s】,错误:%s", res.AuthorizerRefreshToken, err.Error())
	}
	return err
}
