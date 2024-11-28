# http代码

## 代码目录结构

> 以下代码，去掉了pkg的操作。为什么去掉，请查看mvc的目录结构介绍。

```bash
.
├── cmd                           # 入口
│   └── app                       # app 取名例如前台就叫  front，后台叫 backend，一个目录一个入口
│       ├── app                   # 固定词汇
│       │   ├── options           # 处理选项
│       │   │   └── options.go
│       │   └── server.go         # 基本上可以认为是入口函数
│       └── main.go               # 入口函数，具体的实现都在上面的server之中
├── go.mod                        # go.mod
├── internal                      # 内部资源，外部不允许调用  请查看 https://golang.org/doc/go1.4#internalpackages
│   ├── cors                      # 跨域
│   │   ├── config.go
│   │   ├── cors.go
│   │   └── utils.go
│   ├── handler                   # demo用例
│   │   └── health
│   │       ├── handler.go
│   │       └── service.go
│   ├── response                  # 返回操作
│   │   ├── common.go
│   │   └── response.go
│   └── server                    # gin函数内容
│       ├── router.go
│       └── server.go
└── README.md                     # readme
```

## 使用模块

|功能|工具|
|---|---|
| web框架 | `github.com/gin-gonic/gin` |
| 命令行 | `github.com/urfave/cli/v2` |
