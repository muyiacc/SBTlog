package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB, logger *zap.Logger) error {
	// 设置路由全局前缀
	api := r.Group("/api")
	// 版本管理
	v1 := api.Group("/v1")
	// 认证：注册、登陆
	err := RegisterAuthRouter(v1)
	if err != nil {
		logger.Warn(err.Error())
	}
	// 用户管理
	err = RegisterUserRouter(v1)
	if err != nil {
		logger.Warn(err.Error())
	}
	// 文章管理
	err = RegisterPostRouter(v1)
	if err != nil {
		logger.Warn(err.Error())
	}
	// 评论管理
	err = RegisterCommentRouter(v1)
	if err != nil {
		logger.Warn(err.Error())
	}
	// 标签管理
	err = RegisterTagRouter(v1)
	if err != nil {
		logger.Warn(err.Error())
	}
	return nil
}
