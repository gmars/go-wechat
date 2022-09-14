package jumpqrcode

import (
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/official/wxopen"
)

type JumpQrcode struct {
	wxopen.WxOpen
}

func NewJumpQrcode(authorizerAccessToken core.AccessToken) *JumpQrcode {
	return &JumpQrcode{wxopen.WxOpen{Request: core.NewApiRequest(authorizerAccessToken)}}
}

// DownloadQRCodeText 获取校验文件名称及内容
func (s *JumpQrcode) DownloadQRCodeText() (*DownloadQRCodeTextRes, error) {
	var res DownloadQRCodeTextRes
	_, err := s.Request.JsonPost("/cgi-bin/wxopen/qrcodejumpdownload", nil, nil, &res)
	return &res, err
}
