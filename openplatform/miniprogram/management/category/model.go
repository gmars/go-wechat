package category

type CateItem struct {
	Id       int   `json:"id"`
	Children []int `json:"children"`
	Qualify  struct {
		ExterList []struct {
			InnerList []struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"inner_list"`
		} `json:"exter_list"`
		Remark string `json:"remark"`
	} `json:"qualify"`
	Name          string `json:"name"`
	Level         int    `json:"level"`
	Father        int    `json:"father"`
	SensitiveType int    `json:"sensitive_type"`
}

type GetAllCategoriesRes struct {
	CategoriesList struct {
		Categories []CateItem `json:"categories"`
	} `json:"categories_list"`
}

type GetSettingCategoriesRes struct {
	Categories []struct {
		First       int    `json:"first"`
		FirstName   string `json:"first_name"`
		Second      int    `json:"second"`
		SecondName  string `json:"second_name"`
		AuditStatus int    `json:"audit_status"`
		AuditReason string `json:"audit_reason"`
	} `json:"categories"`
	Limit         int `json:"limit"`
	Quota         int `json:"quota"`
	CategoryLimit int `json:"category_limit"`
}

type AddCategoryParams struct {
	Categories []struct {
		First      int `json:"first"`
		Second     int `json:"second"`
		Certicates []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"certicates"`
	} `json:"categories"`
}

type ModifyCategoryParams struct {
	First      int `json:"first"`
	Second     int `json:"second"`
	Certicates []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"certicates"`
}

type GetAllCategoryNameRes struct {
	CategoryList []struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
		ThirdClass  string `json:"third_class"`
		FirstId     int    `json:"first_id"`
		SecondId    int    `json:"second_id"`
		ThirdId     int    `json:"third_id"`
	} `json:"category_list"`
}
