package model

import "time"

// Account 数据模型
type Account struct {
	ID        int64     `gorm:"id;primary_key"`
	UserId    string    `gorm:"user_id"`
	Password  string    `gorm:"password"`
	Nickname  string    `gorm:"nickname"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"created_at"`
}

// TableName 重写需要查找的数据库表名设置
func (a *Account) TableName() string {
	table := "cms_account.account"
	return table
}
