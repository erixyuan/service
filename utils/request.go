package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"service/static"
)

// 获取头部中的跟踪id
func GetTraceId(ctx *gin.Context) string {
	return ctx.Request.Header.Get(static.X_TRACE_ID)
}

// 获取当前请求的request id
func GetRequestId(ctx *gin.Context) string {
	return ctx.Request.Header.Get(static.X_REQUEST_ID)
}

func GetToken(ctx *gin.Context) string {
	return ctx.Request.Header.Get(static.TOKEN)
}

// 设置并返回request id
func SetRequestId(ctx *gin.Context) string {
	reqId := uuid.New().String()
	// 设置request 头部
	ctx.Request.Header.Set(static.X_REQUEST_ID, reqId)
	// 设置response 头部
	ctx.Header(static.X_REQUEST_ID, reqId)
	return reqId
}

func SetTraceId(ctx *gin.Context, traceId string) {
	ctx.Request.Header.Set(static.X_TRACE_ID, traceId)
}
