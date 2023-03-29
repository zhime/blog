package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhime/blog/server/core"
	"github.com/zhime/blog/server/global"
	"github.com/zhime/blog/server/initialize"
	"net/http"
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

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": map[string]string{
				"name":  "golang",
				"price": "29.8",
			},
			"msg": "success",
		})
	})

	_ = router.Run()

	//u := model.User{
	//	UserName: "zhangsan",
	//}
	//
	//var userService *controller.UserService
	//user, err := userService.Register(u)
	//if err != nil {
	//	global.Logger.Warn("register error")
	//}
	//fmt.Println(user)
}
