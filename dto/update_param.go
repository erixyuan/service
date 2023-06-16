package dto

type UpdateInfoReqDTO struct {
	BaseReqParam
	Id          int    `json:"id" binding:"required"`
	UpdateKey   string `json:"update_key" binding:"required"`
	UpdateValue string `json:"update_value" binding:"required"`
}
