// Package migrations 嵌入数据库迁移所需的 SQL 文件。
package migrations

import "embed"

// FS 包含所有数据库迁移 SQL 文件。
//
//go:embed *.sql
var FS embed.FS
