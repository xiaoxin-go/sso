# SSO统一单点登录系统

### 项目描述
本系统使用go+gin+mysql开发，权限管理使用casbin实现

### 目录结构
- api
- libs
- database
- model
- pkg
- api
- routers
- main.go

### 安装启动
1. 安装mysql
2. 安装redis
3. 初始化sql文件 doc/init.sql
4. 执行启动命令 go run main.go