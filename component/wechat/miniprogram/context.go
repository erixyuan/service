package miniprogram

import "service/component/wechat/credential"

type Context struct {
	*Config
	credential.AccessTokenHandle
}
