package management

import (
	"go-wechat/core"
	"go-wechat/openplatform/miniprogram/management/basic_info"
	"go-wechat/openplatform/miniprogram/management/category"
	"go-wechat/openplatform/miniprogram/management/domain"
	"go-wechat/openplatform/miniprogram/management/jumpqrcode"
	"go-wechat/openplatform/miniprogram/management/member"
	"go-wechat/openplatform/miniprogram/management/privacy"
	"go-wechat/openplatform/miniprogram/management/record"
	"go-wechat/openplatform/miniprogram/management/subscribe"
)

type Management struct {
	authorizerAccessToken core.AccessToken
}

func NewManagement(authorizerAccessToken core.AccessToken) *Management {
	return &Management{authorizerAccessToken: authorizerAccessToken}
}

func (s *Management) BasicInfo() *basic_info.BasicInfo {
	return basic_info.NewBasicInfoManager(s.authorizerAccessToken)
}

func (s *Management) Category() *category.Category {
	return category.NewCategory(s.authorizerAccessToken)
}

func (s *Management) Domain() *domain.Domain {
	return domain.NewDomain(s.authorizerAccessToken)
}

func (s *Management) JumpQrcode() *jumpqrcode.JumpQrcode {
	return jumpqrcode.NewJumpQrcode(s.authorizerAccessToken)
}

func (s *Management) Member() *member.Member {
	return member.NewMemberManagement(s.authorizerAccessToken)
}

func (s *Management) Privacy() *privacy.Privacy {
	return privacy.NewPrivacyManagement(s.authorizerAccessToken)
}

func (s *Management) Record() *record.Record {
	return record.NewRecordManagement(s.authorizerAccessToken)
}

func (s *Management) Subscribe() *subscribe.Subscribe {
	return subscribe.NewSubscribeManagement(s.authorizerAccessToken)
}
