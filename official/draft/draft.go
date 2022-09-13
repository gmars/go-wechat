package draft

import (
	"go-wechat/core"
	"go-wechat/util"
)

type Draft struct {
	Request *core.ApiRequest
}

func NewDraft(token core.AccessToken) *Draft {
	return &Draft{Request: core.NewApiRequest(token)}
}

// AddDraft 创建草稿箱
func (s *Draft) AddDraft(articleList *[]ArticleItem) (string, error) {
	var res AddArticleRes
	if err := validateArticles(articleList); err != nil {
		return "", nil
	}

	_, err := s.Request.JsonPost("/cgi-bin/draft/add", nil, map[string]interface{}{
		"articles": articleList,
	}, &res)
	if err != nil {
		return "", err
	}

	return res.MediaId, nil
}

// GetDraft 获取草稿
func (s *Draft) GetDraft(mediaId string) (*NewsRes, error) {
	var res NewsRes
	_, err := s.Request.JsonPost("/cgi-bin/draft/get", nil, map[string]string{
		"media_id": mediaId,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteDraft 删除草稿
func (s *Draft) DeleteDraft(mediaId string) error {
	_, err := s.Request.JsonPost("/cgi-bin/draft/delete", nil, map[string]string{
		"media_id": mediaId,
	}, nil)
	return err
}

// UpdateDraft 更新草稿
func (s *Draft) UpdateDraft(params *UpdateArticle) error {
	_, err := s.Request.JsonPost("/cgi-bin/draft/update", nil, params, nil)
	return err
}

// DraftCount 草稿箱统计
func (s *Draft) DraftCount() (int, error) {
	var res CountRes
	_, err := s.Request.JsonPost("/cgi-bin/draft/count", nil, nil, &res)
	if err != nil {
		return 0, err
	}

	return res.TotalCount, nil
}

// DraftPageList 草稿箱列表
func (s *Draft) DraftPageList(page, pageSize int, noContent bool) (*PageListDraftRes, error) {
	var res PageListDraftRes
	body := make(map[string]int)
	offset, count := util.PageCondition(page, pageSize, 20)
	body["offset"] = offset
	body["count"] = count
	body["no_content"] = 0
	if noContent {
		body["no_content"] = 1
	}
	_, err := s.Request.JsonPost("/cgi-bin/draft/batchget", nil, body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SubmitPublish 提交发布
func (s *Draft) SubmitPublish(mediaId string) (string, error) {
	var res SubmitPublishRes
	_, err := s.Request.JsonPost("/cgi-bin/freepublish/submit", nil, map[string]string{
		"media_id": mediaId,
	}, &res)
	if err != nil {
		return "", err
	}

	return res.PublishId, nil
}

// GetPublishStatus 查询发布状态
func (s *Draft) GetPublishStatus(publishId string) (*GetPublishStatusRes, error) {
	var res GetPublishStatusRes
	_, err := s.Request.JsonPost("/cgi-bin/freepublish/get", nil, map[string]string{
		"publish_id": publishId,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeletePublish 删除已发布文章
func (s *Draft) DeletePublish(articleId string, index int) error {
	_, err := s.Request.JsonPost("/cgi-bin/freepublish/delete", nil, map[string]interface{}{
		"article_id": articleId,
		"index":      index,
	}, nil)
	return err
}

// GetPublishedArticle 获取已发布文章
func (s *Draft) GetPublishedArticle(articleId string) (*GetPublishedArticle, error) {
	var res GetPublishedArticle
	_, err := s.Request.JsonPost("/cgi-bin/freepublish/getarticle", nil, map[string]interface{}{
		"article_id": articleId,
	}, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetPublishedList 获取已发布文章列表
func (s *Draft) GetPublishedList(page, pageSize int, noContent bool) (*GetPublishedList, error) {
	var res GetPublishedList
	body := make(map[string]int)
	offset, count := util.PageCondition(page, pageSize, 20)
	body["offset"] = offset
	body["count"] = count
	body["no_content"] = 0
	if noContent {
		body["no_content"] = 1
	}

	_, err := s.Request.JsonPost("/cgi-bin/freepublish/batchget", nil, body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
