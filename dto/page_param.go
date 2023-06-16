package dto

type PageParam struct {
	PageNumber int `json:"page_size"`
	PageIndex  int `json:"current"`
}

func (p *PageParam) GetOfferSet() int {
	if p.PageIndex == 0 {
		return 0
	}
	return p.PageNumber * (p.PageIndex - 1)
}

func (p *PageParam) GetLimit() int {
	if p.PageNumber == 0 {
		return 10
	}
	return p.PageNumber
}

type PageRespParam struct {
	PageParam
	Total int `json:"total"`
}
