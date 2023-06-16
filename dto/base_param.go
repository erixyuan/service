package dto

import (
	"service/enum/err_enum"
)

type BaseParam struct {
	RequestId string `json:"request_id"`
	TraceId   string `json:"trace_id"`
}

func (b BaseParam) GetTraceId() string {
	return b.TraceId
}

func (b BaseParam) GetRequestId() string {
	return b.RequestId
}

func (b *BaseParam) SetTraceId(traceId string) {
	b.TraceId = traceId
}

func (b *BaseParam) SetRequestId(requestId string) {
	b.RequestId = requestId
}

// BaseReqParam
type BaseReqParam struct {
	BaseParam
	Uid int `json:"uid"`
}

func (b *BaseReqParam) SetUid(uid int) {
	b.Uid = uid
}

func (b BaseReqParam) GetUid() int {
	return b.Uid
}

type BaseRespParam struct {
	BaseParam
	Code err_enum.ErrorCode `json:"code"`
	Msg  string             `json:"msg"`
}

func (r *BaseRespParam) SetCode(code err_enum.ErrorCode) {
	r.Code = code
}

func (r *BaseRespParam) SetMsg(msg string) {
	r.Msg = msg
}

func (r *BaseRespParam) GetCode() err_enum.ErrorCode {
	return r.Code
}

func (r *BaseRespParam) GetMsg() string {
	return r.Msg
}
