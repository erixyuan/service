package alipay

import (
	"net/url"
	"strings"
)

// TradePagePay 统一收单下单并支付页面接口 https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func (cli *Client) TradePagePay(param TradePagePay) (result *url.URL, err error) {
	p, err := cli.URLValues(param)
	if err != nil {
		return nil, err
	}

	result, err = url.Parse(cli.apiDomain + "?" + HandlerSpace(p))
	if err != nil {
		return nil, err
	}
	return result, err
}

// TradeAppPay App支付接口 https://docs.open.alipay.com/api_1/alipay.trade.app.pay
func (cli *Client) TradeAppPay(param TradeAppPay) (result string, err error) {
	p, err := cli.URLValues(param)
	if err != nil {
		return "", err
	}
	return HandlerSpace(p), err
}

// TradeFastPayRefundQuery 统一收单交易退款查询接口 https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
func (cli *Client) TradeFastPayRefundQuery(param TradeFastPayRefundQuery) (result *TradeFastPayRefundQueryRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradeOrderSettle 统一收单交易结算接口 https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (cli *Client) TradeOrderSettle(param TradeOrderSettle) (result *TradeOrderSettleRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradeClose 统一收单交易关闭接口 https://docs.open.alipay.com/api_1/alipay.trade.close/
func (cli *Client) TradeClose(param TradeClose) (result *TradeCloseRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradeCancel 统一收单交易撤销接口 https://docs.open.alipay.com/api_1/alipay.trade.cancel/
func (cli *Client) TradeCancel(param TradeCancel) (result *TradeCancelRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradeRefund 统一收单交易退款接口 https://docs.open.alipay.com/api_1/alipay.trade.refund/
func (cli *Client) TradeRefund(param TradeRefund) (result *TradeRefundRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradePreCreate 统一收单线下交易预创建接口 https://docs.open.alipay.com/api_1/alipay.trade.precreate/
func (cli *Client) TradePreCreate(param TradePreCreate) (result *TradePreCreateRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradeQuery 统一收单线下交易查询接口 https://docs.open.alipay.com/api_1/alipay.trade.query/
func (cli *Client) TradeQuery(param TradeQuery) (result *TradeQueryRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradeCreate 统一收单交易创建接口 https://docs.open.alipay.com/api_1/alipay.trade.create/
func (cli *Client) TradeCreate(param TradeCreate) (result *TradeCreateRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradePay 统一收单交易支付接口 https://docs.open.alipay.com/api_1/alipay.trade.pay/
func (cli *Client) TradePay(param TradePay) (result *TradePayRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradeOrderInfoSync 支付宝订单信息同步接口 https://docs.open.alipay.com/api_1/alipay.trade.orderinfo.sync/
func (cli *Client) TradeOrderInfoSync(param TradeOrderInfoSync) (result *TradeOrderInfoSyncRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

// TradeRefundAsync 统一收单交易退款(异步)接口 https://opendocs.alipay.com/pre-apis/api_pre/alipay.trade.refund.apply
func (cli *Client) TradeRefundAsync(param TradeRefundAsync) (result *TradeRefundAsyncRsp, err error) {
	err = cli.doRequest("POST", param, &result)
	return result, err
}

func HandlerSpace(p url.Values) string {
	return strings.Replace(p.Encode(), "+", "%20", -1)
}
