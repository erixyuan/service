package wechatpay

type PrepareOrderRequest struct {
	//商户订单号
	OutTradeNo string
	//订单总金额，单位为分
	OrderTotal int
	//商品描述
	Description string
	//通知地址.通知URL必须为直接可访问的URL，不允许携带查询串，要求必须为https地址。
	NotifyUrl string
	//附加数据，可选
	Attach string
	//订单优惠标记
	GoodsTag *string
	//优惠功能
	Detail *DiscountDetail
}

// MiniPlaceOrderRequest 小程序下单参数
type MiniPlaceOrderRequest struct {
	PrepareOrderRequest
	//用户的Openid
	OpenId string
}

// NativeOrderRequest Native下单参数
type NativeOrderRequest struct {
	PrepareOrderRequest
}

type DiscountDetail struct {
	// 1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。 2.当订单原价与支付金额不相等，则不享受优惠。 3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	CostPrice *int `json:"cost_price,omitempty"`
	// 商家小票ID。
	InvoiceId   *string       `json:"invoice_id,omitempty"`
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）。
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// 用户购买的数量。
	Quantity *int `json:"quantity"`
	// 商品单价，单位为分。
	UnitPrice *int `json:"unit_price"`
}

type OrderPayCallBackRequest struct {
	//加密算法类型
	Algorithm string `json:"algorithm"`
	//数据密文
	Ciphertext string `json:"ciphertext"`
	//附加数据
	AssociatedData string `json:"associated_data"`
	//原始类型
	OriginalType string `json:"original_type,omitempty"`
	//随机串
	Nonce string `json:"nonce"`
}

type AppOrderRequest struct {
	PrepareOrderRequest
}

type H5OrderRequest struct {
	PrepareOrderRequest
	SceneInfo SceneInfo `json:"scene_info"`
}

type SceneInfo struct {
	// 用户终端IP
	PayerClientIp string `json:"payer_client_ip"`
	H5Info        H5Info `json:"h5_info"`
}

// H5Info
type H5Info struct {
	// 场景类型 示例值：iOS, Android, Wap
	Type string `json:"type"`
}
