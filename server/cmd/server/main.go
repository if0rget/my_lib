// Command server 是图书管理系统的后端入口。
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/if0rget/my_lib/server/internal/config"
	"github.com/if0rget/my_lib/server/internal/database"
)

func main() {
	// .env 不存在时跳过，直接使用系统环境变量
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(cfg.DatabaseURL); err != nil {
		log.Fatal(err)
	}

	db, err := database.Open(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = sqlDB.Close() // 进程退出时关闭数据库连接池
	}()

	r := gin.Default()

	// 健康检查：部署时给容器 / 负载均衡探活用
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 业务接口统一挂在 /api/v1 下
	v1 := r.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok", "data": "pong"})
	})

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
