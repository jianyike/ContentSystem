package dao

import (
	"cmsProject/internal/model"
	"fmt"
	"gorm.io/gorm"
)

// 创建AccountDao对象
type AccountDao struct {
	db *gorm.DB
}

// 创建AccountDao的构造方法 传入数据库连接对象
func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

// 根据userId查询数据是否存在
func (accountDao *AccountDao) IsExit(userId string) (bool, error) {
	var account model.Account
	err := accountDao.db.Where("user_id = ?", userId).First(&account).Error
	// 第一种错误是没查询到，正常返回false
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	// 第二种错误是其他查询错误，返回err
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	// 没错误 返回true
	return true, nil
}

// 插入一个新账户数据
func (accountDao *AccountDao) AddUser(account model.Account) error {
	if err := accountDao.db.Create(&account).Error; err != nil {
		fmt.Printf("add account error [%v]", err)
		return err
	}
	return nil
}

// 根据userId查询用户
func (accountDao *AccountDao) QueryUser(userId string) (*model.Account, error) {
	var account model.Account
	err := accountDao.db.Where("user_id = ?", userId).First(&account).Error
	// 第一种错误是没查询到，返回nil
	if err == gorm.ErrRecordNotFound {
		fmt.Println("未查询到用户")
		return nil, nil
	}
	// 第二种错误是其他查询错误，返回err
	if err != nil {
		fmt.Printf("query user error: [%v]\n", err)
		return nil, err
	}
	// 没错误 返回account
	return &account, nil
}
