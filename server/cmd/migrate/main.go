// Command migrate 提供命令行数据库迁移工具，支持 up 和 down N。
package main

import (
	"log"
	"log/slog"
	"os"
	"strconv"

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

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("用法: migrate up | migrate down <N>")
	}

	switch args[0] {
	case "up":
		if len(args) != 1 {
			log.Fatal("用法: migrate up")
		}
		if err := database.Migrate(cfg.DatabaseURL); err != nil {
			log.Fatal(err)
		}
		slog.Info("数据库迁移升级完成")
	case "down":
		if len(args) != 2 {
			log.Fatal("用法: migrate down <N>")
		}
		steps, err := strconv.Atoi(args[1])
		if err != nil || steps <= 0 {
			log.Fatalf("无效的回滚步数: %q", args[1])
		}
		if err := database.MigrateDown(cfg.DatabaseURL, steps); err != nil {
			log.Fatal(err)
		}
		slog.Info("数据库迁移回滚完成", "steps", steps)
	default:
		log.Fatalf("未知子命令 %q，用法: migrate up | migrate down <N>", args[0])
	}
}
