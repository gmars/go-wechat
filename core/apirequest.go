package core

import (
	"context"
	"encoding/json"
	"github.com/gmars/go-wechat/util"
	"mime/multipart"
)

const ApiHost = "https://api.weixin.qq.com"

type ApiRequest struct {
	AccessToken AccessToken
}

func NewApiRequest(token AccessToken) *ApiRequest {
	return &ApiRequest{AccessToken: token}
}

// JsonPost post json请求
func (a *ApiRequest) JsonPost(urlPath string, query map[string]string, body interface{}, res interface{}) ([]byte, error) {
	req := util.NewRequest(a.generateApiUrl(urlPath), "POST")
	err := a.generateQuery(req, query)
	if err != nil {
		return nil, err
	}
	if body != nil {
		if err = req.WithJsonBody(body); err != nil {
			return nil, err
		}
	}

	return a.parseResponse(req, res)
}

// JsonGet get json请求
func (a *ApiRequest) JsonGet(urlPath string, query map[string]string, res interface{}) ([]byte, error) {
	req := util.NewRequest(a.generateApiUrl(urlPath), "GET")
	err := a.generateQuery(req, query)
	if err != nil {
		return nil, err
	}

	return a.parseResponse(req, res)
}

// FormPost post form请求
func (a *ApiRequest) FormPost(urlPath string, query map[string]string, file map[string]*multipart.FileHeader, body map[string]string, res interface{}) ([]byte, error) {
	req := util.NewRequest(a.generateApiUrl(urlPath), "POST")
	err := a.generateQuery(req, query)
	if err != nil {
		return nil, err
	}

	err = req.WithFormBody(body, file)
	if err != nil {
		return nil, err
	}

	return a.parseResponse(req, res)
}

// 处理query参数
func (a *ApiRequest) generateApiUrl(path string) string {
	return ApiHost + path
}

// 处理query参数
func (a *ApiRequest) parseResponse(r *util.Request, res interface{}) ([]byte, error) {
	var ret ApiError
	resp, err := r.Do(context.Background())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &ret)
	if ret.ErrCode != 0 {
		return nil, &ret
	}

	if res != nil {
		if err = json.Unmarshal(resp, &res); err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// 处理query参数
func (a *ApiRequest) generateQuery(r *util.Request, query map[string]string) error {
	if a.AccessToken != nil {
		token, err := a.AccessToken.GetAccessToken(context.Background())
		if err != nil {
			return err
		}
		if query == nil {
			query = make(map[string]string)
		}
		query["access_token"] = token
	}

	if query == nil {
		return nil
	}

	return r.WithQuery(query)
}
