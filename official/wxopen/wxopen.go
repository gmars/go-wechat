package wxopen

import (
	"github.com/gmars/go-wechat/core"
)

type WxOpen struct {
	Request *core.ApiRequest
}

func NewWxOpen(token core.AccessToken) *WxOpen {
	return &WxOpen{Request: core.NewApiRequest(token)}
}

// SaveJumpQrCode 创建服务号跳转小程序二维码
func (s *WxOpen) SaveJumpQrCode(params *SaveJumpQrCodeParams) error {
	_, err := s.Request.JsonPost("/cgi-bin/wxopen/qrcodejumpadd", nil, params, nil)
	return err
}

// DeleteJumpQrCode 删除服务号跳转小程序二维码
func (s *WxOpen) DeleteJumpQrCode(prefix, appId string) error {
	_, err := s.Request.JsonPost("/cgi-bin/wxopen/qrcodejumpdelete", nil, map[string]string{
		"prefix": prefix,
		"appid":  appId,
	}, nil)
	return err
}

// GetJumpQrCodeByPrefix 使用prefix获取已设置的服务号跳转小程序规则
func (s *WxOpen) GetJumpQrCodeByPrefix(miniAppId string, prefixList []string) (*JumpQrCodeListRes, error) {
	if prefixList == nil {
		return nil, core.NewError(500, "请传入前缀列表")
	}

	body := GetJumpQrCodeParams{
		Appid:      miniAppId,
		GetType:    1,
		PrefixList: prefixList,
	}
	return s.fetchJumpQrCode(&body)
}

// PageListJumpQrCode 分页获取已设置的服务号跳转小程序规则
func (s *WxOpen) PageListJumpQrCode(miniAppId string, page, pageSize uint) (*JumpQrCodeListRes, error) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 || pageSize > 200 {
		pageSize = 20
	}

	body := GetJumpQrCodeParams{
		Appid:    miniAppId,
		GetType:  2,
		PageNum:  page,
		PageSize: pageSize,
	}
	return s.fetchJumpQrCode(&body)
}

// PublishJumpQrCode 发布已设置的二维码规则
func (s *WxOpen) PublishJumpQrCode(prefix string) error {
	_, err := s.Request.JsonPost("/cgi-bin/wxopen/qrcodejumppublish", nil, map[string]string{
		"prefix": prefix,
	}, nil)
	return err
}

// WxAmpLinkGet 获取公众号关联小程序
func (s *WxOpen) WxAmpLinkGet() ([]WxAmpLinkItem, error) {
	var res WxAmpLinkRes
	_, err := s.Request.JsonPost("/cgi-bin/wxopen/wxamplinkget", nil, map[string]string{}, &res)
	return res.WxOpens.Items, err
}

// WxAmpLink 公众号关联小程序
func (s *WxOpen) WxAmpLink(miniAppAppId string, notifyUsers int, showProfile int) error {
	_, err := s.Request.JsonPost("/cgi-bin/wxopen/wxamplink", nil, map[string]interface{}{
		"appid":        miniAppAppId,
		"notify_users": notifyUsers,
		"show_profile": showProfile,
	}, nil)
	return err
}

// WxAmpUnlink 公众号取消关联小程序
func (s *WxOpen) WxAmpUnlink(miniAppAppId string) error {
	_, err := s.Request.JsonPost("/cgi-bin/wxopen/wxampunlink", nil, map[string]interface{}{
		"appid": miniAppAppId,
	}, nil)
	return err
}

// 获取已设置的服务号跳转小程序规则
func (s *WxOpen) fetchJumpQrCode(params *GetJumpQrCodeParams) (*JumpQrCodeListRes, error) {
	var res JumpQrCodeListRes
	_, err := s.Request.JsonPost("/cgi-bin/wxopen/qrcodejumpget", nil, params, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
