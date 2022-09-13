package subscribe

type SubscribeBizFlag = int

const (
	SubscribeBizFlagOpen  SubscribeBizFlag = 1
	SubscribeBizFlagClose SubscribeBizFlag = 0
)

type GetShowItemRes struct {
	CanOpen  int    `json:"can_open"`
	IsOpen   int    `json:"is_open"`
	Appid    string `json:"appid"`
	Nickname string `json:"nickname"`
	Headimg  string `json:"headimg"`
}

type GetLinkForShowRes struct {
	TotalNum    int `json:"total_num"`
	BizInfoList []struct {
		Nickname string `json:"nickname"`
		Appid    string `json:"appid"`
		Headimg  string `json:"headimg"`
	} `json:"biz_info_list"`
}
