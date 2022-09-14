package user

import (
	"github.com/gmars/go-wechat/core"
)

type User struct {
	Request *core.ApiRequest
}

func NewUserManagement(token core.AccessToken) *User {
	return &User{Request: core.NewApiRequest(token)}
}

// CreateTag 创建标签
func (s *User) CreateTag(name string) (int, error) {
	var res CreateTagRes
	_, err := s.Request.JsonPost("/cgi-bin/tags/create", nil, map[string]map[string]string{
		"tag": {"name": name},
	}, &res)
	if err != nil {
		return 0, err
	}

	return res.Tag.Id, nil
}

// GetTags 获取所有标签
func (s *User) GetTags() (*TagsRes, error) {
	var res TagsRes
	_, err := s.Request.JsonGet("/cgi-bin/tags/get", nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UpdateTags 编辑标签
func (s *User) UpdateTags(id int, name string) error {
	_, err := s.Request.JsonPost("/cgi-bin/tags/update", nil, map[string]map[string]interface{}{
		"tag": {"id": id, "name": name},
	}, nil)
	return err
}

// DeleteTag 删除标签
func (s *User) DeleteTag(id int) error {
	_, err := s.Request.JsonPost("/cgi-bin/tags/delete", nil, map[string]map[string]int{
		"tag": {"id": id},
	}, nil)
	return err
}

// TagUsers 获取标签下的粉丝列表
func (s *User) TagUsers(tagId int, nextOpenid string) (*TagUsersRes, error) {
	var res TagUsersRes
	_, err := s.Request.JsonPost("/cgi-bin/user/tag/get", nil, map[string]interface{}{
		"tagid":       tagId,
		"next_openid": nextOpenid,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// BatchTagging 批量打标签
func (s *User) BatchTagging(tagId int, openIdList []string) error {
	_, err := s.Request.JsonPost("/cgi-bin/tags/members/batchtagging", nil, map[string]interface{}{
		"tagid":       tagId,
		"openid_list": openIdList,
	}, nil)

	return err
}

// BatchUntagging 批量取消打标签
func (s *User) BatchUntagging(tagId int, openIdList []string) error {
	_, err := s.Request.JsonPost("/cgi-bin/tags/members/batchuntagging", nil, map[string]interface{}{
		"tagid":       tagId,
		"openid_list": openIdList,
	}, nil)

	return err
}

// GetUserTags 获取用户所有标签
func (s *User) GetUserTags(openId string) ([]int, error) {
	var res TagsUserRes
	_, err := s.Request.JsonPost("/cgi-bin/tags/getidlist", nil, map[string]string{
		"openid": openId,
	}, &res)
	if err != nil {
		return nil, err
	}

	return res.TagIdList, nil
}

// UpdateRemark 设置用户备注名
func (s *User) UpdateRemark(openId, remark string) error {
	_, err := s.Request.JsonPost("/cgi-bin/user/info/updateremark", nil, map[string]string{
		"openid": openId,
		"remark": remark,
	}, nil)
	return err
}

// UserInfo 获取用户基本信息
func (s *User) UserInfo(openId, lang string) (*BaseInfo, error) {
	var res BaseInfo
	query := make(map[string]string)
	query["openid"] = openId
	query["lang"] = lang
	if lang == "" {
		query["lang"] = "zh_CN"
	}
	_, err := s.Request.JsonGet("/cgi-bin/user/info", query, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UserInfoListWithSameLang 批量获取用户信息相同的lang
func (s *User) UserInfoListWithSameLang(openIds []string, lang string) (*[]BaseInfo, error) {

	body := make([]map[string]string, 0)
	if lang == "" {
		lang = "zh_CN"
	}

	for _, item := range openIds {
		body = append(body, map[string]string{
			"openid": item,
			"lang":   lang,
		})
	}
	return s.userInfoListByOpenidAndLang(body)
}

// UserInfoListWithDiffLang 批量获取用户信息不同的lang
func (s *User) UserInfoListWithDiffLang(params []InfoListParamsItem) (*[]BaseInfo, error) {
	body := make([]map[string]string, 0)
	for _, item := range params {
		lang := item.Lang
		if lang == "" {
			lang = "zh_CN"
		}
		body = append(body, map[string]string{
			"openid": item.Openid,
			"lang":   lang,
		})
	}
	return s.userInfoListByOpenidAndLang(body)
}

// OpenIdPageList 分页获取openid列表
func (s *User) OpenIdPageList(nextOpenId string) (*OpenIdPageListRes, error) {
	var res OpenIdPageListRes
	_, err := s.Request.JsonGet("/cgi-bin/user/get", map[string]string{
		"next_openid": nextOpenId,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetBlackList 获取黑名单
func (s *User) GetBlackList(nextOpenId string) (*OpenIdPageListRes, error) {
	var res OpenIdPageListRes
	_, err := s.Request.JsonPost("/cgi-bin/tags/members/getblacklist", nil, map[string]string{
		"begin_openid": nextOpenId,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// BlackListOperation 加入/解除黑名单
func (s *User) BlackListOperation(openIdList []string, join bool) error {
	path := "/cgi-bin/tags/members/batchunblacklist"
	if join {
		path = "/cgi-bin/tags/members/batchblacklist"
	}
	_, err := s.Request.JsonPost(path, nil, map[string][]string{
		"openid_list": openIdList,
	}, nil)
	return err
}

// 获取用户列表信息
func (s *User) userInfoListByOpenidAndLang(body interface{}) (*[]BaseInfo, error) {
	var res InfoListRes
	_, err := s.Request.JsonPost("/cgi-bin/user/info/batchget", nil, body, &res)
	if err != nil {
		return nil, err
	}

	return &res.UserInfoList, nil
}
