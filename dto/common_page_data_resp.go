package dto

type CommonPageDataResp struct {
	BaseRespParam
	Data PageDataBody `json:"data"`
}

type PageDataBody struct {
	PageRespParam
	DataList interface{} `json:"data_list"`
}
