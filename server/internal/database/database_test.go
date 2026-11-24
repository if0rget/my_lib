package database_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/if0rget/my_lib/server/internal/database"
)

func TestMigrate_UpDownUp(t *testing.T) {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("未设置 TEST_DATABASE_URL")
	}
	db, err := database.Open(url)
	require.NoError(t, err)
	t.Cleanup(func() {
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close() // 测试结束关闭连接池
		}
	})

	exists := func(table string) bool {
		var name *string
		require.NoError(t, db.Raw("SELECT to_regclass(?)::text", "public."+table).Scan(&name).Error)
		return name != nil
	}

	require.NoError(t, database.Migrate(url))
	require.NoError(t, database.Migrate(url)) // 重复执行不报错
	for _, tbl := range []string{"users", "categories", "books"} {
		assert.True(t, exists(tbl), tbl)
	}
	require.NoError(t, database.MigrateDown(url, 3))
	for _, tbl := range []string{"users", "categories", "books"} {
		assert.False(t, exists(tbl), tbl)
	}
	require.NoError(t, database.Migrate(url))
}
