package authorizer_token

import "golang.org/x/sync/singleflight"

var gsf singleflight.Group

type AuthorizerTokenRes struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

const (
	authorizerAccessTokenCacheKeyPrefix  = "AUTHORIZER_ACCESS_TOKEN_CACHE_"
	authorizerRefreshTokenCacheKeyPrefix = "AUTHORIZER_REFRESH_TOKEN_"
)
