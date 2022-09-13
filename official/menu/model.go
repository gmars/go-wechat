package menu

type NewsInfo struct {
	Title      string `json:"title"`
	Digest     string `json:"digest"`
	Author     string `json:"author"`
	ShowCover  string `json:"show_cover"`
	CoverUrl   string `json:"cover_url"`
	ContentUrl string `json:"content_url"`
	SourceUrl  string `json:"source_url"`
}

type BaseButton struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Key  string `json:"key"`
}

type ButtonItem struct {
	BaseButton
	Value    string `json:"value"`
	NewsInfo struct {
		List []NewsInfo `json:"list"`
	} `json:"news_info"`
	SubButton struct {
		List []ButtonItem `json:"list"`
	} `json:"sub_button"`
}

type ButtonRes struct {
	IsMenuOpen   int `json:"is_menu_open"`
	SelfMenuInfo struct {
		Button []ButtonItem `json:"button"`
	} `json:"selfmenu_info"`
}

type CreateButtonItem struct {
	BaseButton
	MediaId   string             `json:"media_id"`
	Appid     string             `json:"appid"`
	Pagepath  string             `json:"pagepath"`
	ArticleId string             `json:"article_id"`
	SubButton []CreateButtonItem `json:"sub_button"`
}

type CreateMenuParams struct {
	Button []CreateButtonItem `json:"button"`
}

type ButtonMatchRule struct {
	TagId              string `json:"tag_id"`
	ClientPlatformType string `json:"client_platform_type"`
}

type CreateConditionalParams struct {
	Button    []CreateButtonItem `json:"button"`
	Matchrule ButtonMatchRule    `json:"matchrule"`
}

type GetMenuRes struct {
	Menu struct {
		Button []CreateButtonItem `json:"button"`
		Menuid string             `json:"menuid"`
	} `json:"menu"`
	Conditionalmenu []struct {
		Button    []CreateButtonItem `json:"button"`
		Matchrule ButtonMatchRule    `json:"matchrule"`
		Menuid    string             `json:"menuid"`
	} `json:"conditionalmenu"`
}
