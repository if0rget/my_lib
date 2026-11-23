.PHONY: db db-down server web test fmt

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

fmt: ## 格式化后端代码
	cd server && gofmt -w .
