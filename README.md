# 代码生成器

## 生成ansible代码框

```
go run main.go -i ansible -n ansible-test
```

## 生成 gitbook 文档框架

```
go run main.go -i gitbook -n gitbook-test
```

### 参数

|参数|描述|默认值|
|---|---|---|
| `baidu` | 添加百度统计的token | "" |
| `google` | 添加谷歌统计的token | "" |
| `extra-plugin` | 添加需要添加的额外插件 | "" |
| `keywords` | 添加关键词 | "keywords" |
| `description` | 添加描述词 | "description" |
| `sidebar-title` | 添加侧边栏题目 | "" |
| `sidebar-link` | 添加侧边栏对应链接 | "" |

> 该文档默认安装插件: `-highlight`、`toggle-chapters`、`codeblock-filename`、`sectionx`、`splitter`、`-search`、`-lunr`、`search-pro`、`theme-default`、`prism`、`prism-themes`、`theme-comscore`、`include`、`favicon`、`anchors`、`tbfed-pagefooter`、`hide-element`

## 生成 简单的go代码 脚本框架

```
go run main.go -i simple -n simple-test
```

## 生成 简单的go代码并且通过传参方式执行 脚本框架（需要优化）

```
go run main.go -i simplecobra -n simplecobra-test
```

## 生成 简单go的http代码 脚本框架(未做)

```
go run main.go -i simplehttp -n simplehttp-test
```

## 生成 基于gorm的mvc 代码框架(未做)

```
go run main.go -i mvcgorm -n mvcgorm-test
```

## 生成 基于gorm的mvc 代码框架(未做)

```
go run main.go -i mvcsqlx -n mvcsqlx-test
```
