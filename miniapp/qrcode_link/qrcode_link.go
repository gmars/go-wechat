package qrcode_link

import "github.com/gmars/go-wechat/core"

type QrcodeLink struct {
	request *core.ApiRequest
}

func NewQrcodeLink(accessToken core.AccessToken) *QrcodeLink {
	return &QrcodeLink{request: core.NewApiRequest(accessToken)}
}

// CreateQRCode 获取小程序二维码
func (s *QrcodeLink) CreateQRCode(path string, width int) (*WxaQrcodeRes, error) {
	var res WxaQrcodeRes
	if width == 0 {
		width = 430
	}
	_, err := s.request.JsonPost("/cgi-bin/wxaapp/createwxaqrcode", nil, map[string]interface{}{
		"path":  path,
		"width": width,
	}, &res)
	return &res, err
}

// GetQRCode 获取小程序码
func (s *QrcodeLink) GetQRCode(path string, addition *GetQrcodeParams) (*WxaQrcodeRes, error) {
	var (
		res   WxaQrcodeRes
		param = map[string]interface{}{
			"path": path,
		}
	)

	if addition != nil {
		if addition.Width == 0 {
			addition.Width = 430
		}
		param["with"] = addition.Width
		param["auto_color"] = addition.AutoColor
		param["line_color"] = addition.LineColor
		param["is_hyaline"] = addition.IsHyaline
	}

	_, err := s.request.JsonPost("/wxa/getwxacode", nil, param, &res)
	return &res, err
}

// GetUnlimitedQRCode 获取无限制小程序码
func (s *QrcodeLink) GetUnlimitedQRCode(scene string, addition *GetUnlimitedQrcodeParams) (*WxaQrcodeRes, error) {
	var (
		res   WxaQrcodeRes
		param = map[string]interface{}{
			"scene": scene,
		}
	)

	if addition != nil {
		if addition.Width == 0 {
			addition.Width = 430
		}
		param["with"] = addition.Width
		param["auto_color"] = addition.AutoColor
		param["line_color"] = addition.LineColor
		param["is_hyaline"] = addition.IsHyaline
		param["path"] = addition.Path
		param["check_path"] = addition.CheckPath
		param["env_version"] = addition.EnvVersion
	}

	_, err := s.request.JsonPost("/wxa/getwxacodeunlimit", nil, param, &res)
	return &res, err
}

// QueryScheme 查询 scheme 码
func (s *QrcodeLink) QueryScheme(scheme string) (*SchemeData, error) {
	var res SchemeData
	_, err := s.request.JsonPost("/wxa/queryscheme", nil, map[string]string{
		"scheme": scheme,
	}, &res)
	return &res, err
}

// GenerateScheme 获取 scheme 码
func (s *QrcodeLink) GenerateScheme(params *GenerateSchemeParams) (string, error) {
	var res GenerateSchemeRes
	_, err := s.request.JsonPost("/wxa/generatescheme", nil, params, &res)
	return res.OpenLink, err
}

// GenerateNFCScheme 获取 NFC 的小程序 scheme
func (s *QrcodeLink) GenerateNFCScheme(params *GenerateNFCSchemeParams) (string, error) {
	var res GenerateSchemeRes
	_, err := s.request.JsonPost("/wxa/generatenfcscheme", nil, params, &res)
	return res.OpenLink, err
}

// GenerateUrlLink 获取 URL Link
func (s *QrcodeLink) GenerateUrlLink(params *GenerateUrlLinkParams) (string, error) {
	var res GenerateUrlLinkRes
	_, err := s.request.JsonPost("/wxa/generate_urllink", nil, params, &res)
	return res.UrlLink, err
}

// QueryUrlLink 查询 URL Link
func (s *QrcodeLink) QueryUrlLink(urlLink string) (*QueryUrlLinkRes, error) {
	var res QueryUrlLinkRes
	_, err := s.request.JsonPost("/wxa/query_urllink", nil, map[string]string{
		"url_link": urlLink,
	}, &res)
	return &res, err
}

// GenerateShortLink 获取 Short Link
func (s *QrcodeLink) GenerateShortLink(pageUrl, pageTitle string, isPermanent bool) (string, error) {
	var res GenerateShortLinkRes
	_, err := s.request.JsonPost("/wxa/genwxashortlink", nil, map[string]interface{}{
		"page_url":     pageUrl,
		"page_title":   pageTitle,
		"is_permanent": isPermanent,
	}, &res)
	return res.Link, err
}
