package dev

import (
	"context"
	"github.com/gmars/go-wechat/core"
	"net/url"
	"strconv"
)

type Dev struct {
	request     *core.ApiRequest
	accessToken core.AccessToken
}

func NewDev(accessToken core.AccessToken) *Dev {
	return &Dev{request: core.NewApiRequest(accessToken), accessToken: accessToken}
}

// GetFeedbackMedia 获取 mediaId 图片
func (s *Dev) GetFeedbackMedia(recordId int, mediaId string) (string, error) {
	token, err := s.accessToken.GetAccessToken(context.Background())
	if err != nil {
		return "", err
	}
	imageUrl, err := url.Parse("https://api.weixin.qq.com/cgi-bin/media/getfeedbackmedia")
	if err != nil {
		return "", err
	}

	query := imageUrl.Query()
	query.Add("access_token", token)
	query.Add("record_id", strconv.Itoa(recordId))
	query.Add("media_id", mediaId)
	imageUrl.RawQuery = query.Encode()
	return imageUrl.String(), nil
}

// GetDomainInfo 查询域名配置
func (s *Dev) GetDomainInfo(action string) (*GetDomainInfoRes, error) {
	var res GetDomainInfoRes
	_, err := s.request.JsonPost("/wxa/getwxadevinfo", nil, map[string]string{
		"action": action,
	}, &res)
	return &res, err
}

// GetPerformance 获取性能数据
func (s *Dev) GetPerformance(params *GetPerformanceParams) (*GetPerformanceRes, error) {
	var res GetPerformanceRes
	_, err := s.request.JsonPost("/wxaapi/log/get_performance", nil, params, &res)
	return &res, err
}

// GetSceneList 获取访问来源
func (s *Dev) GetSceneList() (*GetSceneListRes, error) {
	var res GetSceneListRes
	_, err := s.request.JsonGet("/wxaapi/log/get_scene", nil, &res)
	return &res, err
}

// GetVersionList 获取客户端版本
func (s *Dev) GetVersionList() (*GetVersionListRes, error) {
	var res GetVersionListRes
	_, err := s.request.JsonGet("/wxaapi/log/get_client_version", nil, &res)
	return &res, err
}

// RealtimeLogSearch 查询实时日志
func (s *Dev) RealtimeLogSearch(params *RealtimeLogSearchParams) (*RealtimeLogSearchRes, error) {
	var (
		res   RealtimeLogSearchRes
		query = map[string]string{
			"date":      params.Date,
			"begintime": strconv.Itoa(params.BeginTime),
			"endtime":   strconv.Itoa(params.EndTime),
			"start":     strconv.Itoa(params.Start),
			"limit":     strconv.Itoa(params.Limit),
			"traceId":   params.TraceId,
			"url":       params.Url,
			"id":        params.Id,
			"filterMsg": params.FilterMsg,
			"level":     strconv.Itoa(params.Level),
		}
	)

	_, err := s.request.JsonGet("/wxaapi/userlog/userlog_search", query, &res)
	return &res, err
}

// GetFeedback 获取用户反馈列表
func (s *Dev) GetFeedback(params *GetFeedbackParams) (*GetFeedbackRes, error) {
	var (
		res   GetFeedbackRes
		query = map[string]string{
			"type": strconv.Itoa(params.Type),
			"page": strconv.Itoa(params.Page),
			"num":  strconv.Itoa(params.Num),
		}
	)

	_, err := s.request.JsonGet("/wxaapi/feedback/list", query, &res)
	return &res, err
}

// GetJsErrDetail 查询 js 错误详情
func (s *Dev) GetJsErrDetail(params *GetJsErrDetailParams) (*GetJsErrDetailRes, error) {
	var (
		res GetJsErrDetailRes
	)

	_, err := s.request.JsonPost("/wxaapi/log/jserr_detail", nil, params, &res)
	return &res, err
}

// GetJsErrList 查询错误列表
func (s *Dev) GetJsErrList(params *GetJsErrListParams) (*GetJsErrListRes, error) {
	var (
		res GetJsErrListRes
	)
	_, err := s.request.JsonPost("/wxaapi/log/jserr_list", nil, params, &res)
	return &res, err
}

// GetGrayReleasePlan 获取分阶段发布详情
func (s *Dev) GetGrayReleasePlan() (*GetGrayReleasePlanRes, error) {
	var (
		res GetGrayReleasePlanRes
	)
	_, err := s.request.JsonGet("/wxa/getgrayreleaseplan", nil, &res)
	return &res, err
}

// AddNearbyPoi 添加地点
func (s *Dev) AddNearbyPoi(params *AddNearbyPoiParams) (*AddNearbyPoiRes, error) {
	var (
		res AddNearbyPoiRes
	)
	_, err := s.request.JsonPost("/wxa/addnearbypoi", nil, params, &res)
	return &res, err
}

// DeleteNearbyPoi 删除地点
func (s *Dev) DeleteNearbyPoi(poiId string) error {
	var (
		res AddNearbyPoiRes
	)
	_, err := s.request.JsonPost("/wxa/delnearbypoi", nil, map[string]string{
		"poi_id": poiId,
	}, &res)
	return err
}

// GetNearbyPoiList 查看地点
func (s *Dev) GetNearbyPoiList(page, pageSize int) (*PositionList, error) {
	var (
		res PositionListRes
	)
	_, err := s.request.JsonPost("/wxa/getnearbypoilist", nil, map[string]int{
		"page":      page,
		"page_rows": pageSize,
	}, &res)
	return &res.Data, err
}

// SetShowStatus 设置展示状态
func (s *Dev) SetShowStatus(poiId string, status int) error {
	_, err := s.request.JsonPost("/wxa/getnearbypoilist", nil, map[string]interface{}{
		"poi_id": poiId,
		"status": status,
	}, nil)
	return err
}
