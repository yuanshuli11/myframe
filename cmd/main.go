package main

import (
	"flag"
	"log"
	"myproject/app"
	"myproject/infrastructure/database"
	"myproject/infrastructure/mlog"
	"myproject/infrastructure/viper"

	"github.com/fvbock/endless"
)

var (
	f string
)

func init() {
	//配置文件的路径 默认值设为/usr/local/src/tencent/gocode/src/myframe/config/develop.yml 可修改
	flag.StringVar(&f, `f`, "/work/src/myframe/config/develop.yml", `配置文件`)
	flag.Parse()
	viper.LoadConfigFromYaml(f)
	database.LoadMysql()
	mlog.ConfigLocalFilesystemLogger(viper.Config.GetString("log.applog.path"), viper.Config.GetString("log.applog.level"))

}

func main() {
	router := app.GetRouters()
	err := endless.ListenAndServe(":"+viper.Config.GetString("listen.port"), router)
	if err != nil {
		log.Println("err:", err)
	}
}
