package api

import (
	"cmsProject/internal/services"
	"github.com/gin-gonic/gin"
)

const (
	rootPath   string = "/api/"
	noAuthPath string = "/out/api/"
)

func CmsRouter(r *gin.Engine) {
	// 创建自定义的中间件对象
	m := NewSessionAuth()
	// 创建路由组
	root := r.Group(rootPath).Use(m.Auth)
	// 调用CmsApp的构造函数获取CmsApp实例对象
	cmsApp := services.NewCmsApp()

	{ // /api/组下的路由方法
		// get方法
		root.GET("cms/hello", cmsApp.Hello)
		// 其他路由方法...
	}

	// 创建/out/api/路由组
	noAuthRoot := r.Group(noAuthPath)
	{
		//post方法
		noAuthRoot.POST("/cms/register", cmsApp.Register)
		// post方法
		noAuthRoot.POST("/cms/login", cmsApp.Login)
	}
}
