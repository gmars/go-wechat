package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

type AesCbcService struct {
	AesKey string
}

func NewAesCbc(key string) *AesCbcService {
	return &AesCbcService{AesKey: key}
}

// Pkcs7Padding pkcs7模式填充
func (a *AesCbcService) Pkcs7Padding(msg *string) {
	padding := 32 - len(*msg)%32
	*msg += strings.Repeat(string(rune(padding)), padding)
}

// Encryption aes-cbc-256加密解密
func (a *AesCbcService) Encryption(msg *string, encrypt bool) ([]byte, error) {
	var (
		resBytes    []byte
		aesInstance cipher.BlockMode
		err         error
		msgBytes    []byte
	)
	//处理key
	decodeKey, err := base64.StdEncoding.DecodeString(a.AesKey + "=")
	if err != nil {
		return nil, err
	}

	cp, err := aes.NewCipher(decodeKey)
	if err != nil {
		return nil, err
	}

	if encrypt {
		aesInstance = cipher.NewCBCEncrypter(cp, decodeKey[:cp.BlockSize()])
		a.Pkcs7Padding(msg)
		msgBytes = StringToBytes(*msg)
	} else {
		aesInstance = cipher.NewCBCDecrypter(cp, decodeKey[:cp.BlockSize()])
		if msgBytes, err = base64.StdEncoding.DecodeString(*msg); err != nil {
			return nil, err
		}
	}
	//初始化结果
	resBytes = make([]byte, len(msgBytes))
	aesInstance.CryptBlocks(resBytes, msgBytes)
	if encrypt {
		encode := base64.StdEncoding
		buf := make([]byte, encode.EncodedLen(len(resBytes)))
		encode.Encode(buf, resBytes)
		return buf, nil
	} else {
		msgLength := BytesToInt32(resBytes[16:20])
		return resBytes[20 : msgLength+20], nil
	}
}
