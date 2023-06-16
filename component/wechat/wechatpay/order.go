package wechatpay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"io/ioutil"
	"log"
	"service/global"
)

var (
	WeChatMiniPayClient *WeChatOrder
	WeChatAppPayClient  *WeChatOrder
)

type WeChatOrder struct {
	*PayContext
}

func InitWeChatMiniPayClient(context *PayContext) {
	WeChatMiniPayClient = &WeChatOrder{
		PayContext: context,
	}
}

func InitWeChatAppPayClient(context *PayContext) {
	WeChatAppPayClient = &WeChatOrder{
		PayContext: context,
	}
}

// MiniPlaceOrder 小程序下单
func (weChatOrder *WeChatOrder) MiniPlaceOrder(request *MiniPlaceOrderRequest) (*jsapi.PrepayWithRequestPaymentResponse, error) {
	svc := jsapi.JsapiApiService{Client: weChatOrder.Client}
	// 得到prepay_id，以及调起支付所需的参数和签名
	ctx := context.Background()
	prepayRequest := jsapi.PrepayRequest{
		Appid:       core.String(weChatOrder.AppID),
		Mchid:       core.String(weChatOrder.MchID),
		Description: core.String(request.Description),
		OutTradeNo:  core.String(request.OutTradeNo),
		Attach:      core.String(request.Attach),
		NotifyUrl:   core.String(request.NotifyUrl),
		Amount: &jsapi.Amount{
			Total: core.Int64(cast.ToInt64(request.OrderTotal)),
		},
		Payer: &jsapi.Payer{
			Openid: core.String(request.OpenId),
		},
		GoodsTag: request.GoodsTag,
	}
	if request.Detail != nil {
		var detail jsapi.Detail
		copier.Copy(&detail, &request.Detail)
		prepayRequest.Detail = &detail
	}
	resp, result, err := svc.PrepayWithRequestPayment(ctx, prepayRequest)
	if err != nil {
		body, _ := ioutil.ReadAll(result.Response.Body)
		global.GetLogger().Errorf("MiniPlaceOrder is err:%s", string(body))
		return nil, err
	}
	return resp, nil
}

// NativeOrder Native下单API
func (weChatOrder *WeChatOrder) NativeOrder(request *NativeOrderRequest) (*native.PrepayResponse, error) {
	svc := native.NativeApiService{Client: weChatOrder.Client}
	ctx := context.Background()
	prepayRequest := native.PrepayRequest{
		Appid:       core.String(weChatOrder.AppID),
		Mchid:       core.String(weChatOrder.MchID),
		Description: core.String(request.Description),
		OutTradeNo:  core.String(request.OutTradeNo),
		Attach:      core.String(request.Attach),
		NotifyUrl:   core.String(request.NotifyUrl),
		Amount: &native.Amount{
			Total: core.Int64(cast.ToInt64(request.OrderTotal)),
		},
		GoodsTag: request.GoodsTag,
	}
	if request.Detail != nil {
		var detail native.Detail
		copier.Copy(&detail, &request.Detail)
		prepayRequest.Detail = &detail
	}
	resp, result, err := svc.Prepay(ctx, prepayRequest)
	if err != nil {
		body, _ := ioutil.ReadAll(result.Response.Body)
		global.GetLogger().Errorf("NativeOrder is error: %s\n", string(body))
		return nil, err
	}
	return resp, nil
}

// QueryOrderByOutTradeNo 查询订单
func (weChatOrder *WeChatOrder) QueryOrderByOutTradeNo(outTradeNo string) (*payments.Transaction, error) {
	svc := native.NativeApiService{Client: weChatOrder.Client}
	ctx := context.Background()
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx, native.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(weChatOrder.MchID),
	})
	if err != nil {
		body, _ := ioutil.ReadAll(result.Response.Body)
		global.GetLogger().Errorf("QueryOrderByOutTradeNo is error: %s\n", string(body))
		return nil, err
	}
	return resp, nil
}

// OrderCallBackDecrypt 支付结果通知内容的解密
func (weChatOrder *WeChatOrder) OrderCallBackDecrypt(request *OrderPayCallBackRequest) (*payments.Transaction, error) {
	plaintext, err := utils.DecryptAES256GCM(weChatOrder.ApiV3Key, request.AssociatedData, request.Nonce, request.Ciphertext)
	if err != nil {
		return nil, fmt.Errorf("decrypt request error: %v", err)
	}
	var orderResp payments.Transaction
	if err = json.Unmarshal([]byte(plaintext), &orderResp); err != nil {
		return nil, fmt.Errorf("unmarshal plaintext to content failed: %v", err)
	}
	return &orderResp, nil
}

// CloseOrder @Description: 关闭订单
func (weChatOrder *WeChatOrder) CloseOrder(outTradeNo string) error {
	svc := jsapi.JsapiApiService{Client: weChatOrder.Client}
	ctx := context.Background()
	result, err := svc.CloseOrder(ctx, jsapi.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(weChatOrder.PayContext.Config.MchID),
	})
	if err != nil {
		log.Printf("wechat close order is error, %s\n", err.Error())
		return err
	}
	global.GetLogger().Infof("wechat close order, order no: %s, http status code: %d", outTradeNo, result.Response.StatusCode)
	return nil
}

// AppPlaceOrder @Description: app下单
func (weChatOrder *WeChatOrder) AppPlaceOrder(request *AppOrderRequest) (*app.PrepayWithRequestPaymentResponse, error) {
	svc := app.AppApiService{Client: weChatOrder.Client}
	ctx := context.Background()
	prepayRequest := app.PrepayRequest{
		Appid:       core.String(weChatOrder.AppID),
		Mchid:       core.String(weChatOrder.MchID),
		Description: core.String(request.Description),
		OutTradeNo:  core.String(request.OutTradeNo),
		Attach:      core.String(request.Attach),
		NotifyUrl:   core.String(request.NotifyUrl),
		Amount: &app.Amount{
			Total: core.Int64(cast.ToInt64(request.OrderTotal)),
		},
		GoodsTag: request.GoodsTag,
	}
	if request.Detail != nil {
		var detail app.Detail
		copier.Copy(&detail, &request.Detail)
		prepayRequest.Detail = &detail
	}
	resp, result, err := svc.PrepayWithRequestPayment(ctx, prepayRequest)
	if err != nil {
		body, _ := ioutil.ReadAll(result.Response.Body)
		global.GetLogger().Errorf("AppPlaceOrder is error: %s\n", string(body))
		return nil, err
	}
	return resp, nil
}

// H5PlaceOrder @Description: H5下单
func (weChatOrder *WeChatOrder) H5PlaceOrder(request *H5OrderRequest) (*h5.PrepayResponse, error) {
	svc := h5.H5ApiService{Client: weChatOrder.Client}
	ctx := context.Background()
	prepayRequest := h5.PrepayRequest{
		Appid:       core.String(weChatOrder.AppID),
		Mchid:       core.String(weChatOrder.MchID),
		Description: core.String(request.Description),
		OutTradeNo:  core.String(request.OutTradeNo),
		Attach:      core.String(request.Attach),
		NotifyUrl:   core.String(request.NotifyUrl),
		GoodsTag:    request.GoodsTag,
		Amount: &h5.Amount{
			Total: core.Int64(cast.ToInt64(request.OrderTotal)),
		},
		SceneInfo: &h5.SceneInfo{
			PayerClientIp: core.String(request.SceneInfo.PayerClientIp),
			H5Info: &h5.H5Info{
				Type: core.String(request.SceneInfo.H5Info.Type),
			},
		},
	}
	if request.Detail != nil {
		var detail h5.Detail
		copier.Copy(&detail, &request.Detail)
		prepayRequest.Detail = &detail
	}
	resp, result, err := svc.Prepay(ctx, prepayRequest)
	if err != nil {
		body, _ := ioutil.ReadAll(result.Response.Body)
		global.GetLogger().Errorf("H5PlaceOrder is error: %s\n", string(body))
		return nil, err
	}
	return resp, nil
}
