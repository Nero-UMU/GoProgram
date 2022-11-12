package errcode

import (
	"fmt"
	"net/http"
)

// 定义 Error 结构体
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

// 用以存储所有的状态信息
var codes = map[int]string{}

// 新建状态信息
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在,请更换一个", code))
	}
	// 如果状态信息是新的,则添加到codes中
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// 返回code值的方法
func (e *Error) Code() int {
	return e.code
}

// 返回msg值的方法
func (e *Error) Msg() string {
	return e.msg
}

// 重写 Error 结构体的 Error 方法,其出错后会返回如下信息
func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d,错误信息: %s", e.Code(), e.Msg())
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// 返回状态详细信息
func (e *Error) Details() []string {
	return e.details
}

// 设置状态详细信息
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

// 获取状态码
func (e *Error) StatusCode() int {
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequest.Code():
		return http.StatusTooManyRequests
	case NotFound.Code():
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
