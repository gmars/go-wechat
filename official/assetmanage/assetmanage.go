package assetmanage

import (
	"bytes"
	"encoding/json"
	"github.com/gmars/go-wechat/core"
	"github.com/gmars/go-wechat/util"
	"mime/multipart"
)

type AssetsManage struct {
	Request *core.ApiRequest
}

func NewAssetsManage(token core.AccessToken) *AssetsManage {
	return &AssetsManage{Request: core.NewApiRequest(token)}
}

// UploadTemporaryMaterials 上传临时文件
func (s *AssetsManage) UploadTemporaryMaterials(mediaType string, file *multipart.FileHeader) (*TemporaryUploadRes, error) {
	var res TemporaryUploadRes

	if err := s.validateMedia(mediaType, file, false); err != nil {
		return nil, err
	}

	_, err := s.Request.FormPost("/cgi-bin/media/upload", map[string]string{"type": mediaType},
		map[string]*multipart.FileHeader{"media": file}, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FetchTemporaryMaterial 临时媒体资源下载
func (s *AssetsManage) FetchTemporaryMaterial(mediaId string, isVoice bool) (*TemporaryDownloadRes, error) {
	var videoRes struct {
		VideoUrl string `json:"video_url"`
	}
	urlPath := "/cgi-bin/media/get"
	if isVoice {
		urlPath = "/cgi-bin/media/get/jssdk"
	}
	resp, err := s.Request.JsonGet(urlPath, map[string]string{
		"media_id": mediaId,
	}, nil)
	if err != nil {
		return nil, err
	}

	if bytes.Contains(resp, []byte("video_url")) {
		if err = json.Unmarshal(resp, &videoRes); err != nil {
			return nil, err
		}
		return &TemporaryDownloadRes{
			ResponseUrl:   true,
			ResponseValue: []byte(videoRes.VideoUrl),
		}, nil
	}

	return &TemporaryDownloadRes{
		ResponseUrl:   false,
		ResponseValue: resp,
	}, nil
}

// UploadArticleImage 上传图文中的图片
func (s *AssetsManage) UploadArticleImage(file *multipart.FileHeader) (string, error) {
	var res struct {
		Url string `json:"url"`
	}

	if err := s.validateMedia("article_image", file, false); err != nil {
		return "", err
	}

	_, err := s.Request.FormPost("/cgi-bin/media/uploadimg", nil, map[string]*multipart.FileHeader{"media": file}, nil, &res)
	if err != nil {
		return "", err
	}

	return res.Url, nil
}

// UploadPermanentMaterial 上传永久素材
func (s *AssetsManage) UploadPermanentMaterial(mediaType string, file *multipart.FileHeader, videoParams *PermanentVideoParams) (*PermanentUploadRes, error) {
	var (
		res  PermanentUploadRes
		body = make(map[string]string)
	)

	if err := s.validateMedia(mediaType, file, true); err != nil {
		return nil, err
	}

	if mediaType == MediaTypeVideo {
		if videoParams == nil {
			return nil, &core.ApiError{
				ErrCode: 500,
				ErrMsg:  "永久视频必须设置标题和描述信息",
			}
		} else if videoParams.Title == "" {
			return nil, &core.ApiError{
				ErrCode: 500,
				ErrMsg:  "永久视频标题不能为空",
			}
		} else if videoParams.Introduction == "" {
			return nil, &core.ApiError{
				ErrCode: 500,
				ErrMsg:  "用就是怕描述信息不能为空",
			}
		}

		description, err := json.Marshal(videoParams)
		if err != nil {
			return nil, err
		}
		body["description"] = string(description)
	}

	_, err := s.Request.FormPost("/cgi-bin/material/add_material", nil, map[string]*multipart.FileHeader{"media": file}, body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetPermanentMaterial 获取永久素材信息
func (s *AssetsManage) GetPermanentMaterial(mediaId string) (*PermanentMediaRes, error) {
	var res PermanentMediaRes
	resp, err := s.Request.JsonPost("/cgi-bin/material/get_material", nil, map[string]string{
		"media_id": mediaId,
	}, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}

	if bytes.Contains(resp, []byte("news_item")) {
		res.MediaType = "news"
	} else if bytes.Contains(resp, []byte("down_url")) {
		res.MediaType = MediaTypeVideo
	} else {
		res.MediaType = "other"
		res.Content = resp
	}

	return &res, nil
}

// DelPermanentMaterial 删除永久素材信息
func (s *AssetsManage) DelPermanentMaterial(mediaId string) error {
	_, err := s.Request.JsonPost("/cgi-bin/material/del_material", nil, map[string]string{
		"media_id": mediaId,
	}, nil)
	return err
}

// MaterialCount 素材总数
func (s *AssetsManage) MaterialCount() (*MaterialCountRes, error) {
	var res MaterialCountRes
	_, err := s.Request.JsonGet("/cgi-bin/material/get_materialcount", nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetMaterialList 获取素材列表
func (s *AssetsManage) GetMaterialList(mediaType string, page, pageSize int) (*MaterialListRes, error) {
	var res MaterialListRes
	if mediaType == "" {
		return nil, &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "缺少必要参数",
		}
	}

	if mediaType == "news" {
		return nil, &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "查询图文列表请使用方法GetNewsList",
		}
	}
	offset, count := util.PageCondition(page, pageSize, 20)
	_, err := s.Request.JsonPost("/cgi-bin/material/batchget_material", nil, map[string]interface{}{
		"type":   mediaType,
		"offset": offset,
		"count":  count,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetNewsList 获取News素材列表
func (s *AssetsManage) GetNewsList(page, pageSize int) (*NewsListRes, error) {
	var res NewsListRes

	offset, count := util.PageCondition(page, pageSize, 20)
	_, err := s.Request.JsonPost("/cgi-bin/material/batchget_material", nil, map[string]interface{}{
		"type":   "news",
		"offset": offset,
		"count":  count,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
