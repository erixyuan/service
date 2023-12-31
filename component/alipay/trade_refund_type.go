package alipay

// TradeRefundAsync 统一收单交易退款接口(异步)请求参数 https://opendocs.alipay.com/pre-apis/api_pre/alipay.trade.refund.apply
type TradeRefundAsync struct {
	AppAuthToken string `json:"-"`                      // 可选
	NotifyURL    string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	RefundAmount string `json:"refund_amount"`          // 必须 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundReason string `json:"refund_reason"`          // 可选 退款的原因说明
	OutRequestNo string `json:"out_request_no"`         // 必须 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	OperatorId   string `json:"operator_id"`            // 可选 商户的操作员编号
	StoreId      string `json:"store_id"`               // 可选 商户的门店编号
	TerminalId   string `json:"terminal_id"`            // 可选 商户的终端编号
}

func (tra TradeRefundAsync) APIName() string {
	return "alipay.trade.refund.apply"
}

func (tra TradeRefundAsync) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = tra.AppAuthToken
	m["notify_url"] = tra.NotifyURL
	return m
}

// TradeRefundAsyncRsp 统一收单交易退款接口(异步)响应参数
type TradeRefundAsyncRsp struct {
	Content struct {
		Code         Code   `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		TradeNo      string `json:"trade_no"`       // 支付宝交易号
		OutTradeNo   string `json:"out_trade_no"`   // 商户订单号
		OutRequestNo string `json:"out_request_no"` // 本笔退款对应的退款请求号
		RefundAmount string `json:"refund_amount"`  // 本次退款请求，对应的退款金额
		RefundStatus string `json:"refund_status"`  // REFUND_PROCESSING 退款处理中；REFUND_SUCCESS 退款处理成功；REFUND_FAIL 退款失败
	} `json:"alipay_trade_refund_apply_response"`
	Sign string `json:"sign"`
}

func (trar *TradeRefundAsyncRsp) IsSuccess() bool {
	return trar.Content.Code == CodeSuccess
}
