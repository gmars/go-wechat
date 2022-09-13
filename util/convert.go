package util

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"unsafe"
)

// BytesToInt32 bytes转int32
func BytesToInt32(data []byte) int32 {
	var decimal int32
	byteBuffer := bytes.NewBuffer(data)
	binary.Read(byteBuffer, binary.BigEndian, &decimal)
	return decimal
}

// Int32ToBytes 32位整数转byte
func Int32ToBytes(data int32) []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, int32(data))
	return byteBuffer.Bytes()
}

// NonceStr 生成随机字符串
func NonceStr(strLen int) (string, error) {
	seedPool := "abcdefghijklmnopqrstuvwxy1234567891234567890"
	randBytes := make([]byte, strLen)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(seedPool))
	for i, b := range randBytes {
		randBytes[i] = seedPool[b%symbolsByteLength]
	}
	return string(randBytes), nil
}

// StringToBytes 字符串转byte数组
func StringToBytes(str string) []byte {
	strX := (*[2]uintptr)(unsafe.Pointer(&str))
	strH := [3]uintptr{strX[0], strX[1], strX[1]}
	return *(*[]byte)(unsafe.Pointer(&strH))
}

// BytesToString byte数组转字符串
func BytesToString(bytesData []byte) string {
	return *(*string)(unsafe.Pointer(&bytesData))
}
