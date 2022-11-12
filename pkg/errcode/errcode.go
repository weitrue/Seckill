package errcode

import (
	"net/http"

	"google.golang.org/grpc/status"
)

const (
	// CodeOK 请求成功业务状态码
	CodeOK = 200
	// MsgOK 请求成功消息
	MsgOK = "OK"

	// CodeCustom 自定义错误业务状态码
	CodeCustom = 7000
	// MsgCustom 自定义错误消息
	MsgCustom = "自定义错误"
)

// Err 业务错误结构
type Err struct {
	code     uint32
	httpCode int
	msg      string
}

// Code 业务状态码
func (e *Err) Code() uint32 {
	return e.code
}

// HTTPCode HTTP状态码
func (e *Err) HTTPCode() int {
	return e.httpCode
}

// Error 消息
func (e *Err) Error() string {
	return e.msg
}

// 业务错误
var (
	NoErr = NewErr(CodeOK, MsgOK)

	ErrCustom     = NewErr(CodeCustom, MsgCustom)
	ErrUnexpected = NewErr(7777, "服务器繁忙，请稍后重试")

	ErrInvalidOpt       = NewErr(9998, "请求操作非法")
	ErrTokenNotValidYet = NewErr(9999, "token暂未生效", http.StatusUnauthorized)
	ErrInvalidUrl       = NewErr(10000, "请求URL非法")
	ErrInvalidHeader    = NewErr(10001, "请求头部错误")
	ErrInvalidParams    = NewErr(10002, "请求参数错误")
	ErrTokenVerify      = NewErr(10003, "token验证出错", http.StatusUnauthorized)
	ErrTokenExpire      = NewErr(10004, "token过期", http.StatusUnauthorized)
	ErrOptAuth          = NewErr(10005, "无操作权限")
	ErrReadAuth         = NewErr(10006, "无查看权限")
	ErrNotFind          = NewErr(10007, "无对应信息")
	ErrApiUse           = NewErr(10008, "接口未开放")
	ErrSignature        = NewErr(10009, "请求签名校验错误")
)

var codeToErr = map[uint32]*Err{
	200: NoErr,

	7000: ErrCustom,
	7777: ErrUnexpected,

	9998:  ErrInvalidOpt,
	9999:  ErrTokenNotValidYet,
	10000: ErrInvalidUrl,
	10001: ErrInvalidHeader,
	10002: ErrInvalidParams,
	10003: ErrTokenVerify,
	10004: ErrTokenExpire,
	10005: ErrOptAuth,
	10006: ErrReadAuth,
	10007: ErrNotFind,
	10008: ErrApiUse,
	10009: ErrSignature,
}

// NewErr 创建新的业务错误
func NewErr(code uint32, msg string, httpCode ...int) *Err {
	hc := http.StatusOK
	if len(httpCode) != 0 {
		hc = httpCode[0]
	}

	return &Err{code: code, httpCode: hc, msg: msg}
}

// NewCustomErr 创建新的自定义错误
func NewCustomErr(msg string, httpCode ...int) *Err {
	return NewErr(CodeCustom, msg, httpCode...)
}

// IsErr 判断是否为业务错误
func IsErr(err error) bool {
	if err == nil {
		return true
	}

	_, ok := err.(*Err)
	return ok
}

// ParseErr 解析业务错误
func ParseErr(err error) *Err {
	if err == nil {
		return NoErr
	}

	if e, ok := err.(*Err); ok {
		return e
	}

	s, _ := status.FromError(err)
	c := uint32(s.Code())
	if c == CodeCustom {
		return NewCustomErr(s.Message())
	}

	return ParseCode(c)
}

// ParseCode 解析业务状态码对应的业务错误
func ParseCode(code uint32) *Err {
	if e, ok := codeToErr[code]; ok {
		return e
	}

	return ErrUnexpected
}
