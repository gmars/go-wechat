package template_msg

type GetIndustryRes struct {
	PrimaryIndustry struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
	} `json:"primary_industry"`
	SecondaryIndustry struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
	} `json:"secondary_industry"`
}

type AddTemplateRes struct {
	TemplateId string `json:"template_id"`
}

type TemplateListRes struct {
	TemplateList []TemplateInfo `json:"template_list"`
}

type TemplateInfo struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type MsgDataItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type MiniProgram struct {
	Appid    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

type Message struct {
	TouSer      string                 `json:"touser"`
	TemplateId  string                 `json:"template_id"`
	Url         string                 `json:"url"`
	MiniProgram *MiniProgram           `json:"miniprogram"`
	ClientMsgId string                 `json:"client_msg_id"`
	Data        map[string]MsgDataItem `json:"data"`
}

type SendRes struct {
	MsgId int `json:"msgid"`
}
