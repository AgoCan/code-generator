
# 代码脚手架

一键生成一个代码目录架构。

## 工具使用

使用`go run main.go --help`查看帮助

帮助文件显示:

- `common:` 则代表给所有应用使用
- `gitbook:` 则代表是专门给`gitbook`进行配置的

|功能|demo|介绍|
|---|---|---|
|【golang脚手架】基于gin的mvc框架|`go run main.go -i mvcgorm -n mvcgorm-test`|[代码结构介绍](./docs/mvc.md)|
|【ansible脚手架】基于shell|`go run main.go -i ansible -n ansible-test`|[代码结构介绍](./docs/ansible.md)|
|【golang脚手架】命令行方式|`go run main.go -i command -n command-test`|[代码结构介绍](./docs/command.md)|
|【gitbook脚手架】文档框架|`go run main.go -i gitbook -n gitbook-test`|[代码结构介绍](./docs/gitbook.md)|
|【mdbook脚手架】 文档框架|`go run main.go -i mdbook -n mdbook-test`|[代码结构介绍](./docs/mdbook.md)|
|【golang脚手架】简单函数|`go run main.go -i simple -n simple-test`|[代码结构介绍](./docs/simple.md)|
|【golang脚手架】简单http|`go run main.go -i simplehttp -n simplehttp-test`|[代码结构介绍](./docs/http.md)|



mvn生成后样式: https://github.com/AgoCan/mvc-demo


## gitee和github仓库地址

> github: https://github.com/AgoCan/code-generator.git  
> gitee: https://gitee.com/agocan/code-generator.git  