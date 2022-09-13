package component

import "golang.org/x/sync/singleflight"

var gsf singleflight.Group

type ComAccessToken struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}

const (
	componentTicketCacheKeyPrefix      = "COMPONENTS_TICKET_CACHE_"
	componentAccessTokenCacheKeyPrefix = "COMPONENT_ACCESS_TOKEN_CACHE_"
)

func getComponentTicketCacheKey(appid string) string {
	return componentTicketCacheKeyPrefix + appid
}
