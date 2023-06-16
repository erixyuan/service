package apibase

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/dto"
	"service/enum/err_enum"
	"service/global"
	"service/interfaces"
	"service/utils"
)

func Response(ctx *gin.Context, errorObj *err_enum.ErrorObj, resp interfaces.BaseResponseInterface) {
	if resp == nil {
		resp = &dto.BaseRespParam{
			Code: 20000,
			Msg:  "success",
		}
	} else {
		if resp.GetCode() == 0 {
			resp.SetCode(20000)
		}
		if resp.GetMsg() == "" {
			resp.SetMsg("success")
		}
	}

	// 手动填充
	resp.SetRequestId(utils.GetRequestId(ctx))
	resp.SetTraceId(utils.GetTraceId(ctx))
	if errorObj != nil {
		resp.SetCode(errorObj.Code)
		resp.SetMsg(errorObj.Msg)
	}
	// 返回
	ctx.JSON(http.StatusOK, resp)
}

func Success(ctx *gin.Context, resp interfaces.BaseResponseInterface) {
	Response(ctx, nil, resp)
	return
}

func Fail(ctx *gin.Context, errorObj *err_enum.ErrorObj) {
	Response(ctx, errorObj, nil)
}

func GetHandler[T interfaces.BaseRequestInterface, G interfaces.BaseResponseInterface](handler func(T) (G, *err_enum.ErrorObj)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if p := err.(error); p != nil {
					global.GetLogger().Error("error : " + p.Error()) // this is error
					if err_my, ok := p.(*err_enum.ErrorObj); ok {
						Fail(c, &err_enum.ErrorObj{
							Code: err_my.Code,
							Msg:  err_my.Error(),
						})
					} else {
						Fail(c, &err_enum.ErrorObj{
							Code: 501,
							Msg:  p.Error(),
						})

					}

				} else {
					global.GetLogger().Info("not a normal error, continue panic") // this is error
					panic(p)
				}
			}
		}()

		var reqParam T
		ctx := context.Background()
		if err := c.ShouldBindJSON(&reqParam); err != nil {
			Fail(c, err_enum.ParamErr)
			return
		}

		// 设置trace id 和 request id
		// 从头部获取trace id 和 request id
		traceId := utils.GetTraceId(c)
		requestId := utils.GetRequestId(c)
		token := utils.GetToken(c)

		// 如果trace id 没有，用request id 设置 trace id
		if reqParam.GetTraceId() == "" {
			if traceId != "" {
				reqParam.SetTraceId(traceId)
				context.WithValue(ctx, "trace_id", traceId)
			} else {
				if reqParam.GetRequestId() != "" {
					reqParam.SetTraceId(reqParam.GetRequestId())
					context.WithValue(ctx, "trace_id", reqParam.GetRequestId())
				} else {
					reqParam.SetTraceId(requestId)
					context.WithValue(ctx, "trace_id", requestId)
				}
			}
		}
		if traceId == "" {
			utils.SetTraceId(c, reqParam.GetTraceId())
		}
		if token != "" {
			global.GetLogger().InfoByTrace(reqParam, "获取到token %+v", token)
			if ok := c.Value("uid"); ok != nil {
				reqParam.SetUid(ok.(int))
			}
		}

		reqParam.SetRequestId(requestId)
		if respParam, err := handler(reqParam); err != nil {
			Fail(c, err)
		} else {
			//请求成功 设置response 的公共参数
			Success(c, respParam)
		}
	}
}
