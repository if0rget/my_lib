# my_lib

图书管理系统，一个学习 Go 的练手项目。后端 Go，前端 React。

> 当前状态：前后端项目已初始化（后端只有 `/healthz` 和 `/api/v1/ping`，前端是 Vite 默认页面），业务功能尚未开始。

## 功能

- 账号：注册、登录，分管理员和读者两种角色
- 图书：列表、搜索、分类筛选、详情；管理员可增删改和调整库存
- 借阅：借书、还书、我的借阅记录；管理员可查看全部记录和逾期列表
- 分类管理、用户管理（管理员）

## 技术栈

| 部分 | 选型 |
|---|---|
| 后端 | Go · Gin · GORM · PostgreSQL |
| 数据库迁移 | golang-migrate |
| 登录鉴权 | JWT + bcrypt |
| 前端 | React 19 · TypeScript · Vite · Ant Design 5 · TanStack Query |
| 部署 | 前端 `embed` 进后端，单个镜像 + Docker Compose |

## 本地运行

需要 Go 1.27+、Node 22+、pnpm、Docker。

```bash
make db        # 启动 PostgreSQL（目前后端还没用到）
make server    # 启动后端 http://localhost:8080
pnpm -C web install   # 首次需要安装前端依赖
make web       # 启动前端 http://localhost:5173，/api 自动转发到后端
```

验证后端：`curl localhost:8080/api/v1/ping`

## 目录规划

```
my_lib/
├── server/     # Go 后端
├── web/        # React 前端
└── docs/       # 项目文档
```

## 文档

- [docs/项目文档.md](docs/项目文档.md)：选型理由、数据模型、接口设计、部署方案和学习顺序
- [docs/开发计划-一期.md](docs/开发计划-一期.md)：一期（后端基础、账号、分类、图书）的任务拆解
- [docs/代码规范.md](docs/代码规范.md)：代码规范；提交前跑 `make check`
