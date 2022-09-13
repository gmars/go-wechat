package user

type CreateTagRes struct {
	Tag struct {
		Id int `json:"id"`
	} `json:"tag"`
}

type TagsRes struct {
	Tags []struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"tags"`
}

type TagUsersRes struct {
	Count int64 `json:"count"`
	Data  struct {
		OpenId []string `json:"openid"`
	}
	NextOpenId string `json:"next_openid"`
}

type TagsUserRes struct {
	TagIdList []int `json:"tagid_list"`
}

type BaseInfo struct {
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Language       string `json:"language"`
	SubscribeTime  int    `json:"subscribe_time"`
	UnionId        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupId        int    `json:"groupid"`
	TagIdList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

type InfoListParamsItem struct {
	Openid string `json:"openid"`
	Lang   string `json:"lang"`
}

type InfoListRes struct {
	UserInfoList []BaseInfo `json:"user_info_list"`
}

type OpenIdPageListRes struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}
