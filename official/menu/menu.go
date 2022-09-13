package menu

import (
	"go-wechat/core"
)

type Menu struct {
	Request *core.ApiRequest
}

func NewMenu(token core.AccessToken) *Menu {
	return &Menu{Request: core.NewApiRequest(token)}
}

func (m *Menu) CurrentSelfMenuInfo() (*ButtonRes, error) {
	var btn ButtonRes
	_, err := m.Request.JsonGet("/cgi-bin/get_current_selfmenu_info", nil, &btn)
	if err != nil {
		return nil, err
	}

	return &btn, nil
}

// Create 创建菜单
func (m *Menu) Create(buttons []CreateButtonItem) error {
	body := CreateMenuParams{Button: buttons}
	_, err := m.Request.JsonPost("/cgi-bin/menu/create", nil, body, nil)
	return err
}

// CreateConditionalMenu 创建个性化菜单
func (m *Menu) CreateConditionalMenu(buttons []CreateButtonItem, mathRule ButtonMatchRule) error {
	body := CreateConditionalParams{Button: buttons, Matchrule: mathRule}
	_, err := m.Request.JsonPost("/cgi-bin/menu/addconditional", nil, body, nil)
	return err
}

// DeleteConditionalMenu 删除个性化菜单
func (m *Menu) DeleteConditionalMenu(menuId string) error {
	_, err := m.Request.JsonPost("/cgi-bin/menu/addconditional", nil, map[string]string{
		"menuid": menuId,
	}, nil)
	return err
}

// TryMatch 匹配个性化菜单
func (m *Menu) TryMatch(userId string) (*CreateConditionalParams, error) {
	var menu CreateConditionalParams
	_, err := m.Request.JsonPost("/cgi-bin/menu/trymatch", nil, map[string]string{
		"user_id": userId,
	}, &menu)
	if err != nil {
		return nil, err
	}

	return &menu, nil
}

// DeleteMenu 删除菜单
func (m *Menu) DeleteMenu() error {
	_, err := m.Request.JsonGet("/cgi-bin/menu/delete", nil, nil)
	return err
}

// GetMenu 获取所有菜单
func (m *Menu) GetMenu() (*GetMenuRes, error) {
	var res GetMenuRes
	_, err := m.Request.JsonGet("/cgi-bin/menu/get", nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
