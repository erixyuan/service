package middlebase

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"service/dto"
	"service/global"
	"service/utils"
	"strings"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinLoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 设置request id
		requestId := utils.SetRequestId(c)
		ctx := context.Background()
		context.WithValue(ctx, "request_id", requestId)

		startTime := time.Now()
		// 获取全局跟踪id，用于全链路追踪
		traceId := utils.GetTraceId(c)

		// 获取当前请求的id
		reqId := utils.GetRequestId(c)

		baseParam := dto.BaseParam{
			RequestId: reqId,
			TraceId:   traceId,
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		reqBodyContent := ""

		if body, err := ioutil.ReadAll(c.Request.Body); err != nil {
			reqBodyContent = err.Error()
		} else {
			//判断如果是文件上传，就不打印请求体内容
			contentType := c.Request.Header.Get("Content-Type")
			if !strings.Contains(strings.ToLower(contentType), "multipart") {
				reqBodyContent = string(body)
			}
			// 关闭流
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		}

		method := c.Request.Method
		reqUrl := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// 打印入口请求内容
		global.GetLogger().InfoRequest("Gin Http Request: ", traceId, reqId, method, clientIP, reqUrl, reqBodyContent, statusCode)

		c.Next()

		// 重新全局跟踪id，用于全链路追踪
		traceId = utils.GetTraceId(c)

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime).Milliseconds()

		// 如果请求时间大于3秒，打个日志
		if latencyTime > 2000 {
			global.GetLogger().InfoByTrace(&baseParam, "请求时间过长请求：url [%s], latency [%d]", reqUrl, latencyTime)
		}

		respBodyContent := blw.body.String()

		// 打印出口返回内容
		global.GetLogger().InfoResponse("Gin Http Response: ", traceId, reqId, method, clientIP, reqUrl, reqBodyContent, respBodyContent, statusCode, latencyTime)
	}
}
