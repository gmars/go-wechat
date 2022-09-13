package official

import (
	"go-wechat/core"
	"go-wechat/official/message"
)

type Options interface {
	Apply(c *Client) error
}

type Client struct {
	accessToken    core.AccessToken
	messageHandler *message.Message
}

func NewOfficialClient(opts ...Options) (*Client, error) {
	c := new(Client)
	err := initOfficialClient(c, opts)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func initOfficialClient(c *Client, opts []Options) error {
	for _, opt := range opts {
		if err := opt.Apply(c); err != nil {
			return err
		}
	}

	return nil
}
