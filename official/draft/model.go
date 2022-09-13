package draft

import (
	"fmt"
	"go-wechat/core"
)

type ArticleItem struct {
	Title              string `json:"title"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	ContentSourceUrl   string `json:"content_source_url"`
	ThumbMediaId       string `json:"thumb_media_id"`
	NeedOpenComment    uint32 `json:"need_open_comment"`
	OnlyFansCanComment uint32 `json:"only_fans_can_comment"`
}

type AddArticleRes struct {
	MediaId string `json:"media_id"`
}

type NewsItem struct {
	Title              string `json:"title"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	ContentSourceUrl   string `json:"content_source_url"`
	ThumbMediaId       string `json:"thumb_media_id"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
	Url                string `json:"url"`
}

type NewsRes struct {
	NewsItem []NewsItem `json:"news_item"`
}

type UpdateArticle struct {
	MediaId  string        `json:"media_id"`
	Index    int           `json:"index"`
	Articles []ArticleItem `json:"articles"`
}

type CountRes struct {
	TotalCount int `json:"total_count"`
}

type ArticlesItem struct {
	MediaId string `json:"media_id"`
	Content struct {
		NewsItem []NewsItem `json:"news_item"`
	} `json:"content"`
	UpdateTime int `json:"update_time"`
}

type PageListDraftRes struct {
	TotalCount int            `json:"total_count"`
	ItemCount  int            `json:"item_count"`
	Item       []ArticlesItem `json:"item"`
}

// SubmitPublishRes 提交发布
type SubmitPublishRes struct {
	PublishId string `json:"publish_id"`
}

// GetPublishStatusRes 查询发布状态
type GetPublishStatusRes struct {
	PublishId     string `json:"publish_id"`
	PublishStatus int    `json:"publish_status"`
	ArticleId     string `json:"article_id"`
	ArticleDetail struct {
		Count int `json:"count"`
		Item  []struct {
			Idx        int    `json:"idx"`
			ArticleUrl string `json:"article_url"`
		} `json:"item"`
	} `json:"article_detail"`
	FailIdx []int `json:"fail_idx"`
}

type PublishArticleItem struct {
	NewsItem
	IsDeleted bool `json:"is_deleted"`
}

type GetPublishedArticle struct {
	NewsItem []PublishArticleItem `json:"news_item"`
}

type GetPublishedList struct {
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	Item       []struct {
		ArticleId string `json:"article_id"`
		Content   struct {
			NewsItem []PublishArticleItem `json:"news_item"`
		} `json:"content"`
		UpdateTime int `json:"update_time"`
	} `json:"item"`
}

// 文章创建参数校验
func validateArticles(articleList *[]ArticleItem) error {
	if len(*articleList) == 0 {
		return &core.ApiError{
			ErrCode: 500,
			ErrMsg:  "不能创建空草稿箱内容",
		}
	}
	for key, item := range *articleList {
		if item.Title == "" || item.Content == "" || item.ThumbMediaId == "" {
			return &core.ApiError{
				ErrCode: 500,
				ErrMsg:  fmt.Sprintf("第%d篇文章缺少必要字段，文章标题，内容，封面图都不能为空", key),
			}
		}
	}
	return nil
}
