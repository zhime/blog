package main

import (
	"fmt"
	"github.com/zhime/blog/server/core"
	"github.com/zhime/blog/server/global"
)

func main() {
	global.Config = core.Viper() // 配置文件初始化
	fmt.Println(global.Config.Mysql)

	global.DB = core.Gorm() // gorm初始化
	fmt.Println(global.DB)
}
