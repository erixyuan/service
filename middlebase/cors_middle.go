package middlebase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length, Authorization, Accept,X-Requested-With")
		//ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET,OPTIONS")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
		} else {
			ctx.Next()
		}
	}
}
