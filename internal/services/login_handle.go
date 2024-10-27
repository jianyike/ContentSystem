package services

import (
	"cmsProject/internal/dao"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// LoginReq 登录的请求消息模型
type LoginReq struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRsp 登录的响应消息模型
type LoginRsp struct {
	SessionId string `json:"session_id"`
	UserId    string `json:"user_id" binding:"required"`
	NickName  string `json:"nick_name"`
}

func (cmsApp *CmsApp) Login(ctx *gin.Context) {
	// 登录的请求消息模型绑定
	var loginReq LoginReq
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	accountDao := dao.NewAccountDao(cmsApp.db)
	account, err := accountDao.QueryUser(loginReq.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if account == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "账号不存在！"})
		return
	}
	// 密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password),
		[]byte(loginReq.Password)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "密码错误！"})
		return
	}

	// 生成session_id
	sessionId, err := cmsApp.generateSessionId(context.Background(), account.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// err为nil则校验通过
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "登录成功",
		"message": &LoginRsp{
			SessionId: sessionId,
			UserId:    account.UserId,
			NickName:  account.Nickname,
		},
	})

}

func (cmsApp *CmsApp) generateSessionId(ctx context.Context, userId string) (string, error) {
	// 生成session_id
	sessionId := uuid.New().String()
	// session_id持久化
	// 得到sessionkey
	sessionKey := fmt.Sprintf("session_id:%s", userId)
	// 存储到redis里，过期时间设置8小时
	if err := cmsApp.rdb.Set(context.Background(), sessionKey, sessionId, time.Hour*8).Err(); err != nil {
		fmt.Printf("session key set error : [%v]\n", err)
		return "", err
	}

	// session_id生成时间持久化
	// 得到session生成时间的key名
	authKey := fmt.Sprintf("session_auth:%s", sessionId)
	// 持久化，值为当前时间的时间戳 过期时间8小时
	if err := cmsApp.rdb.Set(ctx, authKey, time.Now().Unix(), time.Hour*8).Err(); err != nil {
		fmt.Printf("auth key set error : [%v]\n", err)
		return "", err
	}

	return sessionId, nil
}
