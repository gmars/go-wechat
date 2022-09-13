package core

import "errors"

const (
	msgMustConfigParams = "缺少必要配置参数"
)

func errorHandle(msg string) error {
	return errors.New(msg)
}

var IsLackMustConfigParams = errorHandle(msgMustConfigParams)

type ApiError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *ApiError) Error() string {
	return e.ErrMsg
}

func NewError(code int, msg string) error {
	return &ApiError{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func ErrorCodes(err error) int {
	if err == nil {
		return 0
	}

	switch typed := err.(type) {
	case *ApiError:
		return typed.ErrCode
	default:
		return -500
	}
}
