package kf_message

import (
	"github.com/gmars/go-wechat/core"
	"mime/multipart"
)

type KfMessage struct {
	request *core.ApiRequest
}

func NewKfMessage(accessToken core.AccessToken) *KfMessage {
	return &KfMessage{request: core.NewApiRequest(accessToken)}
}

// UploadTempMedia 上传临时文件，只支持图片
func (s *KfMessage) UploadTempMedia(file *multipart.FileHeader) (*UploadRes, error) {
	var res UploadRes
	_, err := s.request.FormPost("/cgi-bin/media/upload", map[string]string{"type": "image"},
		map[string]*multipart.FileHeader{"media": file}, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetTempMedia 获取临时文件
func (s *KfMessage) GetTempMedia(mediaId string) ([]byte, error) {
	var res GetTempMediaRes
	_, err := s.request.JsonGet("/cgi-bin/media/get", map[string]string{"media_id": "mediaId"}, &res)
	if err != nil {
		return nil, err
	}

	return res.Buffer, nil
}

// SetTyping 下发客服当前输入状态
func (s *KfMessage) SetTyping(toUser, command string) error {
	_, err := s.request.JsonPost("/cgi-bin/message/custom/business/typing", nil, map[string]interface{}{
		"touser":  toUser,
		"command": command,
	}, nil)
	return err
}

// SendText 下发文字
func (s *KfMessage) SendText(toUser, content string) error {
	msg := map[string]interface{}{
		"touser":  toUser,
		"msgtype": "text",
		"text": map[string]string{
			"content": content,
		},
	}
	return s.send(msg)
}

// SendImage 下发图片
func (s *KfMessage) SendImage(toUser, mediaId string) error {
	msg := map[string]interface{}{
		"touser":  toUser,
		"msgtype": "image",
		"image": map[string]string{
			"media_id": mediaId,
		},
	}
	return s.send(msg)
}

// SendLink 下发图文链接
func (s *KfMessage) SendLink(toUser string, linkMsg *LinkMsg) error {
	msg := map[string]interface{}{
		"touser":  toUser,
		"msgtype": "link",
		"link":    linkMsg,
	}
	return s.send(msg)
}

// SendMiniProgramPage 下发小程序卡片
func (s *KfMessage) SendMiniProgramPage(toUser string, miniProgramPage *MiniProgramPageMsg) error {
	msg := map[string]interface{}{
		"touser":          toUser,
		"msgtype":         "miniprogrampage",
		"miniprogrampage": miniProgramPage,
	}
	return s.send(msg)
}

// SetTyping 下发客服当前输入状态
func (s *KfMessage) send(msg map[string]interface{}) error {
	_, err := s.request.JsonPost("/cgi-bin/message/custom/business/send", nil, msg, nil)
	return err
}
