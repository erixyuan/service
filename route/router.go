package router

import (
	"fmt"
	"log"
	"path"
	"service/api"
	"service/config"
	"service/db"
	"service/global"
	"service/middlebase"
	"service/utils"
	"strconv"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter(projectPath string) {
	Router := gin.Default()
	pprof.Register(Router)
	//8<<20  namely  8*2^20=8M 设置上传文件大小
	Router.MaxMultipartMemory = 8 << 20
	Router.Use(
		middlebase.CorsMiddleWare(),
		middlebase.RecoverMiddleWare(),
		middlebase.GinLoggerMiddleWare(),
		//middlebase.JwtAuthMiddleware(),
	)
	Router.GET("/ping", api.Ping)
	Router.GET("/debug/dbstate", func(ctx *gin.Context) {
		out := fmt.Sprintf("entdb.Stats().OpenConnections: %v\n sqlxdb..Stats().OpenConnections: %v\n",
			db.DBDirver.Stats().OpenConnections,
		)
		ctx.String(200, out)
	})
	ApiV1Group := Router.Group("/v1")
	initRouteV1(ApiV1Group)
	gin.SetMode(gin.ReleaseMode)
	s := endless.NewServer(fmt.Sprintf(":%d", +config.ServerGlobalConfig.ServerConfig.Port), Router)
	s.BeforeBegin = func(add string) {
		time.Sleep(2 * time.Second)
		log.Printf("Actual pid is %d", syscall.Getpid())
		// save it somehow
		if syscall.Getpid() > 50 {
			utils.WriteFile(path.Join(projectPath, "pid"), strconv.Itoa(syscall.Getpid()))
		} else {
			log.Printf("Actual pid error is %d", syscall.Getpid())
		}
	}
	s.ReadHeaderTimeout = 5 * time.Second
	s.WriteTimeout = 90 * time.Second
	s.ReadTimeout = 5 * time.Second
	s.MaxHeaderBytes = 1 << 20
	if err := s.ListenAndServe(); err != nil {
		global.GetLogger().Error("Server end:" + err.Error())
	} else {
		global.GetLogger().Info("Server normal")
	}
}
