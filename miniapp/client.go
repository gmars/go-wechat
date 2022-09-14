package miniapp

import (
	"github.com/gmars/go-wechat/core"
)

type Options interface {
	Apply(c *Client) error
}

type Client struct {
	appId       string
	appSecret   string
	cache       core.Cache
	accessToken core.AccessToken
}

func NewMiniAppClient(opts ...Options) (*Client, error) {
	c := new(Client)
	err := initMiniAppClient(c, opts)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func initMiniAppClient(c *Client, opts []Options) error {
	for _, opt := range opts {
		if err := opt.Apply(c); err != nil {
			return err
		}
	}

	return nil
}
