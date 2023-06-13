# mvc

## 代码目录结构

```bash
.
├── api                              # 接口文档,主要是swagger
├── bin                              # 生成的二进制文件
├── build                            # 构建相关的脚本或者内容
│   └── Dockerfile
├── cmd                              # 程序的入口
│   └── app                          # app1 取名例如前台就叫  front，后台叫 backend，一个目录一个入口
│       ├── app                      # app 固定词
│       │   ├── options              # 处理选项
│       │   │   └── options.go
│       │   └── server.go            # 基本上可以认为是入口函数
│       └── main.go                  # 入口函数，具体的实现都在上面的server之中
├── config                           # 默认配置文件
│   └── config.yaml
├── docs                             # 文档
├── go.mod                           # go.mod
├── internal                         # 内部资源，外部不允许调用  请查看 https://golang.org/doc/go1.4#internalpackages
│   ├── config                       # 配置文件，主要就是设置一下配置的相关结构体之类
│   │   ├── config.go      
│   │   ├── db.go                    # 数据库的配置项
│   │   └── log.go                   # 日志的配置项
│   ├── handler                      # handler函数，这里是把数据模块，handler 放置一起
│   │   └── health                   # 案例
│   │       ├── handler.go
│   │       └── service.go
│   ├── pkg                          # 内部调用模块，其实不推荐加一个pkg，因为这样导入模块的时候 需要中间多加一个pkg字样，看个人习惯
│   │   ├── database                 # 数据库相关，例如redis和mysql
│   │   │   └── database.go
│   │   ├── generator                # id生成器
│   │   │   └── id.go
│   │   ├── middleware               # 中间件，上面其实也算，这里应该算是gin的使用工具
│   │   │   ├── cors                 # 跨域
│   │   │   │   ├── config.go
│   │   │   │   ├── cors.go
│   │   │   │   └── utils.go
│   │   │   └── log                  # 日志
│   │   │       └── log.go
│   │   └── response                 # 返回内容
│   │       ├── common.go
│   │       └── response.go
│   └── server                       # 这里也可以改成api， 主要就是上面的入口函数调用的地方，其他的函数主要作用是插件，这里是主要处理逻辑部分
│       └── server.go
│   │   └── router.go                # 路由在这里，通过传参的方式把配置文件等服务传到handler进行使用
├── pkg                              # 跟上面的pkg一样的说法，不推荐，但是很多代码库有这个习惯，看个人使用，这里是可以其他代码库进行调用
│   └── pkg.go                       # 查看 https://github.com/golang-standards/project-layout/issues/10 只看讨论即可
└── README.md                        # readme内容
```

## 使用模块

|功能|工具|
|---|---|
| web框架 | `github.com/gin-gonic/gin` |
| 数据库1 | `github.com/jmoiron/sqlx` |
| 数据库2 | `github.com/jinzhu/gorm` |
| 日志 | `go.uber.org/zap` |
| 命令行 | `github.com/urfave/cli/v2` |
| 配置 | `github.com/spf13/viper` |
| id生成器 | `github.com/sony/sonyflake` |
