package component

import (
	"context"
	"fmt"
	"go-wechat/core"
	"go-wechat/official/message"
	"net/http"
)

type Notify struct {
	cache   core.Cache
	message *message.Message
}

func NewAuthorizationNotify(message *message.Message, cache core.Cache) *Notify {
	return &Notify{
		cache:   cache,
		message: message,
	}
}

// AuthorizationHandler 授权事件处理器
func (c *Notify) AuthorizationHandler(req *http.Request) (*message.ReplyMsg, error) {
	msg, err := c.message.Handler(req)
	if err != nil {
		return nil, err
	}

	if msg.ReplayType != message.ReplyTypeInfo {
		return nil, core.NewError(400, "该消息非开放平台info type")
	}

	if msg.TypeValue == message.InfoTypeComponentVerifyTicket {
		_ = c.cacheComponentTicket(msg.Data.(*message.ComponentVerifyTicketInfo))
		return msg, nil
	}
	return msg, nil
}

// 缓存验证票据
func (c *Notify) cacheComponentTicket(ticketMsg *message.ComponentVerifyTicketInfo) error {
	var (
		cacheKey       = getComponentTicketCacheKey(ticketMsg.AppId)
		expiresSeconds = 12 * 3600
	)

	if err := c.cache.SetData(context.Background(), cacheKey, ticketMsg.ComponentVerifyTicket, int64(expiresSeconds)); err != nil {
		fmt.Printf("缓存开放平台%s的ticket出错：%s", ticketMsg.AppId, err.Error())
	}
	return nil
}
