package lanpice_core

import (
	"service/config"
	"sync"
)

type WeilaService struct {
	AppId      string
	AppSecret  string
	HostDomain string
}

var lanpiceService WeilaService
var once sync.Once

func GetWeilaServiceInstance() *WeilaService {
	once.Do(func() {
		lanpiceService = WeilaService{
			AppId:      config.ServerGlobalConfig.WeilaConfig.AppId,
			AppSecret:  config.ServerGlobalConfig.WeilaConfig.AppKey,
			HostDomain: config.ServerGlobalConfig.CoreConfig.Host,
		}
	})
	return &lanpiceService
}
