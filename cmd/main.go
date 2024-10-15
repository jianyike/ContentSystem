package main

import (
	"cmsProject/internal/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 调用自定义的路由管理函数
	api.CmsRouter(r)
	err := r.Run()
	if err != nil {
		fmt.Printf("r run error = %v", err)
		return
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
