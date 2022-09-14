package website

type LangType = string
type StyleType = string

const (
	LangCn     LangType  = "cn"
	LangEn     LangType  = "en"
	StyleBlack StyleType = "black"
	StyleWhite StyleType = "white"
)

type LoginConfigRes struct {
	SelfRedirect bool      `json:"self_redirect"`
	Id           string    `json:"id"`
	AppId        string    `json:"appid"`
	Scope        string    `json:"scope"`
	RedirectUri  string    `json:"redirect_uri"`
	State        string    `json:"state"`
	Style        StyleType `json:"style"`
	Href         string    `json:"href"`
}

type AccessTokenRes struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"unionid"`
}

type UserInfoRes struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionId    string   `json:"unionid"`
}
