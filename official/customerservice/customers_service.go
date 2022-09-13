package customerservice

import (
	"go-wechat/core"
	"mime/multipart"
	"regexp"
	"unicode/utf8"
)

type CustomerService struct {
	Request *core.ApiRequest
}

func NewCustomerService(token core.AccessToken) *CustomerService {
	return &CustomerService{Request: core.NewApiRequest(token)}
}

// KfList 获取客服列表
func (s *CustomerService) KfList() (*[]KfInfo, error) {
	var res KfListRes
	_, err := s.Request.JsonGet("/cgi-bin/customservice/getkflist", nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.KfList, nil
}

// OnlineKfList 获取在线客服列表
func (s *CustomerService) OnlineKfList() (*[]KfOnlineInfo, error) {
	var res OnlineKfRes
	_, err := s.Request.JsonGet("/cgi-bin/customservice/getonlinekflist", nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.KfOnlineList, nil
}

// AddKfAccount 添加客服账号
func (s *CustomerService) AddKfAccount(account, wxOriginAccount, nickName string) error {
	rg := regexp.MustCompile(`^[a-zA-Z0-9_]{1,10}$`)
	if !rg.MatchString(account) {
		return &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "客服账号只能包含字母数字或下划线且长度不能超过10个字符",
		}
	}

	if utf8.RuneCountInString(nickName) > 16 {
		return &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "客服昵称不能超过16个字",
		}
	}

	_, err := s.Request.JsonPost("/customservice/kfaccount/add", nil, map[string]string{
		"kf_account": account + "@" + wxOriginAccount,
		"nickname":   nickName,
	}, nil)
	return err
}

// InviteKfWorker 邀请客服
func (s *CustomerService) InviteKfWorker(kfAccount, inviteWx string) error {
	_, err := s.Request.JsonPost("/customservice/kfaccount/inviteworker", nil, map[string]string{
		"kf_account": kfAccount,
		"invite_wx":  inviteWx,
	}, nil)
	return err
}

// UpdateKfAccount 修改客服账号信息
func (s *CustomerService) UpdateKfAccount(kfAccount, nickname string) error {
	if utf8.RuneCountInString(nickname) > 16 {
		return &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "客服昵称不能超过16个字",
		}
	}

	_, err := s.Request.JsonPost("/customservice/kfaccount/update", nil, map[string]string{
		"kf_account": kfAccount,
		"nickname":   nickname,
	}, nil)
	return err
}

// KfHeadImgUpload 上传客服头像
func (s *CustomerService) KfHeadImgUpload(kfAccount string, file *multipart.FileHeader) error {
	if file.Size > 5<<20 {
		return &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "头像大小不能超过5MB",
		}
	}
	_, err := s.Request.FormPost("/customservice/kfaccount/uploadheadimg", map[string]string{
		"kf_account": kfAccount,
	}, map[string]*multipart.FileHeader{
		"media": file,
	}, nil, nil)
	return err
}

// DelKfAccount 删除客服账号
func (s *CustomerService) DelKfAccount(kfAccount string) error {
	_, err := s.Request.JsonGet("/customservice/kfaccount/del", map[string]string{
		"kf_account": kfAccount,
	}, nil)
	return err
}

// DialogueManage 会话管理
func (s *CustomerService) DialogueManage(kfAccount string, openId string, close bool) error {
	path := "/customservice/kfsession/create"
	if close {
		path = "/customservice/kfsession/close"
	}
	_, err := s.Request.JsonPost(path, nil, map[string]string{
		"kf_account": kfAccount,
		"openid":     openId,
	}, nil)
	return err
}

// GetUserDialogueInfo 获取会话状态
func (s *CustomerService) GetUserDialogueInfo(openId string) (*DialogStatus, error) {
	var res DialogStatus
	_, err := s.Request.JsonGet("/customservice/kfsession/getsession", map[string]string{
		"openid": openId,
	}, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetKfDialogueRecords 获取客服会话记录
func (s *CustomerService) GetKfDialogueRecords(kfAccount string) (*[]KfDialogItem, error) {
	var res KfDialogRecords
	_, err := s.Request.JsonGet("/customservice/kfsession/getsessionlist", map[string]string{
		"kf_account": kfAccount,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res.SessionList, nil
}

// WaitingList 获取待接入列表
func (s *CustomerService) WaitingList() (*WaitingRes, error) {
	var res WaitingRes
	_, err := s.Request.JsonGet("/customservice/kfsession/getwaitcase", nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ChattingRecords 获取聊天记录
func (s *CustomerService) ChattingRecords(params *ChartRecordsParams) (*RecordsList, error) {
	var res RecordsList
	_, err := s.Request.JsonPost("/customservice/msgrecord/getmsglist", nil, params, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// KfTyping 客服输入状态
func (s *CustomerService) KfTyping(toUserOpenId string, typing bool) error {
	body := map[string]string{
		"touser":  toUserOpenId,
		"command": "CancelTyping",
	}
	if typing {
		body["command"] = "Typing"
	}
	_, err := s.Request.JsonPost("/cgi-bin/message/custom/typing", nil, body, nil)
	return err
}

// SendText 客服发送文字消息
func (s *CustomerService) SendText(toUserOpenId string, content string, kfAccount string) error {

	return s.sendMsg(map[string]interface{}{
		"touser":  toUserOpenId,
		"msgtype": "text",
		"text": map[string]string{
			"content": content,
		},
	}, kfAccount)
}

// SendImage 客服发送图片消息
func (s *CustomerService) SendImage(toUserOpenId string, mediaId string, kfAccount string) error {
	return s.sendMsg(map[string]interface{}{
		"touser":  toUserOpenId,
		"msgtype": "image",
		"image": map[string]string{
			"media_id": mediaId,
		},
	}, kfAccount)
}

// SendVoice 客服发送语音消息
func (s *CustomerService) SendVoice(toUserOpenId string, mediaId string, kfAccount string) error {
	return s.sendMsg(map[string]interface{}{
		"touser":  toUserOpenId,
		"msgtype": "voice",
		"voice": map[string]string{
			"media_id": mediaId,
		},
	}, kfAccount)
}

// SendVideo 客服发送视频消息
func (s *CustomerService) SendVideo(toUserOpenId string, params *VideoMsgParams, kfAccount string) error {
	return s.sendMsg(map[string]interface{}{
		"touser":  toUserOpenId,
		"msgtype": "video",
		"video":   params,
	}, kfAccount)
}

// SendMusic 客服发送音乐消息
func (s *CustomerService) SendMusic(toUserOpenId string, params *MusicMsgParams, kfAccount string) error {
	return s.sendMsg(map[string]interface{}{
		"touser":  toUserOpenId,
		"msgtype": "music",
		"music":   params,
	}, kfAccount)
}

// SendLinkArticle 发送外部图文
func (s *CustomerService) SendLinkArticle(toUserOpenId string, params LinkArticlesParams, kfAccount string) error {
	return s.sendMsg(map[string]interface{}{
		"touser":  toUserOpenId,
		"msgtype": "news",
		"news": map[string][]LinkArticlesParams{
			"articles": {params},
		},
	}, kfAccount)
}

// SendMpArticle 发送公众号图文
func (s *CustomerService) SendMpArticle(toUserOpenId string, articleId string, kfAccount string) error {
	return s.sendMsg(map[string]interface{}{
		"touser":  toUserOpenId,
		"msgtype": "mpnewsarticle",
		"mpnewsarticle": map[string]string{
			"article_id": articleId,
		},
	}, kfAccount)
}

// SendMenu 发送操作菜单
func (s *CustomerService) SendMenu(toUserOpenId string, head string, list []MenuMsgListItem, tail string, kfAccount string) error {
	return s.sendMsg(map[string]interface{}{
		"touser":  toUserOpenId,
		"msgtype": "msgmenu",
		"msgmenu": map[string]interface{}{
			"head_content": head,
			"list":         list,
			"tail_content": tail,
		},
	}, kfAccount)
}

// 发送消息给用户
func (s *CustomerService) sendMsg(body map[string]interface{}, kfAccount string) error {
	if kfAccount != "" {
		body["customservice"] = map[string]string{
			"kf_account": kfAccount,
		}
	}
	_, err := s.Request.JsonPost("/cgi-bin/message/custom/send", nil, body, nil)
	return err
}
