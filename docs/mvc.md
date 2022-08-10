# mvc代码结构介绍

```bash
.
├── api                            # 接口文档,主要是swagger
├── bin                            # 生成的二进制文件
│   └── app
├── build                          # 构建相关的脚本或者内容
│   └── Dockerfile
├── cmd                            # 入口
│   └── app                        # app1 取名例如前台就叫  front，后台叫 backend，一个目录一个入口
│       ├── app                    # app 固定词
│       │   ├── options            # 处理配置文件
│       │   │   └── options.go     # 这里一般有一堆文件
│       │   └── server.go          # 配置文件入口，还有各种初始化，例如MySQL
│       └── main.go                # 整理入口函数
├── config                         # 配置文件存放位置
│   └── config.yaml                # 配置文件
├── docs                           # 乱七八糟的文档
├── go.mod
├── go.sum
├── internal                       # 内部函数，除了当前项目，其他项目导入会报错
│   ├── config                     # 这个是需要优化的，放在上面
│   │   ├── config.go
│   │   └── option.go
│   ├── handler                    # 也就是controller目录
│   │   └── health                 # 我这里把三个步骤当成一个包，还有另外一种写法，就是model、controller、service
│   │       ├── handler.go
│   │       ├── model.go
│   │       └── service.go
│   ├── pkg                        # 内部的包文件，当前项目自己用
│   │   ├── database               # 这里可能也需要优化的 数据库入口
│   │   │   ├── database.go
│   │   │   └── ping.go
│   │   ├── generator              # 唯一ID生成器
│   │   │   └── id.go
│   │   ├── middleware             # 中间件
│   │   │   ├── cors               # 跨域
│   │   │   │   └── cors.go
│   │   │   └── log                # 日志
│   │   │       └── log.go
│   │   └── response               # 我也不知道这个应该怎么写
│   │       ├── common.go
│   │       ├── ping.go
│   │       └── response.go
│   └── router                     # 统一路由
│       └── router.go
├── pkg                            # 外部可以调用的包。也就是等于开源（开源项目没有internal目录）
│   └── pkg.go
└── README.md
```