.PHONY: db db-down server web test fmt lint check

GOLANGCI := go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0

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

fmt: ## 格式化前后端代码
	cd server && $(GOLANGCI) fmt
	cd web && pnpm format

lint: ## 检查前后端代码
	cd server && $(GOLANGCI) run
	cd web && pnpm lint && pnpm format:check

check: lint test ## 提交前跑一遍：检查 + 测试
