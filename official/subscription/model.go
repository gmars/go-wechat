package subscription

type AddTempParams struct {
	Tid       string `json:"tid"`
	KidList   []int  `json:"kidList"`
	SceneDesc string `json:"sceneDesc"`
}

type AddTempRes struct {
	PriTmplId string `json:"priTmplId"`
}

type CategoryItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryRest struct {
	Data []CategoryItem `json:"data"`
}

type PubKeywordsRes struct {
	Count int `json:"count"`
	Data  []struct {
		Kid     int64  `json:"kid"`
		Name    string `json:"name"`
		Example string `json:"example"`
		Rule    string `json:"rule"`
	} `json:"data"`
}

type TemplateTitleListRes struct {
	Count int `json:"count"`
	Data  []struct {
		Tid        int64  `json:"tid"`
		Title      string `json:"title"`
		Type       int    `json:"type"`
		CategoryId int    `json:"categoryId"`
	} `json:"data"`
}

type PrivateTemplateRes struct {
	Data []struct {
		PriTmplId string `json:"priTmplId"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Example   string `json:"example"`
		Type      int    `json:"type"`
	} `json:"data"`
}

type MiniAppParams struct {
	Appid    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

type DataValue struct {
	Value string `json:"value"`
}

type SendParams struct {
	ToUser      string               `json:"touser"`
	TemplateId  string               `json:"template_id"`
	Page        string               `json:"page"`
	MiniProgram MiniAppParams        `json:"miniprogram"`
	Data        map[string]DataValue `json:"data"`
}
