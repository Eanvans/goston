package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000, "系统繁忙, 请稍后重试")
	InvalidParams             = NewError(10001, "无效输入")
	NotFound                  = NewError(10002, "未找到数据")
	UnauthorizedAuthNotExist  = NewError(10003, "账户不存在")
	UnauthorizedAuthFailed    = NewError(10004, "账户密码错误")
	UnauthorizedTokenError    = NewError(10005, "未登录")
	UnauthorizedTokenTimeout  = NewError(10006, "登录信息已过期")
	UnauthorizedTokenGenerate = NewError(10007, "Token 生成失败")
	TooManyRequests           = NewError(10008, "请求过多")
	InvalidLink               = NewError(10009, "链接失效")
	OutDateWeiXinAccessToken  = NewError(10010, "微信token超时")
	FunctionInService         = NewError(10011, "功能维护中,请先试用替代功能")

	GatewayMethodsLimit    = NewError(10109, "网关仅接受GET或POST请求")
	GatewayLostSign        = NewError(10110, "网关请求缺少签名")
	GatewayLostAppKey      = NewError(10111, "网关请求缺少APP KEY")
	GatewayAppKeyInvalid   = NewError(10112, "网关请求无效APP KEY")
	GatewayAppKeyClosed    = NewError(10113, "网关请求APP KEY已停用")
	GatewayParamSignError  = NewError(10114, "网关请求参数签名错误")
	GatewayTooManyRequests = NewError(10115, "网关请求频次超限")

	FileUploadFailed = NewError(10200, "文件上传失败")
	FileInvalidExt   = NewError(10201, "文件类型不合法")
	FileInvalidSize  = NewError(10202, "文件大小超限")
)

type Error struct {
	code    int
	msg     string
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedAuthFailed.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
