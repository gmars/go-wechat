package message_management

import (
	"go-wechat/core"
	"go-wechat/official/subscription"
)

type MessageManagement struct {
	request            *core.ApiRequest
	templateManagement *subscription.Subscription
}

func NewMessageManagement(accessToken core.AccessToken) *MessageManagement {
	return &MessageManagement{
		request:            core.NewApiRequest(accessToken),
		templateManagement: subscription.NewSubscription(accessToken),
	}
}

// SendUniformMessage 下发统一消息
func (s *MessageManagement) SendUniformMessage(toUser string, weAppMsg WeAppTemplateMsg, msg MpTemplateMsg) error {
	body := SendUniformMessageParams{
		ToUser:           toUser,
		WeAppTemplateMsg: weAppMsg,
		MpTemplateMsg:    msg,
	}
	_, err := s.request.JsonPost("/cgi-bin/message/wxopen/template/uniform_send", nil, body, nil)
	return err
}

// CreateActivityId 创建activity_id
func (s *MessageManagement) CreateActivityId(unionId, openid string) (*CreateActivityIdRes, error) {
	var res CreateActivityIdRes
	_, err := s.request.JsonGet("/cgi-bin/message/wxopen/activityid/create", map[string]string{
		"unionid": unionId,
		"openid":  openid,
	}, nil)
	return &res, err
}

// SetUpdatableMsg 修改动态消息
func (s *MessageManagement) SetUpdatableMsg(params *SetUpdatableMsgRes) error {
	_, err := s.request.JsonPost("/cgi-bin/message/wxopen/updatablemsg/send", nil, params, nil)
	return err
}

// DeleteMessageTemplate 删除模板
func (s *MessageManagement) DeleteMessageTemplate(priTmplId string) error {
	return s.templateManagement.DeleteTemplate(priTmplId)
}

// GetCategory 获取类目
func (s *MessageManagement) GetCategory(priTmplId string) (*[]subscription.CategoryItem, error) {
	return s.templateManagement.GetCategory()
}

// GetPubTemplateKeyWordsById 获取关键词列表
func (s *MessageManagement) GetPubTemplateKeyWordsById(tid string) (*subscription.PubKeywordsRes, error) {
	return s.templateManagement.GetPubTemplateKeywords(tid)
}

// GetPubTemplateTitleList 获取所属类目下的公共模板
func (s *MessageManagement) GetPubTemplateTitleList(ids string, page, pageSize int) (*subscription.TemplateTitleListRes, error) {
	return s.templateManagement.GetPubTemplateTitles(ids, page, pageSize)
}

// GetMessageTemplateList 获取个人模板列表
func (s *MessageManagement) GetMessageTemplateList() (*subscription.PrivateTemplateRes, error) {
	return s.templateManagement.GetTemplate()
}

// AddMessageTemplate 添加模板
func (s *MessageManagement) AddMessageTemplate(params *subscription.AddTempParams) (string, error) {
	return s.templateManagement.AddTemplate(params)
}

// SendMessage 发送订阅消息
func (s *MessageManagement) SendMessage(params *SendMessageParams) error {
	_, err := s.request.JsonPost("/cgi-bin/message/subscribe/send", nil, params, nil)
	return err
}
