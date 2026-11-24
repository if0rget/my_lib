.PHONY: db db-down server web test test-db test-int migrate-up migrate-down migrate-new fmt lint check

GOLANGCI := go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0
TEST_DATABASE_URL := postgres://my_lib:my_lib@localhost:5432/my_lib_test?sslmode=disable

db: ## 启动开发数据库
	docker compose up -d

db-down: ## 停止开发数据库
	docker compose down

server: ## 启动后端（:8080）
	cd server && go run ./cmd/server

web: ## 启动前端（:5173）
	cd web && pnpm dev

test: ## 运行后端测试
	cd server && go test ./...

test-db: ## 创建测试库（已存在则忽略）
	docker compose exec -T postgres createdb -U my_lib my_lib_test || true

test-int: test-db ## 运行包含数据库的全部测试
	cd server && TEST_DATABASE_URL='$(TEST_DATABASE_URL)' go test -p 1 ./...

migrate-up: ## 升级到最新
	cd server && go run ./cmd/migrate up

migrate-down: ## 回滚一个版本
	cd server && go run ./cmd/migrate down 1

migrate-new: ## 新建一对迁移文件：make migrate-new name=add_xxx
	@n=$$(printf '%06d' $$(( $$(ls server/migrations/*.up.sql 2>/dev/null | wc -l) + 1 ))); \
	touch server/migrations/$${n}_$(name).up.sql server/migrations/$${n}_$(name).down.sql; \
	echo "created $${n}_$(name)"

fmt: ## 格式化前后端代码
	cd server && $(GOLANGCI) fmt
	cd web && pnpm format

lint: ## 检查前后端代码
	cd server && $(GOLANGCI) run
	cd web && pnpm lint && pnpm format:check

check: lint test ## 提交前跑一遍：检查 + 测试
