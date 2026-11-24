// Package testdb 为需要真实 PostgreSQL 的集成测试提供数据库连接与清表辅助。
package testdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/if0rget/my_lib/server/internal/database"
)

// Open 连接测试库、迁移到最新并清空所有业务表；未设置 TEST_DATABASE_URL 时跳过当前测试。
func Open(t *testing.T) *gorm.DB {
	t.Helper()

	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("未设置 TEST_DATABASE_URL")
	}

	require.NoError(t, database.Migrate(url))

	db, err := database.Open(url)
	require.NoError(t, err)
	t.Cleanup(func() {
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close() // 测试结束关闭连接池
		}
	})

	require.NoError(t, db.Exec("TRUNCATE users, categories, books RESTART IDENTITY CASCADE").Error)
	return db
}
