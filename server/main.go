package main

import (
	"fmt"
	"github.com/zhime/blog/server/core"
	"github.com/zhime/blog/server/global"
	"github.com/zhime/blog/server/initialize"
)

func main() {
	global.Config = core.Viper() // 配置文件初始化
	fmt.Println(global.Config.Mysql)

	global.Logger = core.Zap() // 日志初始化
	fmt.Println(global.Logger)

	global.DB = core.Gorm() // gorm初始化
	fmt.Println(global.DB)
	if global.DB != nil {
		initialize.RegisterTables(global.DB) // 初始化表
		db, _ := global.DB.DB()
		defer db.Close()
	}
}
