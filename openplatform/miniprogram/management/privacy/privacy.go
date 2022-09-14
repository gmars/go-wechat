package privacy

import (
	"github.com/gmars/go-wechat/core"
	"mime/multipart"
)

type Privacy struct {
	request *core.ApiRequest
}

func NewPrivacyManagement(authorizerAccessToken core.AccessToken) *Privacy {
	return &Privacy{request: core.NewApiRequest(authorizerAccessToken)}
}

// SetPrivacySetting 上传小程序用户隐私保护指引
func (s *Privacy) SetPrivacySetting(params *UploadPrivacySettingParams) error {
	_, err := s.request.JsonPost("/cgi-bin/component/setprivacysetting", nil, params, nil)
	return err
}

// GetPrivacySetting 获取小程序用户隐私保护指引
func (s *Privacy) GetPrivacySetting(privacyVer int) (*GetPrivacySettingRes, error) {
	var res GetPrivacySettingRes
	_, err := s.request.JsonPost("/cgi-bin/component/getprivacysetting", nil, map[string]int{
		"privacy_ver": privacyVer,
	}, &res)
	return &res, err
}

// UploadPrivacySetting 上传小程序用户隐私保护指引
func (s *Privacy) UploadPrivacySetting(file *multipart.FileHeader) (string, error) {
	var res UploadPrivacySettingRes
	_, err := s.request.FormPost("/cgi-bin/component/uploadprivacyextfile", nil, map[string]*multipart.FileHeader{
		"file": file,
	}, nil, &res)
	return res.ExtFileMediaId, err
}

// ApplyPrivacyInterface 申请地理位置接口
func (s *Privacy) ApplyPrivacyInterface(params *ApplyPrivacyInterfaceParams) (int, error) {
	var res ApplyPrivacyInterfaceRes
	_, err := s.request.JsonPost("/wxa/security/apply_privacy_interface", nil, params, &res)
	return res.AuditId, err
}

// GetPrivacyInterface 获取地理位置接口列表
func (s *Privacy) GetPrivacyInterface() (*[]InterfaceItem, error) {
	var res GetPrivacyInterfaceRes
	_, err := s.request.JsonGet("/wxa/security/get_privacy_interface", nil, &res)
	return &res.InterfaceList, err
}
