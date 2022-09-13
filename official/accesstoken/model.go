package accesstoken

import "golang.org/x/sync/singleflight"

var gsf singleflight.Group

const (
	accessTokenUrl = "/cgi-bin/token"
	cachePrefix    = "ACCESS_TOKEN_CACHE_"
)

type TokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
