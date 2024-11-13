# 简单的命令行工具

## 代码目录结构

> 按照 https://github.com/golang-standards/project-layout 规范进行目录结构设置

```bash
.
├── build                        # 构建golang二进制等操作
│   └── build.sh
├── cmd                          # 程序的入口
│   └── app                      # 应用名称
│       ├── app                  # app固定词
│       │   ├── options          # 处理选项
│       │   │   └── options.go   # 如不使用配置文件，可以在这里进行代码删除
│       │   └── server.go        # 服务启动入口
│       └── main.go              # 执行入口
├── config                       # 默认配置文件
│   └── config.yaml
├── docs                         # 文档相关
│   └── README.md
├── go.mod                       # go.mod
├── internal                     # 内部函数
│   ├── command                  # 执行command主要函数
│   │   └── command.go
│   ├── config                   # 配置文件
│   │   ├── common.go
│   │   └── config.go
│   └── log                      # 日志库，暂未使用
│       └── log.go
└── README.md                    # readme
```

## 使用模块

|功能|工具|
|---|---|
| 日志 | `go.uber.org/zap` |
| 命令行 | `github.com/urfave/cli/v2` |
| 配置 | `github.com/spf13/viper` |

