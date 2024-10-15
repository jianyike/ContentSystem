package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
	// TODO : 账号校验，是否存在
	// TODO : 账号密码持久化
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
