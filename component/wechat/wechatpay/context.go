package wechatpay

import "github.com/wechatpay-apiv3/wechatpay-go/core"

type PayContext struct {
	*Config
	*core.Client
}
