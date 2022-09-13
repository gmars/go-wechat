package category

import "go-wechat/core"

type Category struct {
	request *core.ApiRequest
}

func NewCategory(authorizerAccessToken core.AccessToken) *Category {
	return &Category{request: core.NewApiRequest(authorizerAccessToken)}
}

// GetAllCategories 获取可设置的所有类目
func (s *Category) GetAllCategories() (*[]CateItem, error) {
	var res GetAllCategoriesRes
	_, err := s.request.JsonGet("/cgi-bin/wxopen/getallcategories", nil, &res)
	return &res.CategoriesList.Categories, err
}

// GetSettingCategories 获取已设置的所有类目
func (s *Category) GetSettingCategories() (*GetSettingCategoriesRes, error) {
	var res GetSettingCategoriesRes
	_, err := s.request.JsonGet("/cgi-bin/wxopen/getcategory", nil, &res)
	return &res, err
}

// GetAllCategoriesByType 获取不同类型主体可设置的类目
func (s *Category) GetAllCategoriesByType() (*[]CateItem, error) {
	var res GetAllCategoriesRes
	_, err := s.request.JsonPost("/cgi-bin/wxopen/getcategoriesbytype", nil, nil, &res)
	return &res.CategoriesList.Categories, err
}

// AddCategory 添加类目
func (s *Category) AddCategory(params *AddCategoryParams) error {
	_, err := s.request.JsonPost("/cgi-bin/wxopen/addcategory", nil, params, nil)
	return err
}

// DeleteCategory 删除类目
func (s *Category) DeleteCategory(first, second int) error {
	_, err := s.request.JsonPost("/cgi-bin/wxopen/deletecategory", nil, map[string]int{
		"first":  first,
		"second": second,
	}, nil)
	return err
}

// ModifyCategory 修改类目资质信息
func (s *Category) ModifyCategory(params *ModifyCategoryParams) error {
	_, err := s.request.JsonPost("/cgi-bin/wxopen/modifycategory", nil, params, nil)
	return err
}

// GetAllCategoryName 获取类目名称信息
func (s *Category) GetAllCategoryName() (*GetAllCategoryNameRes, error) {
	var res GetAllCategoryNameRes
	_, err := s.request.JsonPost("/cgi-bin/wxopen/get_category", nil, nil, &res)
	return &res, err
}
