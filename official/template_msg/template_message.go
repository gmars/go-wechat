package template_msg

import "github.com/gmars/go-wechat/core"

type TemplateMessage struct {
	request *core.ApiRequest
}

func NewTemplateMessage(token core.AccessToken) *TemplateMessage {
	return &TemplateMessage{request: core.NewApiRequest(token)}
}

// SetIndustry 设置所属行业
func (t *TemplateMessage) SetIndustry(industryIdOne, industryIdTwo int) error {
	_, err := t.request.JsonPost("/cgi-bin/template/get_industry", nil, map[string]int{
		"industry_id1": industryIdOne,
		"industry_id2": industryIdTwo,
	}, nil)
	return err
}

// GetIndustry 获取所属行业
func (t *TemplateMessage) GetIndustry() (*GetIndustryRes, error) {
	var res GetIndustryRes
	_, err := t.request.JsonGet("/cgi-bin/template/get_industry", nil, &res)
	return &res, err
}

// AddTemplate 添加模板
func (t *TemplateMessage) AddTemplate(shortId string) (string, error) {
	var res AddTemplateRes
	_, err := t.request.JsonPost("/cgi-bin/template/api_add_template", nil, map[string]string{
		"template_id_short": shortId,
	}, &res)
	return res.TemplateId, err
}

// GetTemplateList 获取已添加模板列表
func (t *TemplateMessage) GetTemplateList() ([]TemplateInfo, error) {
	var res TemplateListRes
	_, err := t.request.JsonGet("/cgi-bin/template/get_all_private_template", nil, &res)
	return res.TemplateList, err
}

// DelTemplate 删除模板
func (t *TemplateMessage) DelTemplate(templateId string) error {
	_, err := t.request.JsonPost("/cgi-bin/template/del_private_template", nil, map[string]string{
		"template_id": templateId,
	}, nil)
	return err
}

// Send 发送模板消息
func (t *TemplateMessage) Send(message *Message) (int, error) {
	var res SendRes
	_, err := t.request.JsonPost("/cgi-bin/message/template/send", nil, message, &res)
	return res.MsgId, err
}
