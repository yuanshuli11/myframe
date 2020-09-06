package database

import (
	"fmt"
	"myproject/infrastructure/mlog"
	"myproject/infrastructure/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func LoadMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", viper.Config.GetString("db.username"), viper.Config.GetString("db.password"), viper.Config.GetString("db.host"), viper.Config.GetString("db.port"), viper.Config.GetString("db.database"))
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "" + defaultTableName
	}
	db, err := gorm.Open("mysql", dsn)
	//db, err := sql.Open("mysql", dsn)
	db.LogMode(true)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	//	defer DB.Close()
	db.SetLogger(mlog.GetLogrus())
	DB = db

}
