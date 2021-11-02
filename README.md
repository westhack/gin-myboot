
<div align=center>
<img src="https://i.loli.net/2021/11/02/5dZV1Oqoxc4R76G.png" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.14-blue"/>
<img src="https://img.shields.io/badge/gin-1.6.3-lightBlue"/>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/antdv-1.7.6-green"/>
<img src="https://img.shields.io/badge/gorm-1.20.7-red"/>
</div>


> gin-myboot 基于 gin-vue-admin 分化而来，原版 git：https://github.com/flipped-aurora/gin-vue-admin

[gitee地址](https://gitee.com/westhack/gin-myboot): https://gitee.com/westhack/gin-myboot

[github地址](https://github.com/westhack/gin-myboot): https://github.com/westhack/gin-myboot


# 相关项目

- [spring-myboot](https://github.com/westhack/spring-myboot): https://github.com/westhack/spring-myboot

- [spring-myboot](https://github.com/westhack/laravel-myboot): https://github.com/westhack/laravel-myboot

- [grpc-myboot - grpc 版](https://github.com/westhack/grpc-myboot): https://github.com/westhack/grpc-myboot

# 项目文档
[在线文档](http://docs.limaopu.com) : http://docs.limaopu.com


# 截图预览

![1.png](https://i.loli.net/2021/11/02/4UikFAHnQO7lJsb.png)

![2.png](https://i.loli.net/2021/11/02/sHGh3qwnoNLptRO.png)

![3.png](https://i.loli.net/2021/11/02/z95V1ntGjKr48xo.png)

![4.png](https://i.loli.net/2021/11/02/AH9vaCQGq2en6uR.png)

![5.png](https://i.loli.net/2021/11/02/xhRFwXJfuHIKZcT.png)

## 1. 基本介绍

### 1.1 项目介绍

> gin-myboot 是一个基于 [vue](https://vuejs.org) 和 [gin](https://gin-gonic.com) 开发的全栈前后端分离的后台管理系统，集成jwt鉴权，动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，提供多种示例文件，让您把更多时间专注在业务开发上。

[在线预览](http://demo.limaopu.com): http://demo.limaopu.com

测试用户名：admin

测试密码：123456

## 2. 使用说明

```
- node版本 > v8.6.0
- golang版本 >= v1.17
- IDE推荐：Goland
```

### 2.1 server项目

使用 `Goland` 等编辑工具，打开server目录，不可以打开 gin-myboot 根目录

```bash

# 克隆项目
git clone https://gitee.com/westhack/gin-myboot.git
# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 运行
go run mian.go

# 开发环境推荐使用 `air` http://github.com/cosmtrek/air 
air -c .air.toml

# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```

### 2.2 web项目

```bash
# 进入web文件夹
cd backend-ui

# 安装依赖
yarn install || npm install

# 启动web项目
yarn serve || npm serve
```

### 2.3 swagger自动化API文档

#### 2.3.1 安装 swagger

##### （1）可以访问外国网站

````
go get -u github.com/swaggo/swag/cmd/swag
````

##### （2）无法访问外国网站

由于国内没法安装 go.org/x 包下面的东西，推荐使用 [goproxy.cn](https://goproxy.cn) 或者 [goproxy.io](https://goproxy.io/zh/)

```bash
# 如果您使用的 Go 版本是 1.13 - 1.15 需要手动设置GO111MODULE=on, 开启方式如下命令, 如果你的 Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 如果嫌弃麻烦,可以使用go generate 编译前自动执行代码, 不过这个不能使用 `Goland` 或者 `Vscode` 的 命令行终端
cd server
go generate -run "go env -w .*?"

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2.3.2 生成API文档

```` shell
cd server
swag init
````

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) 即可查看swagger文档


## 3. 技术选型

- 前端：用基于 [Vue](https://vuejs.org) 的 [Antdv](https://github.com/vueComponent/ant-design-vue) 构建基础页面。
- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 数据库：采用`MySql`(5.6.44)版本，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。

## 4. 项目架构

### 4.1 目录结构

```
    ├── server
        ├── config          (配置包)
        ├── core            (核心文件)
        ├── docs            (swagger文档目录)
        ├── global          (全局对象)                    
        ├── initialize      (初始化)                        
        │   └── internal    (初始化内部函数)                            
        ├── middleware      (中间件层)                        
        ├── modules         (模块)                    
        │   ├── common      (公共模块)   
        │   ├── ...         (更多板块)                           
        │   └── system      (系统业务模块)                            
        ├── packfile        (静态文件打包)                        
        ├── resource        (静态资源文件夹)                        
        │   └── excel       (excel导入导出默认路                         
        ├── router          (路由层)                    
        ├── source          (source层)                    
        └── utils           (工具包)                    
            ├── timer       (定时器接口封装)                        
            └── upload      (oss接口封装)                        
    
    └─backend-ui            （前端文件）
        ├─public            （发布模板）
        └─src               （源码包）
            ├─api	        （基础接口相关）
            ├─assets	    （静态文件）
            ├─config        （前段配置）
            ├─components    （组件）
            ├─constants     （系统常量）
            ├─router	    （前端路由）
            ├─locales       （语言包）
            ├─layouts       （布局）
            ├─mixins         (mixin)
            ├─modules        (模块-对于server)
            │   ├── common   (公共模块)    
            │   ├── ...      (更多板块)                         
            │   └── system   (系统业务模块) 
            ├─store     （vuex 状态管理仓）
            ├─utils     （前端工具库）
            └─views     （前端基础页面）

```

## 5. 主要功能

- 权限管理：基于`jwt`和`casbin`实现的权限管理。
- 文件上传下载：实现基于`七牛云`, `阿里云`, `腾讯云` 的文件上传操作(请开发自己去各个平台的申请对应 `token` 或者对应`key`)。
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。
- 配置管理：配置文件可前台修改。
- 缓存管理：管理reids缓存。
- 条件搜索：动态自定义多条件搜索。
- restful示例：可以参考用户管理模块中的示例API。
	- 前端文件参考: [/backend-ui/src/modules/system/user/Index.vue](https://github.com/westhack/gin-myboot/blob/main/backend-ui/src/modules/system/user/Index.vue)
    - 后台文件参考: [/server/modules/system/api/v1/sys_user.go](https://github.com/westhack/gin-myboot/blob/main/server/modules/system/api/v1/sys_user.go)
- 多点登录限制：需要在`config.yaml`中把`system`中的`use-multipoint`修改为true(需要自行配置Redis和Config中的Redis参数，测试阶段，有bug请及时反馈)。
- 分片长传：提供文件分片上传和大文件分片上传功能示例。
- 表单生成器：参考 [/backend-ui/src/modules/demo/views/view1.vue](https://github.com/westhack/gin-myboot/blob/main/backend-ui/src/modules/demo/views/view1.vue) 。
- 代码生成器：后台基础逻辑以及简单curd的代码生成器。
