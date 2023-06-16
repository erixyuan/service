package dto

import (
	"service/enum/err_enum"
)

type BaseOperationRespData struct {
	Result    int8   `json:"result"`
	ResultMsg string `json:"result_msg"`
}

type BaseOperationRespDTO struct {
	BaseRespParam
	Data BaseOperationRespData `json:"data"`
}

// 返回成功的操作结果
func SuccessOperationResp() *BaseOperationRespDTO {
	respDTO := &BaseOperationRespDTO{}
	respDTO.Data.Result = 1
	respDTO.Data.ResultMsg = "success"
	return respDTO
}

func FailOperationResp() *BaseOperationRespDTO {
	respDTO := &BaseOperationRespDTO{}
	respDTO.Data.Result = 0
	respDTO.Data.ResultMsg = "fail"
	return respDTO
}

func FailOperationRespWithMsg(msg string) *BaseOperationRespDTO {
	respDTO := &BaseOperationRespDTO{}
	respDTO.Data.Result = 0
	respDTO.Data.ResultMsg = msg
	return respDTO
}

func CheckOperationRespIsOk(resp BaseOperationRespDTO) *err_enum.ErrorObj {
	if resp.Code != err_enum.SUCCESS {
		return &err_enum.ErrorObj{
			Code: resp.Code,
			Msg:  resp.Msg,
		}
	}
	return nil
}
