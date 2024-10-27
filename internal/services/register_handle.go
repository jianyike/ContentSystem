package services

import (
	"cmsProject/internal/dao"
	"cmsProject/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterReq struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
}

type RegisterRsp struct {
	Message string `json:"message" binding:"required"`
}

func (cmsApp *CmsApp) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 密码加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
	// 账号校验，是否存在
	// 获取数据库连接对象
	// new一个AccountDao的实例对象
	accountDao := dao.NewAccountDao(cmsApp.db)
	// 调用AccountDao的方法根据userId查询数据是否存在
	isExist, err := accountDao.IsExit(req.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if isExist { // 账号已存在 返回客户端错误码400
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号已存在"})
		return
	}

	// 账号密码持久化
	nowTime := time.Now()
	if err := accountDao.AddUser(model.Account{
		UserId:    req.UserId,
		Password:  hashedPassword,
		Nickname:  req.NickName,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	}); err != nil {
		fmt.Println(nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	fmt.Printf("register info %+v, hassedPassword is [%v]", req, hashedPassword)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"message": &RegisterRsp{
			Message: "register successed !",
		},
	})
}

// 密码加密
func encryptPassword(password string) (string, error) {
	// 使用bcrypt生成加密后的哈希值进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 错误处理
	if err != nil {
		fmt.Printf("generate from password error : %v", err.Error())
		return "", err
	}
	return string(hashedPassword), err
}
