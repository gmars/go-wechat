package record

type RecordsItem struct {
	IllegalRecordId string `json:"illegal_record_id"`
	CreateTime      int    `json:"create_time"`
	IllegalReason   string `json:"illegal_reason"`
	IllegalContent  string `json:"illegal_content"`
	RuleUrl         string `json:"rule_url"`
	RuleName        string `json:"rule_name"`
}

type GetIllegalRecordsRes struct {
	Records []RecordsItem `json:"records"`
}

type GetAppealRecordsItem struct {
	AppealRecordId    int    `json:"appeal_record_id"`
	AppealTime        int    `json:"appeal_time"`
	AppealFrom        int    `json:"appeal_from"`
	AppealStatus      int    `json:"appeal_status"`
	PunishDescription string `json:"punish_description"`
	Materials         []struct {
		IllegalMaterial struct {
			Content    string `json:"content"`
			ContentUrl string `json:"content_url"`
		} `json:"illegal_material"`
		AppealMaterial struct {
			Reason           string   `json:"reason"`
			ProofMaterialIds []string `json:"proof_material_ids"`
		} `json:"appeal_material"`
	} `json:"materials"`
}

type GetAppealRecordsRes struct {
	Records []GetAppealRecordsItem `json:"records"`
}
