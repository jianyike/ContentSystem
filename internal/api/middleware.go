package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const sessionKey = "session_id"

type SessionAuth struct {
}

func (m *SessionAuth) Auth(c *gin.Context) {
	sessionId := c.GetHeader(sessionKey)
	if sessionId == "" {
		fmt.Println("request abort !")
		c.AbortWithStatusJSON(http.StatusForbidden, "session_id is null !")
	}

	c.Next()
}
