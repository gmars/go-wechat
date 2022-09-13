package message

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"go-wechat/core"
	"go-wechat/util"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	AesKey  string
	Token   string
	AppId   string
	AesTool *util.AesCbcService
}

func NewMessageHandler(aesKey, token, AppId string) *Message {
	return &Message{
		AesKey:  aesKey,
		Token:   token,
		AppId:   AppId,
		AesTool: util.NewAesCbc(aesKey),
	}
}

// Handler 消息处理器
func (m *Message) Handler(req *http.Request) (*ReplyMsg, error) {
	var (
		baseType TypeParse
		err      error
	)

	//校验解密消息体
	msgBytes, err := m.decryptMsg(req)
	if err != nil {
		return nil, err
	}

	//解析消息体
	if err = xml.Unmarshal(msgBytes, &baseType); err != nil {
		return nil, err
	}

	if baseType.MsgType == MsgEvent {
		return m.parseEvent(&baseType, msgBytes)
	} else if baseType.MsgType != MsgEvent && baseType.MsgType != "" {
		return m.parseMsg(&baseType, msgBytes)
	} else if baseType.InfoType != "" && baseType.MsgType == "" {
		return m.parseInfo(&baseType, msgBytes)
	}

	return nil, core.NewError(400, "服务器未处理您传入的消息")
}

// EncryptMessage 生成回复加密消息
func (m *Message) EncryptMessage(msg, aesKey, token string) (string, error) {
	nonce, err := util.NonceStr(16)
	if err != nil {
		return "", err
	}

	//aes加密
	respMsg := nonce + string(util.Int32ToBytes(int32(len(msg)))) + msg + m.AppId
	encryptor := util.AesCbcService{AesKey: aesKey}
	enMsg, err := encryptor.Encryption(&respMsg, true)
	if err != nil {
		return "", err
	}

	//签名
	nonceStr, err := util.NonceStr(10)
	if err != nil {
		return "", err
	}
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	enMsgStr := util.BytesToString(enMsg)
	signStr := m.generateSign(token, nonceStr, timestamp, enMsgStr)

	return fmt.Sprintf("<xml>\n<Encrypt><![CDATA[%s]]></Encrypt>\n<MsgSignature><![CDATA[%s]]></MsgSignature>\n<TimeStamp>%s</TimeStamp>\n<Nonce><![CDATA[%s]]></Nonce>\n</xml>",
		string(enMsg), signStr, timestamp, nonceStr), nil
}

// CheckServiceSignature 生成回复加密消息
func (m *Message) CheckServiceSignature(req *http.Request) string {
	nonce := req.FormValue("nonce")
	timestamp := req.FormValue("timestamp")
	sign := req.FormValue("signature")
	echoStr := req.FormValue("echostr")
	tmpSign := m.generateSign(m.Token, timestamp, nonce)
	if tmpSign == sign {
		return echoStr
	} else {
		return "Check Wechat Request Fail"
	}
}

// 解析加密消息体
func (m *Message) decryptMsg(req *http.Request) ([]byte, error) {
	var encMsg EncryptMsg
	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil || body == nil {
		return nil, err
	}

	if err = xml.Unmarshal(body, &encMsg); err != nil {
		return nil, err
	}

	if !m.checkMsgSign(req, encMsg.Encrypt, m.Token) {
		return nil, errors.New("消息验证不通过")
	}

	//解密
	msgBytes, err := m.AesTool.Encryption(&encMsg.Encrypt, false)
	if err != nil {
		return nil, err
	}
	return msgBytes, nil
}

// 消息签名
func (m *Message) generateSign(params ...string) string {
	sort.Strings(params)
	res := sha1.Sum([]byte(strings.Join(params, "")))
	return fmt.Sprintf("%x", res)
}

// 验证签名
func (m *Message) checkMsgSign(req *http.Request, msgEncrypt, token string) bool {
	nonce := req.FormValue("nonce")
	timestamp := req.FormValue("timestamp")
	msgSign := req.FormValue("msg_signature")

	comSign := m.generateSign(token, nonce, timestamp, msgEncrypt)
	return msgSign == comSign
}
