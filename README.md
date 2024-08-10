
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

## 生成 基于gorm的mvc 代码框架(持续优化)

```
go run main.go -i mvcgorm -n mvcgorm-test
```

生成后样式: https://github.com/AgoCan/mvc-demo

[mvc介绍](./docs/mvc.md)

## 生成ansible代码框

```
go run main.go -i ansible -n ansible-test
```

[ansible价绍](./docs/ansible.md)

## 生成 gitbook 文档框架

```
go run main.go -i gitbook -n gitbook-test
```

[gitbook介绍](./docs/gitbook.md)

## 生成 mdbook 文档框架

```
go run main.go -i mdbook -n mdbook-test
```

[mdbook介绍](./docs/mdbook.md)

## 生成 简单的go代码 脚本框架

```
go run main.go -i simple -n simple-test
```

[简单的go代码介绍](./docs/simple.md)

## 生成 简单的go代码并且通过传参方式执行 脚本框架（需要优化）

```
go run main.go -i simplecobra -n simplecobra-test
```

[简单的命令行工具go代码介绍](./docs/cobra.md)

## 生成 简单go的http代码 脚本框架（需要优化）

```
go run main.go -i simplehttp -n simplehttp-test
```

[简单的http的go代码介绍](./docs/http.md)
