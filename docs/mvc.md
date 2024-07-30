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
│   ├── model                        # model层，增加处理数据库的内容
│   │   ├── health                   # 案例
│   │   │   ├── health.go            # 具体的操作内容，一个表结构对应一个文件
│   │   │   ├── model.go             # 初始化的操作
│   │   │   └── types.go             # 定义常量等内容
│   │   └── model.go                 # 这里可以定义接口，这里暂时是一个空文件
│   ├── pkg                          # 内部调用模块，其实不推荐加一个pkg，因为这样导入模块的时候 需要中间多加一个pkg字样，看个人习惯
│   │   ├── database                 # 数据库相关，例如MySQL
│   │   │   ├── database.go          # 定义接口
│   │   │   ├── mysql.go             # 初始化MySQL
│   │   │   └── sqlite.go            # 初始化sqlite
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
│       ├── migrate.go               # 处理model层的migrate
│       ├── router.go                # 路由在这里，通过传参的方式把配置文件等服务传到handler进行使用
│       └── server.go                # web框架入口            
├── pkg                              # 跟上面的pkg一样的说法，不推荐，但是很多代码库有这个习惯，看个人使用，这里是可以其他代码库进行调用
│   └── pkg.go                       # 查看 https://github.com/golang-standards/project-layout/issues/10 只看讨论即可
└── README.md                        # readme内容
```

## 使用模块

|功能|工具|
|---|---|
| web框架 | `github.com/gin-gonic/gin` |
| 数据库2 | `github.com/jinzhu/gorm` |
| 日志 | `go.uber.org/zap` |
| 命令行 | `github.com/urfave/cli/v2` |
| 配置 | `github.com/spf13/viper` |
| id生成器 | `github.com/sony/sonyflake` |

## 其中如果不想用数据库
需要整理的地方有以下几处

- 删除 internal/database
- 删除 internal/server/migrate.go
- 整理 cmd/app/app/options/options.go 下面初始化数据库的相关操作
- 整理 internal/server/server.go
- 整理 internal/handler/health/ 下面的 handler.go 和service.go

## 推荐更改

1. router： 路由，可以创建一个目录，然后把route的变量用参数的方式进行传参。例如下面的方式

    ```go
    package router

    import (
        "time"

        "github.com/gin-gonic/gin"
        cache "github.com/patrickmn/go-cache"

        "xxxx/internal/config"
        "xxxx/internal/handler/auth"
    )

    func Auth(group *gin.RouterGroup, conf *config.Config) {
        authCache := cache.New(5*time.Minute, 10*time.Minute)
        service := auth.Auth{
            Config: conf,
            Cache:  authCache,
        }
        group.POST("/login", auth.LoginHandler(&service))
    }

    ```

2. 数据库，这里使用的是接口模式，小项目推荐直接使用接口体方式，不使用接口

    ```go
    package database

    import (
        "fmt"

        "go.uber.org/zap"
        "gorm.io/gorm"
    )

    type Client struct {
        DB     *gorm.DB
        Driver string // 类型，暂时只支持sqlite
    }

    func New(driver string, options ...string) (client *Client, err error) {
        var db *gorm.DB
        con := options[0]
        db, err = NewSqlite(con)
        if err != nil {
            return nil,err
        }
        
        client = &Client{
            DB:     db,
            Driver: driver,
        }
        return client, err
    }

    ```
