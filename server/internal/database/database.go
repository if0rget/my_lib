// Package database 提供 PostgreSQL 连接与基于 golang-migrate 的数据库迁移能力。
package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	pgxv5 "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v5/stdlib" // 注册 pgx 标准库驱动供迁移器使用
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/if0rget/my_lib/server/migrations"
)

// Open 使用 GORM 打开 PostgreSQL 数据库连接，并启用错误翻译。
func Open(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库: %w", err)
	}
	return db, nil
}

// Migrate 将数据库结构迁移升级到最新版本；若已是最新则返回 nil。
func Migrate(databaseURL string) (retErr error) {
	m, err := newMigrator(databaseURL)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := closeMigrator(m); closeErr != nil && retErr == nil {
			retErr = fmt.Errorf("关闭迁移器: %w", closeErr)
		}
	}()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("执行向上迁移: %w", err)
	}
	return nil
}

// MigrateDown 将数据库结构向后回滚 steps 个版本。
func MigrateDown(databaseURL string, steps int) (retErr error) {
	if steps <= 0 {
		return fmt.Errorf("回滚步数必须大于 0: %d", steps)
	}

	m, err := newMigrator(databaseURL)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := closeMigrator(m); closeErr != nil && retErr == nil {
			retErr = fmt.Errorf("关闭迁移器: %w", closeErr)
		}
	}()

	if err := m.Steps(-steps); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("回滚迁移 %d 步: %w", steps, err)
	}
	return nil
}

func newMigrator(databaseURL string) (*migrate.Migrate, error) {
	sqlDB, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("打开迁移数据库连接: %w", err)
	}

	driver, err := pgxv5.WithInstance(sqlDB, &pgxv5.Config{})
	if err != nil {
		_ = sqlDB.Close() // 初始化失败时释放底层连接
		return nil, fmt.Errorf("初始化迁移驱动: %w", err)
	}

	src, err := iofs.New(migrations.FS, ".")
	if err != nil {
		_ = driver.Close() // 初始化失败时释放驱动及连接
		return nil, fmt.Errorf("加载迁移文件: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
	if err != nil {
		_ = src.Close()
		_ = driver.Close()
		return nil, fmt.Errorf("创建迁移器: %w", err)
	}

	return m, nil
}

func closeMigrator(m *migrate.Migrate) error {
	srcErr, dbErr := m.Close()
	return errors.Join(srcErr, dbErr)
}
