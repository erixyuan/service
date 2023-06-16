package main

import (
	"fmt"
	"os"
	"path/filepath"
	"service/cache"
	"service/component"
	"service/config"
	"service/db"
	"service/global"
	router "service/route"
	"service/task"
)

func main() {
	if workDir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		_ = fmt.Errorf("Load workspace fail, stop !!!!!!!!!!!!!!")
		return
	} else {
		global.InitViperConfig(workDir, config.ServerGlobalConfig)
		// 设置日志配置并且初始化
		global.InitLoginConfig(config.ServerGlobalConfig.LoggerConfig.MaxAge)

		global.GetLogger().Info("Init Config Success")
		global.GetLogger().Info("Starting DB Client......")
		db.InitDB()
		//初始化redis
		cache.InitRedis(&cache.RedisOpts{
			Host:     config.ServerGlobalConfig.RedisConfig.Host,
			Password: config.ServerGlobalConfig.RedisConfig.Password,
			DB:       config.ServerGlobalConfig.RedisConfig.DB,
		})
		component.Init(config.ServerGlobalConfig, cache.RedisClient)
		global.GetLogger().Info("Start DB Client Success")
		//初始化定时任务
		task.Init()
		//global.GetLogger().Info("Starting Task......")
		//c := cron.New()
		//c.AddFunc("*/2 * * * * *", func() {
		//	task.ProductOrderNotifyTask()
		//})
		//c.AddFunc("*/5 * * * * *", func() {
		//	task.ProductOrderExpiredTask()
		//})
		//c.Start()

		global.GetLogger().Info("Starting Service......")
		router.InitRouter(workDir)

	}
}
