package alisms

import (
	"encoding/json"
	"fmt"
	"github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg/errors"
	"service/global"
)

var SmsClient *SmsClientServe

type SmsClientServe struct {
	*client.Client
}

func InitSmsClient(accessKeyId string, accessKeySecret string) {
	clt, err := CreateClient(accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Printf("sms.CreateClient is err: %s\n", err.Error())
		return
	}
	SmsClient = &SmsClientServe{Client: clt}
}

func (service *SmsClientServe) SendSMS(phoneNumber string, signName string, templateCode string, templateParam map[string]string) (*client.SendSmsResponse, error) {
	smsRequest := &client.SendSmsRequest{
		PhoneNumbers: tea.String(phoneNumber),
		SignName:     tea.String(signName),
		TemplateCode: tea.String(templateCode),
	}
	if templateParam != nil && len(templateParam) > 0 {
		marshal, _ := json.Marshal(templateParam)
		smsRequest.TemplateParam = tea.String(string(marshal))
	}
	sendResp, err := service.Client.SendSms(smsRequest)
	if err != nil {
		global.GetLogger().Errorf("ali service is error: %s", err.Error())
		return nil, err
	}
	if *sendResp.Body.Code != "OK" {
		global.GetLogger().Errorf("ali service is fail, error_code: %s, error_message: %s", *sendResp.Body.Code, *sendResp.Body.Message)
		return nil, errors.New("send sms is fail")
	}
	return sendResp, nil
}
