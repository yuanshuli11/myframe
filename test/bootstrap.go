package infrastructure

import (
	"myproject/infrastructure/database"
	"myproject/infrastructure/mlog"
	"myproject/infrastructure/viper"
)

var (
	f string
)

func init() {
	f := "/work/src/myframe/config/develop.yml"
	viper.LoadConfigFromYaml(f)
	database.LoadMysql()
	mlog.ConfigLocalFilesystemLogger(viper.Config.GetString("log.applog.path"), viper.Config.GetString("log.applog.level"))

}
