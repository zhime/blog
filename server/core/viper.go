package core

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhime/blog/server/config"
)

func Viper() *config.Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./settings.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	var conf *config.Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Println("Viper Unmarshal error")
	}
	return conf
}
