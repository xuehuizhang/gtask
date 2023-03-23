package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

var (
	TDb *gorm.DB
)

func init() {
	//执行main之前 先执行init方法
	dataSourceName := fmt.Sprintf("root:Zhangxh106!@tcp(101.35.196.142:61166)/goblog?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	config := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(dataSourceName), config)
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}

	////最大空闲连接数，默认不配置，是2个最大空闲连接
	//db.SetMaxIdleConns(5)
	////最大连接数，默认不配置，是不限制最大连接数
	//db.SetMaxOpenConns(100)
	//// 连接最大存活时间
	//db.SetConnMaxLifetime(time.Minute * 3)
	////空闲连接最大存活时间
	//db.SetConnMaxIdleTime(time.Minute * 1)
	//err = db.Ping()
	//if err != nil {
	//	log.Println("数据库无法连接")
	//	_ = db.Close()
	//	panic(err)
	//}
	TDb = db
}
