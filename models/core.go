package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局定义Db和err 向外导出
var Db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:zhao200048.@tcp(127.0.0.1:3306)/beego?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

}
func Getunix() int {
	time := time.Now()
	fmt.Println("当前时间为", time)
	return int(time.Unix())
}
