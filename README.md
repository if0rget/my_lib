# my_lib

图书管理系统，一个学习 Go 的练手项目。后端 Go，前端 React。

> 当前状态：只有项目文档，代码尚未开始。

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

## 目录规划

```
my_lib/
├── server/     # Go 后端
├── web/        # React 前端
└── docs/       # 项目文档
```

## 文档

选型理由、数据模型、接口设计、部署方案和学习顺序见 [docs/项目文档.md](docs/项目文档.md)。
