package miniprogram

import "service/component/wechat/credential"

var (
	WeChatMiniProgram *MiniProgram
)

// MiniProgram 微信小程序相关API
type MiniProgram struct {
	ctx *Context
}

// InitMiniProgram 实例化小程序API
func InitMiniProgram(cfg *Config) {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyMiniProgramPrefix, cfg.Cache)
	ctx := &Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	WeChatMiniProgram = &MiniProgram{ctx}
}

// SetAccessTokenHandle 自定义access_token获取方式
func (miniProgram *MiniProgram) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	miniProgram.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (miniProgram *MiniProgram) GetContext() *Context {
	return miniProgram.ctx
}

// GetAuth 登录/用户信息相关接口
func (miniProgram *MiniProgram) GetAuth() *Auth {
	return NewAuth(miniProgram.ctx)
}

// GetBusiness @Description: 业务接口
func (miniProgram *MiniProgram) GetBusiness() *Business {
	return NewBusiness(miniProgram.ctx)
}
