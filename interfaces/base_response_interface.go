package interfaces

import (
	"service/enum/err_enum"
)

type BaseInterface interface {
	GetTraceId() string
	GetRequestId() string
	SetTraceId(traceId string)
	SetRequestId(traceId string)
}

type BaseResponseInterface interface {
	BaseInterface
	SetCode(code err_enum.ErrorCode)
	SetMsg(msg string)
	GetCode() err_enum.ErrorCode
	GetMsg() string
}

type BaseRequestInterface interface {
	BaseInterface
	SetUid(uid int)
	GetUid() int
}
