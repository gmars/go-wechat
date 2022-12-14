package assetmanage

import (
	"github.com/gmars/go-wechat/core"
	"mime/multipart"
	"path"
	"strconv"
	"strings"
)

const (
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
)

type TemporaryUploadRes struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
}

type TemporaryDownloadRes struct {
	ResponseUrl   bool
	ResponseValue []byte
}

type PermanentVideoParams struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}

type PermanentUploadRes struct {
	MediaId string `json:"media_id"`
	Url     string `json:"url"`
}

type NewsItem struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"content_source_url"`
}

type PermanentMediaRes struct {
	MediaType   string     `json:"media_type"`
	NewsItem    []NewsItem `json:"news_item"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DownUrl     string     `json:"down_url"`
	Content     []byte     `json:"content"`
}

type MaterialCountRes struct {
	VoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"`
}

type MaterialListRes struct {
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	Item       []struct {
		MediaId    string `json:"media_id"`
		Name       string `json:"name"`
		UpdateTime int    `json:"update_time"`
		Url        string `json:"url"`
	} `json:"item"`
}

type NewsListRes struct {
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	Item       []struct {
		MediaId string `json:"media_id"`
		Content struct {
			NewsItem []struct {
				Title            string `json:"title"`
				ThumbMediaId     string `json:"thumbMediaId"`
				ShowCoverPic     int    `json:"showCoverPic"`
				Author           string `json:"author"`
				Digest           string `json:"digest"`
				Content          string `json:"content"`
				Url              string `json:"url"`
				ContentSourceUrl string `json:"contentSourceUrl"`
			} `json:"news_item"`
		} `json:"content"`
		UpdateTime int `json:"update_time"`
	} `json:"item"`
}

// ?????????????????????
func (s *AssetsManage) validateMedia(mediaType string, file *multipart.FileHeader, permanent bool) error {
	fileExt := path.Ext(file.Filename)
	sizeLimit := int64(0)
	typeLimit := ""
	switch mediaType {
	case MediaTypeImage:
		sizeLimit = 10 << 20
		typeLimit = ".png.jpeg.jpg.gif"
		if permanent {
			typeLimit += ".bmp"
		}
		break
	case MediaTypeThumb:
		sizeLimit = 64 << 10
		typeLimit = ".jpg"
		break
	case MediaTypeVideo:
		sizeLimit = 10 << 20
		typeLimit = ".mp4"
		break
	case MediaTypeVoice:
		sizeLimit = 2 << 20
		typeLimit = ".amr.mp3"
		if permanent {
			typeLimit += ".wma.wav"
		}
		break
	case "article_image":
		sizeLimit = 1 << 20
		typeLimit = ".png.jpeg.jpg.gif"
		if permanent {
			typeLimit += ".bmp"
		}
		break
	default:
		return &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "????????????????????????",
		}
	}

	if !strings.Contains(typeLimit, fileExt) {
		return &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "????????????????????????????????????????????????????????????????????????" + typeLimit,
		}
	}

	if file.Size > sizeLimit {
		return &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "???????????????????????????????????????????????????????????????????????????" + strconv.FormatInt(sizeLimit>>20, 10) + "MB",
		}
	}

	return nil
}
