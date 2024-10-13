package router

import "github.com/gin-gonic/gin"

func RegisterAuthRouter(group *gin.RouterGroup) error {
	auth := group.Group("/auth")
	// 注册
	auth.POST("/signup", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})
	// 登陆
	auth.POST("/login")
	return nil
}
