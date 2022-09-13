package subscription

import (
	"go-wechat/core"
	"go-wechat/util"
	"strconv"
)

type Subscription struct {
	Request *core.ApiRequest
}

func NewSubscription(token core.AccessToken) *Subscription {
	return &Subscription{Request: core.NewApiRequest(token)}
}

// AddTemplate 选用公共模板到私有库
func (s *Subscription) AddTemplate(params *AddTempParams) (string, error) {
	var res AddTempRes
	_, err := s.Request.JsonPost("/wxaapi/newtmpl/addtemplate", nil, params, &res)
	if err != nil {
		return "", err
	}

	return res.PriTmplId, nil
}

// DeleteTemplate 删除私有库中的模板
func (s *Subscription) DeleteTemplate(priTmplId string) error {
	_, err := s.Request.JsonPost("/wxaapi/newtmpl/deltemplate", nil, map[string]string{
		"priTmplId": priTmplId,
	}, nil)
	return err
}

// GetCategory 获取公众号设置的类目
func (s *Subscription) GetCategory() (*[]CategoryItem, error) {
	var res CategoryRest
	_, err := s.Request.JsonGet("/wxaapi/newtmpl/getcategory", nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// GetPubTemplateKeywords 获取公共模板下的关键词列表
func (s *Subscription) GetPubTemplateKeywords(tid string) (*PubKeywordsRes, error) {
	var res PubKeywordsRes
	_, err := s.Request.JsonGet("/wxaapi/newtmpl/getpubtemplatekeywords", map[string]string{
		"tid": tid,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetPubTemplateTitles 获取类目下的公共模板
func (s *Subscription) GetPubTemplateTitles(ids string, page, pageSize int) (*TemplateTitleListRes, error) {
	var res TemplateTitleListRes
	offset, count := util.PageCondition(page, pageSize, 30)
	_, err := s.Request.JsonGet("/wxaapi/newtmpl/getpubtemplatetitles", map[string]string{
		"ids":   ids,
		"start": strconv.Itoa(offset),
		"limit": strconv.Itoa(count),
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetTemplate 获取私有模板列表
func (s *Subscription) GetTemplate() (*PrivateTemplateRes, error) {
	var res PrivateTemplateRes

	_, err := s.Request.JsonGet("/wxaapi/newtmpl/gettemplate", nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// BizSend 发送订阅通知
func (s *Subscription) BizSend(params *SendParams) error {
	_, err := s.Request.JsonPost("/cgi-bin/message/subscribe/bizsend", nil, params, nil)
	return err
}
