package account

import (
	"github.com/gmars/go-wechat/core"
	"unicode/utf8"
)

type Account struct {
	Request *core.ApiRequest
}

func NewAccount(token core.AccessToken) *Account {
	return &Account{Request: core.NewApiRequest(token)}
}

// CreateLimitStrQrcode 创建临时二维码场景值是字符串
func (s *Account) CreateLimitStrQrcode(sceneStr string) (*CreateQrCodeRes, error) {
	if utf8.RuneCountInString(sceneStr) > 64 || sceneStr == "" {
		return nil, core.NewError(500, "字符串场景值不能为空且不能超过64个字符")
	}

	body := map[string]interface{}{
		"action_name": "QR_LIMIT_STR_SCENE",
		"action_info": map[string]interface{}{"scene": map[string]string{"scene_str": sceneStr}},
	}
	return s.generateQrcode(body)
}

// CreateLimitIdQrcode 创建临时二维码场景值是数字
func (s *Account) CreateLimitIdQrcode(sceneId uint32) (*CreateQrCodeRes, error) {
	if sceneId > 100000 {
		return nil, core.NewError(500, "永久二维码场景id值不能超过100000")
	}

	body := map[string]interface{}{
		"action_name": "QR_LIMIT_SCENE",
		"action_info": map[string]interface{}{"scene": map[string]uint32{"scene_id": sceneId}},
	}
	return s.generateQrcode(body)
}

// CreateStrQrcode 创建永久二维码场景值是字符串
func (s *Account) CreateStrQrcode(sceneStr string, expireSeconds int) (*CreateQrCodeRes, error) {
	if expireSeconds > 2592000 {
		return nil, core.NewError(500, "二维码有效期不能超过30天")
	}

	if utf8.RuneCountInString(sceneStr) > 64 || sceneStr == "" {
		return nil, core.NewError(500, "字符串场景值不能为空且不能超过64个字符")
	}

	body := map[string]interface{}{
		"expire_seconds": expireSeconds,
		"action_name":    "QR_STR_SCENE",
		"action_info":    map[string]interface{}{"scene": map[string]string{"scene_str": sceneStr}},
	}
	return s.generateQrcode(body)
}

// CreateIdQrcode 创建永久二维码场景值是数字
func (s *Account) CreateIdQrcode(sceneId uint32, expireSeconds int) (*CreateQrCodeRes, error) {
	if expireSeconds > 2592000 {
		return nil, core.NewError(500, "二维码有效期不能超过30天")
	}

	body := map[string]interface{}{
		"expire_seconds": expireSeconds,
		"action_name":    "QR_SCENE",
		"action_info":    map[string]interface{}{"scene": map[string]uint32{"scene_id": sceneId}},
	}
	return s.generateQrcode(body)
}

// GetQrcodeByTicket 获取二维码图片
func (s *Account) GetQrcodeByTicket(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + ticket
}

// GenShorten 生成短key
func (s *Account) GenShorten(longData string, expireSeconds int) (string, error) {
	var res GenShortenRes
	if expireSeconds > 2592000 || expireSeconds < 1 {
		return "", core.NewError(500, "短Key有效期不能超过30天且不能小于1秒")
	}

	if len(longData) > 4<<10 {
		return "", core.NewError(500, "长信息的长度不能超过4KB")
	}

	_, err := s.Request.JsonPost("/cgi-bin/shorten/gen", nil, map[string]interface{}{
		"long_data":      longData,
		"expire_seconds": expireSeconds,
	}, &res)
	if err != nil {
		return "", err
	}

	return res.ShortKey, nil
}

// FetchShorten 查询长信息
func (s *Account) FetchShorten(shortKey string) (*FetchShortenRes, error) {
	var res FetchShortenRes
	_, err := s.Request.JsonPost("/cgi-bin/shorten/fetch", nil, map[string]interface{}{
		"short_key": shortKey,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// 生成二维码
func (s *Account) generateQrcode(body interface{}) (*CreateQrCodeRes, error) {
	var res CreateQrCodeRes
	_, err := s.Request.JsonPost("/cgi-bin/qrcode/create", nil, body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
