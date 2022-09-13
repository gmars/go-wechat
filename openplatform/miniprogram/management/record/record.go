package record

import "go-wechat/core"

type Record struct {
	request *core.ApiRequest
}

func NewRecordManagement(authorizerAccessToken core.AccessToken) *Record {
	return &Record{request: core.NewApiRequest(authorizerAccessToken)}
}

// GetIllegalRecords 获取小程序违规处罚记录
func (s *Record) GetIllegalRecords(startTime, endTime int) (*[]RecordsItem, error) {
	var res GetIllegalRecordsRes
	_, err := s.request.JsonPost("/wxa/getillegalrecords", nil, map[string]int{
		"start_time": startTime,
		"end_time":   endTime,
	}, &res)
	return &res.Records, err
}

// GetAppealRecords 获取小程序申诉记录
func (s *Record) GetAppealRecords(illegalRecordId string) (*[]GetAppealRecordsItem, error) {
	var res GetAppealRecordsRes
	_, err := s.request.JsonPost("/wxa/getappealrecords", nil, map[string]string{
		"illegal_record_id": illegalRecordId,
	}, &res)
	return &res.Records, err
}
