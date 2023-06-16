package component

import (
	"service/cache"
	"service/component/alipay"
	"service/component/alisms"
	"service/component/wechat/miniprogram"
	"service/component/wechat/wechatapp"
	"service/component/wechat/wechatpay"
	"service/config"
	"service/global"
)

func Init(config *config.GlobalConfig, cache cache.Cache) {
	initAliSms(&config.AliyunSmsConfig)
	//initAliPay(&config.AliPayConfig)
	initWechat(&config.WeChatConfig, cache)
}

func initAliSms(config *config.AliyunSmsConfig) {
	alisms.InitSmsClient(config.AccessKeyId, config.AccessKeySecret)
}

func initAliPay(config *config.AliPayConfig) {
	var err error
	alipay.AliPayClient, err = alipay.NewPayClient(config.AppId, config.AppPrivateKey, true)
	if err != nil {
		global.GetLogger().Errorf("aliPayClient init error: %s", err.Error())
	}
	err = alipay.AliPayClient.LoadAliPayPublicKey(config.AlipayPublicKey)
	if err != nil {
		global.GetLogger().Errorf("aliPayClient LoadAliPayPublicKey error: %s", err.Error())
	}
}

func initWechat(config *config.WeChatConfig, cache cache.Cache) {
	wechatPayConfig := wechatpay.Config{
		MchID:                config.Pay.MchID,
		MerchantSerialNumber: config.Pay.MerchantSerialNumber,
		ApiV3Key:             config.Pay.ApiV3Key,
	}
	payClient := wechatpay.InitClient(&wechatPayConfig)
	wechatpay.InitWeChatMiniPayClient(&wechatpay.PayContext{
		Config: &wechatpay.Config{
			AppID:                config.Miniprogram.AppId,
			MchID:                config.Pay.MchID,
			MerchantSerialNumber: config.Pay.MerchantSerialNumber,
			ApiV3Key:             config.Pay.ApiV3Key,
		},
		Client: payClient,
	})
	wechatpay.InitWeChatAppPayClient(&wechatpay.PayContext{
		Config: &wechatpay.Config{
			AppID:                config.WeChatApp.AppId,
			MchID:                config.Pay.MchID,
			MerchantSerialNumber: config.Pay.MerchantSerialNumber,
			ApiV3Key:             config.Pay.ApiV3Key,
		},
		Client: payClient,
	})

	miniprogram.InitMiniProgram(&miniprogram.Config{
		AppID:     config.Miniprogram.AppId,
		AppSecret: config.Miniprogram.AppSecret,
		Cache:     cache,
	})
	wechatapp.InitWeChatAppClient(config.WeChatApp.AppId, config.WeChatApp.AppSecret)
}
