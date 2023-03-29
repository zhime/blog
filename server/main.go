package main

import (
	"fmt"
	"github.com/zhime/blog/server/core"
	"github.com/zhime/blog/server/global"
)

func main() {
	core.InitConfig()
	fmt.Println(global.Config.Mysql)
}
