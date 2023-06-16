package db

import (
	drvsql "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"service/config"
	"service/global"
)

var (
	DBGClient *gorm.DB
	DBDirver  *drvsql.DB
)

func InitDB() {
	global.GetLogger().Info("Start connect db......")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", config.ServerGlobalConfig.DBConfig.Username, config.ServerGlobalConfig.DBConfig.Pwd, config.ServerGlobalConfig.DBConfig.Uri, config.ServerGlobalConfig.DBConfig.DBName)
	dbg, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			TablePrefix:   "hl_",
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	DBGClient = dbg
	if err != nil {
		global.GetLogger().Error("failed opening connection to mysql: " + err.Error())
	}

	global.GetLogger().Info("connect db success:%s" + config.ServerGlobalConfig.DBConfig.Uri)
}
