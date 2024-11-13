
# 代码生成器

> github: https://github.com/AgoCan/code-generator.git  
> gitee: https://gitee.com/agocan/code-generator.git  

代码生成器，主要是解决平时常开代码，总是要复制的操作。

> 个人理解： golang 主推荐比较简单的代码，并且自由风格。所以，只要代码框架是OK的，随性写代码都行。

## 工具使用

使用`go run main.go --help`查看帮助

帮助文件显示:

- `common:` 则代表给所有应用使用
- `gitbook:` 则代表是专门给`gitbook`进行配置的

## 【golang脚手架】基于gin的mvc框架(持续优化)

```
go run main.go -i mvcgorm -n mvcgorm-test
```

生成后样式: https://github.com/AgoCan/mvc-demo

[mvc介绍](./docs/mvc.md)

## 【ansible脚手架】基于shell

```
go run main.go -i ansible -n ansible-test
```

[ansible价绍](./docs/ansible.md)

## 【golang脚手架】命令行方式

```
go run main.go -i command -n command-test
```

[简单的命令行工具go代码介绍](./docs/command.md)

## 【gitbook脚手架】 文档框架

```
go run main.go -i gitbook -n gitbook-test
```

[gitbook介绍](./docs/gitbook.md)

## 【mdbook脚手架】 文档框架

```
go run main.go -i mdbook -n mdbook-test
```

[mdbook介绍](./docs/mdbook.md)

## 【golang脚手架】简单函数

```
go run main.go -i simple -n simple-test
```

[简单的go代码介绍](./docs/simple.md)

## 【golang脚手架】简单http

```
go run main.go -i simplehttp -n simplehttp-test
```

[简单的http的go代码介绍](./docs/http.md)
