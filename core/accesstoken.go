package core

import "context"

type AccessToken interface {
	GetAccessToken(ctx context.Context) (string, error)
	RefreshAccessToken(ctx context.Context) (string, error)
	GetCurrentAppid(ctx context.Context) string
	GetCacheKey(ctx context.Context) string
}
