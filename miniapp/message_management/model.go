package message_management

type WeAppTemplateData struct {
	Value string `json:"value"`
}

type WeAppTemplateMsg struct {
	TemplateId      string                       `json:"template_id"`
	Page            string                       `json:"page"`
	FormId          string                       `json:"form_id"`
	Data            map[string]WeAppTemplateData `json:"data"`
	EmphasisKeyword string                       `json:"emphasis_keyword"`
}

type MpTemplateData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type MpTemplateMsg struct {
	Appid       string `json:"appid"`
	TemplateId  string `json:"template_id"`
	Url         string `json:"url"`
	MiniProgram struct {
		Appid    string `json:"appid"`
		PagePath string `json:"pagepath"`
	} `json:"miniprogram"`
	Data map[string]MpTemplateData `json:"data"`
}

type SendUniformMessageParams struct {
	ToUser           string           `json:"touser"`
	WeAppTemplateMsg WeAppTemplateMsg `json:"weapp_template_msg"`
	MpTemplateMsg    MpTemplateMsg    `json:"mp_template_msg"`
}

type CreateActivityIdRes struct {
	ActivityId     string `json:"activity_id"`
	ExpirationTime int    `json:"expiration_time"`
}

type SetUpdatableMsgParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SetUpdatableMsgRes struct {
	ActivityId   string `json:"activity_id"`
	TargetState  int    `json:"target_state"`
	TemplateInfo struct {
		ParameterList []SetUpdatableMsgParameter `json:"parameter_list"`
	} `json:"template_info"`
}

type SendMessageParams struct {
	ToUser           string                       `json:"touser"`
	TemplateId       string                       `json:"template_id"`
	Page             string                       `json:"page"`
	MiniProgramState string                       `json:"miniprogram_state"`
	Lang             string                       `json:"lang"`
	Data             map[string]WeAppTemplateData `json:"data"`
}
