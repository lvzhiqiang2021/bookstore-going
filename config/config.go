package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if nil != err {
		log.Fatalf("读取配置异常 %s", err)
	}
	return
}
