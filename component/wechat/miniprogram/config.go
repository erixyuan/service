package miniprogram

import "service/cache"

// Config .config for 小程序
type Config struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Cache     cache.Cache
}
