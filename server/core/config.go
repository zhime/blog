package core

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhime/blog/server/global"
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./settings.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Println("Viper Unmarshal error")
	}
}
