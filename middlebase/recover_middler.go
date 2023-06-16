package middlebase

import (
	"github.com/gin-gonic/gin"
	"runtime"
	"service/apibase"
	"service/dto"
	"service/enum/err_enum"
	"service/global"
	"service/utils"
)

func RecoverMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				baseParam := dto.BaseParam{
					RequestId: utils.GetRequestId(c),
					TraceId:   utils.GetTraceId(c),
				}
				global.GetLogger().ErrorByTrace(&baseParam, "panic异常: %v，调用处: %s", err, string(buf))
				apibase.Fail(c, err_enum.SystemErr)
				c.Abort()
			}
		}()
		c.Next()
	}
}
