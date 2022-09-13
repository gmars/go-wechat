package member

import "go-wechat/core"

type Member struct {
	request *core.ApiRequest
}

func NewMemberManagement(authorizerAccessToken core.AccessToken) *Member {
	return &Member{request: core.NewApiRequest(authorizerAccessToken)}
}

func (s *Member) BindTester(wechatId string) (string, error) {
	var res BindTesterRes
	_, err := s.request.JsonPost("/wxa/bind_tester", nil, map[string]string{
		"wechatid": wechatId,
	}, &res)
	return res.UserStr, err
}

// UnbindTester 解除绑定体验者
func (s *Member) UnbindTester(wechatId, userStr string) error {
	_, err := s.request.JsonPost("/wxa/unbind_tester", nil, map[string]string{
		"wechatid": wechatId,
		"userstr":  userStr,
	}, nil)
	return err
}

// GetTester 获取体验者列表
func (s *Member) GetTester(action string) (*[]Members, error) {
	var res GetTesterRes
	if action == "" {
		action = "get_experiencer"
	}
	_, err := s.request.JsonPost("/wxa/memberauth", nil, map[string]string{
		"action": action,
	}, &res)
	return &res.Members, err
}
