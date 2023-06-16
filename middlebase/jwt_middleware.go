package middlebase

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"service/apibase"
	"service/config"
	"service/dto"
	"service/enum/err_enum"
	"service/global"
	"service/static"
	"service/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		baseParam := dto.BaseParam{
			RequestId: utils.GetRequestId(c),
			TraceId:   utils.GetTraceId(c),
		}
		token := c.GetHeader("token")
		if token == "" {
			excuseList := config.ServerGlobalConfig.JwtConfig.JwtExcuse
			for _, path := range excuseList {
				compile := regexp.MustCompile(path)
				if compile.MatchString(c.Request.RequestURI) {
					global.GetLogger().InfoByTrace(&baseParam, "JwtMiddleware path excuse :"+c.Request.RequestURI)
					c.Next()
					return
				}
			}
			global.GetLogger().ErrorByTrace(&baseParam, "JwtMiddleware authentication token is empty")
			apibase.Fail(c, err_enum.TokenInvalid)
			c.Abort()
			return
		} else {
			// 解析token
			claims, err := utils.ParseToken(token)
			if err != nil {
				global.GetLogger().ErrorByTrace(&baseParam, "JwtMiddleware authentication token is invalid :"+err.Error())
				apibase.Fail(c, err_enum.TokenInvalid)
				c.Abort()
				return
			} else if claims.Id < 1 {
				global.GetLogger().ErrorByTrace(&baseParam, "JwtMiddleware authentication user_id is error : %d", claims.Id)
				apibase.Fail(c, err_enum.TokenInvalid)
				c.Abort()
				return
			}
			global.GetLogger().InfoByTrace(&baseParam, "JwtMiddleware claims : %+v", claims)
			c.Set(static.USER_ID_KEY, claims.Id)
		}
		c.Next()
	}
}
