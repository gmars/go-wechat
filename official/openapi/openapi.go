package openapi

import (
	"context"
	"github.com/gmars/go-wechat/core"
)

type OpenApi struct {
	Request      *core.ApiRequest
	TokenHandler core.AccessToken
}

func NewOpenApi(token core.AccessToken) *OpenApi {
	return &OpenApi{Request: core.NewApiRequest(token), TokenHandler: token}
}

// ClearQuota 清空api调用次数
func (m *OpenApi) ClearQuota() error {
	_, err := m.Request.JsonPost("/cgi-bin/clear_quota", nil, map[string]string{
		"appid": m.TokenHandler.GetCurrentAppid(context.Background()),
	}, nil)
	return err
}

// GetQuota 查询api调用quota
func (m *OpenApi) GetQuota(cgiPath string) (*QuotaData, error) {
	var res QuotaRes
	_, err := m.Request.JsonPost("/cgi-bin/openapi/quota/get", nil, map[string]string{
		"cgi_path": cgiPath,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res.Quota, nil
}

// GetRid 查询接口调用报错信息
func (m *OpenApi) GetRid(rid string) (*RidRequest, error) {
	var res RidRequestRes
	_, err := m.Request.JsonPost("/cgi-bin/openapi/rid/get", nil, map[string]string{
		"rid": rid,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res.Request, nil
}
