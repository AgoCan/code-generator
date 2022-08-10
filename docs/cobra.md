# 简单的命令行工具

## 代码目录结构

```bash
.
├── app                  # 主要函数目录
│   ├── app.go           # 主要函数
│   ├── config.go        # 配置模块
│   └── log.go           # 日志模块
├── cmd                  # 命令行目录
│   ├── root.go          # 命令行入口
│   └── version
│       └── version.go   # 打印版本号
├── config               # 配置文件
│   └── config.yaml
├── go.mod
└── main.go              # 入口函数
```