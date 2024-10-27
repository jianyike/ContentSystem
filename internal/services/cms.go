package services

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
	db  *gorm.DB
	rdb *redis.Client
}

// CmsApp的构造函数，返回一个CmsApp实例对象
func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	app.connDB()
	app.connRDB()
	return app
}

// 创建sql连接，获取mysql数据库连接对象
func (cmsApp *CmsApp) connDB() {
	// 使用gorm.Open()获取数据库连接 传入参数mysql驱动连接mysql.Open() 传入数据库连接schema
	mysqlDB, err := gorm.Open(mysql.Open("root:774051432@tcp(localhost:3306)/cms_account?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取数据库实例 设置最大连接数
	db, err := mysqlDB.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	// 数据库连接设置为debug模式
	mysqlDB = mysqlDB.Debug()
	cmsApp.db = mysqlDB
}

// 创建redis连接 获取redis数据库连接对象
func (cmsApp *CmsApp) connRDB() {
	// 连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	// 使用ping命令测试一下是否连接成功
	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("redis connected successed : %s\n", result)
	cmsApp.rdb = rdb
}
