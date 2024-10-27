package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

const sessionKey = "session_id"

type SessionAuth struct {
	rdb *redis.Client
}

func NewSessionAuth() *SessionAuth {
	sessionAuth := &SessionAuth{}
	sessionAuth.connRDB()
	return sessionAuth
}

func (s *SessionAuth) Auth(c *gin.Context) {
	sessionId := c.GetHeader(sessionKey)
	if sessionId == "" {
		fmt.Println("request abort !")
		c.AbortWithStatusJSON(http.StatusForbidden, "session_id is null !")
	}
	// 得到authKey
	authKey := fmt.Sprintf("session_auth:%s", sessionId)
	// 根据key查询redis
	result, err := s.rdb.Get(context.Background(), authKey).Result()
	if err != nil && err != redis.Nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	// 如果值为"" 则表示已经过期，鉴权失败
	if result == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "鉴权失败")
	}
	// 否则 鉴权通过
	c.Next()
}

// 创建redis连接 获取redis数据库连接对象
func (s *SessionAuth) connRDB() {
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
	s.rdb = rdb
}
