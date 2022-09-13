package member

type BindTesterRes struct {
	UserStr string `json:"userstr"`
}
type Members struct {
	UserStr string `json:"userstr"`
}

type GetTesterRes struct {
	Members []Members `json:"members"`
}
