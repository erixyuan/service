package wechatpay

type Config struct {
	AppID                string `json:"app_id"`
	MchID                string `json:"mch_id"`
	MerchantSerialNumber string `json:"merchant_serial_number"`
	ApiV3Key             string `json:"api_v3_key"`
}
