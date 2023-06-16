package static

import "fmt"

const (
	SmsPhoneLoginCodeKey = "sms:phone:login:%s"  //手机号登录验证码cache_key
	WxAppCodeAuthKey     = "wx:app:code:%s:auth" //移动应用微信code 获取 access_token
)

func GetSmsPhoneCodeKey(phone string) string {
	return fmt.Sprintf(SmsPhoneLoginCodeKey, phone)
}

func GetWxAppCodeAuthKey(code string) string {
	return fmt.Sprintf(WxAppCodeAuthKey, code)
}
