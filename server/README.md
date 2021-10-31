## server项目结构

```shell
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── middleware
├── modules
│   ├── common
│   ├── ...
│   └── system
├── packfile
├── resource
│   └── excel
├── router
├── source
└── utils
    ├── timer
    └── upload
```

| 文件夹       | 说明                    | 描述                        | |
| ------------ | ----------------------- | --------------------------- | ---|
| `config`     | 配置包                  | config.yaml对应的配置结构体 |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化 |
| `docs`       | swagger文档目录         | swagger文档目录 |
| `global`     | 全局对象                | 全局对象 |
| `initialize` | 初始化                  | router,redis,gorm,validator, timer的初始化 |
| `--internal` | 初始化内部函数           | gorm 的 longger 自定义,在此文件夹的函数只能由 `initialize` 层进行调用 |
| `middleware` | 中间件层                | 用于存放 `gin` 中间件代码 |
| `modules`    | 系统模块                | 系统接口相关             |
| `--common`   | 系统公共板块             | 用户登录注册，验证码等相关接口。  |
| `--...`      | 更多模块                | 更多模块  |
| `--system`   | 系统管理板块             | 系统设置，用户管理，权限设置等相关接口      |
| `packfile`   | 静态文件打包             | 静态文件打包 |
| `resource`   | 静态资源文件夹            | 负责存放静态文件                |
| `--excel`    | excel导入导出默认路径     | excel导入导出默认路径 |
| `router`     | 路由层                  | 路由层 |
| `source`     | source层               | 存放初始化数据的函数 |
| `utils`      | 工具包                  | 工具函数封装            |
| `--timer`    | timer | 定时器接口封装   |
| `--upload`   | oss                    | oss接口封装        |

