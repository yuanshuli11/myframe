package viper

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfigFromYaml(path string) error {
	Config = viper.New()
	//设置配置文件的名字
	Config.SetConfigFile(path)

	//设置配置文件类型
	Config.SetConfigType("yaml")

	if err := Config.ReadInConfig(); err != nil {
		log.Fatalf("config fail:%v", err)
		return err
	}

	return nil
}
