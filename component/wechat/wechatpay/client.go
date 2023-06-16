package wechatpay

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"os"
	"path"
	"path/filepath"
)

func InitClient(config *Config) *core.Client {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	workDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		_ = fmt.Errorf("Load client  fail, stop !!!!!!!!!!!!!!")
		return nil
	}
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(path.Join(workDir, "apiclient_key.pem"))
	if err != nil {
		log.Println("load merchant private key error")
		return nil
	}
	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(config.MchID, config.MerchantSerialNumber, mchPrivateKey, config.ApiV3Key),
	}
	weChatPayClient, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}
	return weChatPayClient
}
