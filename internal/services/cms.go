package services

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
}

// CmsApp的构造函数，返回一个CmsApp实例对象
func NewCmsApp() *CmsApp {
	return &CmsApp{}
}

// 创建sql连接，返回mysql数据库连接对象
func (cmsApp *CmsApp) connDB() *gorm.DB {
	// 使用gorm.Open()获取数据库连接 传入参数mysql驱动连接mysql.Open() 传入数据库连接schema
	mysqlDB, err := gorm.Open(mysql.Open("root:774051432@tcp(localhost:3306)/cms_account?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 获取数据库实例 设置最大连接数
	db, err := mysqlDB.DB()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	// 数据库连接设置为debug模式
	mysqlDB = mysqlDB.Debug()
	return mysqlDB
}
