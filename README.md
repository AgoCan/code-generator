# 代码生成器

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

## 生成 简单的go代码 脚本框架

```
go run main.go -i simple -n simple-test
```

## 生成 简单的go代码并且通过传参方式执行 脚本框架（需要优化）

```
go run main.go -i simplecobra -n simplecobra-test
```

## 生成 简单go的http代码 脚本框架（需要优化）

```
go run main.go -i simplehttp -n simplehttp-test
```

## 生成 基于gorm的mvc 代码框架(需要优化)

```
go run main.go -i mvcgorm -n mvcgorm-test
```

[mvc介绍](./docs/mvc.md)

## 生成 基于sqlx的mvc 代码框架（需要优化）

```
go run main.go -i mvcsqlx -n mvcsqlx-test
```

[mvc介绍](./docs/mvc.md)