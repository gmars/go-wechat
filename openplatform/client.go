package openplatform

import (
	"go-wechat/core"
	"go-wechat/official/message"
)

type Options interface {
	Apply(c *Client) error
}

type Client struct {
	cache                core.Cache
	componentAccessToken core.AccessToken
	messageHandler       *message.Message
}

func NewOpenPlatformClient(opts ...Options) (*Client, error) {
	c := new(Client)
	err := initOpenPlatformConfig(c, opts)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func initOpenPlatformConfig(c *Client, opts []Options) error {
	for _, opt := range opts {
		if err := opt.Apply(c); err != nil {
			return err
		}
	}

	return nil
}
