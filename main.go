package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muyiacc/STBlog/config"
	"github.com/muyiacc/STBlog/internal/db"
	"github.com/muyiacc/STBlog/internal/logger"
	"github.com/muyiacc/STBlog/internal/router"
)

func main() {
	// 读取配置
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("配置读取失败")
	}

	// 初始化日志
	logger := logger.NewLogger(cfg.Logger.LogLevel)

	// 初始化数据库
	db, err := db.InitDB(cfg.Database.Dialect, cfg.Database.DSN)

	//
	r := gin.Default()

	// 加载路由
	err = router.InitRouter(r, db, logger)

	//
	if err != nil {
		panic(err)
	} else {
		r.Run()
	}
}
