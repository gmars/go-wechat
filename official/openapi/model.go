package openapi

type QuotaData struct {
	DailyLimit int64 `json:"daily_limit"`
	Used       int64 `json:"used"`
	Remain     int64 `json:"remain"`
}

type QuotaRes struct {
	Quota QuotaData `json:"quota"`
}

type RidRequest struct {
	InvokeTime   int64  `json:"invoke_time"`
	CostInMs     int64  `json:"cost_in_ms"`
	RequestUrl   string `json:"request_url"`
	RequestBody  string `json:"request_body"`
	ResponseBody string `json:"response_body"`
	ClientIp     string `json:"client_ip"`
}

type RidRequestRes struct {
	Request RidRequest `json:"request"`
}
