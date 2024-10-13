package router

import "github.com/gin-gonic/gin"

func RegisterUserRouter(group *gin.RouterGroup) error {
	users := group.Group("/users")
	// 修改密码
	users.POST("/")
	// 查看用户信息
	users.GET("/")
	return nil
}
