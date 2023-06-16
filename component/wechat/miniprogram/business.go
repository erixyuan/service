package miniprogram

type Business struct {
	*Context
}

func NewBusiness(ctx *Context) *Business {
	return &Business{ctx}
}
