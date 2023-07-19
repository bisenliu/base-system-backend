## server项目结构

```shell
├── api
│   └── v1
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── middleware
├── model
│   ├── request
│   └── response
├── router
├── service
└── utils
```

| 文件夹       | 说明                    | 描述                                                       |
| ------------ | ----------------------- |----------------------------------------------------------|
| `api`        | api层                   | api层                                                     |
| `--v1`       | v1版本接口              | v1版本接口                                                   |
| `config`     | 配置包                  | config.yaml对应的配置结构体                                      |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化                             |
| `docs`       | swagger文档目录         | swagger文档目录                                              |
| `global`     | 全局对象                | 全局对象                                                     |
| `initialize` | 初始化 | router,redis,gorm,validator, timer的初始化                   |
| `--internal` | 初始化内部函数 | validator 的自定义校验, 初始化默认数据,在此文件夹的函数只能由 `initialize` 层进行调用 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码                                         |
| `model`      | 模型层                  | 模型对应数据表                                                  |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。                                            |
| `--response` | 出参结构体              | 返回给前端的数据结构体                                              |
| `router`     | 路由层                  | 路由层                                                      |
| `service`    | service层               | 存放业务逻辑问题                                                 |
| `utils`      | 工具包                  | 工具函数封装                                                   |

