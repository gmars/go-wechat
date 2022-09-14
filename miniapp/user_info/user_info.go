package user_info

import (
	"github.com/gmars/go-wechat/core"
)

type UserInfo struct {
	request *core.ApiRequest
}

func NewUserInfo(accessToken core.AccessToken) *UserInfo {
	return &UserInfo{request: core.NewApiRequest(accessToken)}
}

// GetPluginOpenPId 获取插件用户openpid
func (s *UserInfo) GetPluginOpenPId(code string) (string, error) {
	var res OpenIdRes
	_, err := s.request.JsonPost("/wxa/getpluginopenpid", nil, map[string]string{
		"code": code,
	}, &res)
	return res.OpenPid, err
}

// CheckEncryptedData 检查加密信息
func (s *UserInfo) CheckEncryptedData(encryptedMsgHash string) (*CheckEncryptedMsg, error) {
	var res CheckEncryptedMsg
	_, err := s.request.JsonPost("/wxa/business/checkencryptedmsg", nil, map[string]string{
		"encrypted_msg_hash": encryptedMsgHash,
	}, &res)
	return &res, err
}

// GetPaidUnionId 支付后获取 UnionId
func (s *UserInfo) GetPaidUnionId(openId string, paymentInfo *PaymentAddition) (string, error) {
	var (
		res    PaymentUnionRes
		params = map[string]string{
			"openid": openId,
		}
	)

	if paymentInfo != nil {
		params["transaction_id"] = paymentInfo.TransactionId
		params["mch_id"] = paymentInfo.MchId
		params["out_trade_no"] = paymentInfo.OutTradeNo
	}

	_, err := s.request.JsonGet("/wxa/getpaidunionid", params, &res)
	return res.UnionId, err
}

// GetUserEncryptKey 获取用户encryptKey
func (s *UserInfo) GetUserEncryptKey(openid, signature, sigMethod string) (*[]UserEncryptKeyInfoItem, error) {
	var res UserEncryptKeyRes
	if sigMethod == "" {
		sigMethod = "hmac_sha256"
	}
	_, err := s.request.JsonPost("/wxa/business/getuserencryptkey", nil, map[string]string{
		"openid":     openid,
		"signature":  signature,
		"sig_method": sigMethod,
	}, &res)
	return &res.KeyInfoList, err
}

// GetPhoneNumber 获取手机号
func (s *UserInfo) GetPhoneNumber(code string) (*PhoneInfo, error) {
	var res PhoneInfoRes
	_, err := s.request.JsonPost("/wxa/business/getuserphonenumber", nil, map[string]string{
		"code": code,
	}, &res)
	return &res.PhoneInfo, err
}
