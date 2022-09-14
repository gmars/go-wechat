package subscribe

import (
	"github.com/gmars/go-wechat/core"
	"strconv"
)

type Subscribe struct {
	request *core.ApiRequest
}

func NewSubscribeManagement(authorizerAccessToken core.AccessToken) *Subscribe {
	return &Subscribe{request: core.NewApiRequest(authorizerAccessToken)}
}

// GetShowItem 获取已设置公众号信息
func (s *Subscribe) GetShowItem() (*GetShowItemRes, error) {
	var res GetShowItemRes
	_, err := s.request.JsonGet("/wxa/getshowwxaitem", nil, &res)
	return &res, err
}

// GetLinkForShow 获取可设置公众号列表
func (s *Subscribe) GetLinkForShow(page, pageSize int) (*GetLinkForShowRes, error) {
	var res GetLinkForShowRes
	_, err := s.request.JsonGet("/wxa/getwxamplinkforshow", map[string]string{
		"page": strconv.Itoa(page),
		"num":  strconv.Itoa(pageSize),
	}, &res)
	return &res, err
}

// SetShowItem 设置扫码关注的公众号
func (s *Subscribe) SetShowItem(wxaSubscribeBizFlag SubscribeBizFlag, appId string) error {
	_, err := s.request.JsonGet("/wxa/updateshowwxaitem", map[string]string{
		"wxa_subscribe_biz_flag": strconv.Itoa(wxaSubscribeBizFlag),
		"appid":                  appId,
	}, nil)
	return err
}
