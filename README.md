## go 学习-博客网站

go 学习练手项目

### 计划实现功能

实现`blog`功能,方便二次开发

- [ ] 文章管理
- [ ] 页面管理
- [ ] 分类管理
- [ ] 标签管理
- [ ] 系统设置
- [ ] 权限校验
- [ ] 登录登出
- [ ] 日志功能

### 已实现功能

- [x] 使用 gorm 完成数据库连接
- [x] 路由区分公用路由和鉴权路由（`jwt`判断是否登录）
- [x] 用户注册：uuid，密码 md5 加密，用户判重
- [x] 用户登录：账号密码校验，`token`签发，自动续期
- [x] 文章

### 技术栈

- gin web 框架
- gorm 连接数据库
-

## 踩坑

### 接入 go-stash 消费 kafka 数据

坑点：

- go-stash 必须要 kafka 的 consumer Group
- kafka 要同时配置对 docker 容器内和对容器外的访问端口，go-stash 访问内部的，项目访问外部的。

问题：
虽然通过 go-stash 接入了 kafka 的数据，但是没有 timestamp，还要排查。
