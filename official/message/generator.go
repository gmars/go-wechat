package message

import (
	"fmt"
	"time"
)

// GenerateText 文本消息
func GenerateText(fromUser, toUser, content string) string {
	return fmt.Sprintf(
		"<xml>\n<ToUserName><![CDATA[%s]]></ToUserName>\n<FromUserName><![CDATA[%s]]></FromUserName>\n<CreateTime>%d</CreateTime>\n<MsgType><![CDATA[text]]></MsgType>\n<Content><![CDATA[%s]]></Content>\n</xml>",
		toUser, fromUser, time.Now().Unix(), content)
}

// GenerateImage 图片消息
func GenerateImage(fromUser, toUser, mediaId string) string {
	return fmt.Sprintf(
		"<xml>\n<ToUserName><![CDATA[%s]]></ToUserName>\n<FromUserName><![CDATA[%s]]></FromUserName>\n<CreateTime>%d</CreateTime>\n<MsgType><![CDATA[image]]></MsgType>\n<Image>\n<MediaId><![CDATA[%s]]></MediaId>\n</Image>\n</xml>",
		toUser, fromUser, time.Now().Unix(), mediaId)
}

// GenerateVoice 语音消息
func GenerateVoice(fromUser, toUser, mediaId string) string {
	return fmt.Sprintf(
		"<xml>\n<ToUserName><![CDATA[%s]]></ToUserName>\n<FromUserName><![CDATA[%s]]></FromUserName>\n<CreateTime>%d</CreateTime>\n<MsgType><![CDATA[voice]]></MsgType>\n<Image>\n<MediaId><![CDATA[%s]]></MediaId>\n</Image>\n</xml>",
		toUser, fromUser, time.Now().Unix(), mediaId)
}

// GenerateVideo 视频消息
func GenerateVideo(fromUser, toUser, mediaId, title, description string) string {
	return fmt.Sprintf(
		"<xml>\n<ToUserName><![CDATA[%s]]></ToUserName>\n<FromUserName><![CDATA[%s]]></FromUserName>\n<CreateTime>%d</CreateTime>\n<MsgType><![CDATA[video]]></MsgType>\n<Video>\n<MediaId><![CDATA[%s]]></MediaId>\n<Title><![CDATA[%s]]></Title>\n<Description><![CDATA[%s]]></Description>\n</Video>\n</xml>",
		toUser, fromUser, time.Now().Unix(), mediaId, title, description)
}

// GenerateMusic 音乐消息
func GenerateMusic(fromUser, toUser, coverImgMediaId, title, description, url, hqUrl string) string {
	return fmt.Sprintf(
		"<xml>\n<ToUserName><![CDATA[%s]]></ToUserName>\n<FromUserName><![CDATA[%s]]></FromUserName>\n<CreateTime>%d</CreateTime>\n<MsgType><![CDATA[music]]></MsgType>\n<Music>\n<Title><![CDATA[%s]]></Title>\n<Description><![CDATA[%s]]></Description>\n<MusicUrl><![CDATA[%s]]></MusicUrl>\n<HQMusicUrl><![CDATA[%s]]></HQMusicUrl>\n<ThumbMediaId><![CDATA[%s]]></ThumbMediaId>\n</Music>\n</xml>",
		toUser, fromUser, time.Now().Unix(), title, description, url, hqUrl, coverImgMediaId)
}

// GenerateNews 图文消息
func GenerateNews(fromUser, toUser string, list []NewsItem) string {
	var (
		msg = fmt.Sprintf(
			"<xml>\n<ToUserName><![CDATA[%s]]></ToUserName>\n<FromUserName><![CDATA[%s]]></FromUserName>\n<CreateTime>%d</CreateTime>\n<MsgType><![CDATA[news]]></MsgType>\n<ArticleCount>%d</ArticleCount>\n<Articles>",
			toUser, fromUser, time.Now().Unix(), len(list))
	)

	for _, item := range list {
		msg += fmt.Sprintf(
			"<item>\n<Title><![CDATA[%s]]></Title>\n<Description><![CDATA[%s]]></Description>\n<PicUrl><![CDATA[%s]]></PicUrl>\n<Url><![CDATA[%s]]></Url>\n</item>",
			item.Title, item.Description, item.PicUrl, item.Url)
	}
	msg += "</Articles>\n</xml>"
	return msg
}
