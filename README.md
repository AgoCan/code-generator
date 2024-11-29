
# 代码脚手架

一键生成一个代码目录架构。

## 工具使用

使用`./cheetach --help`查看帮助

|功能|命令|介绍|demo|
|---|---|---|---|
|【golang】基于gin的mvc框架|`cheetach -i mvcgorm`|[项目结构](./docs/mvc.md)|[demo地址](https://github.com/AgoCan/mvc-demo)|
|【ansible】基于shell|`cheetach -i ansible`|[项目结构](./docs/ansible.md)|暂未创建|
|【golang】命令行方式|`cheetach -i command`|[项目结构](./docs/command.md)|暂未创建|
|【gitbook】文档框架|`cheetach -i gitbook`|[项目结构](./docs/gitbook.md)|暂未创建|
|【mdbook】文档框架|`cheetach -i mdbook`|[项目结构](./docs/mdbook.md)|暂未创建|
|【golang】简单函数|`cheetach -i simple`|[项目结构](./docs/simple.md)|暂未创建|
|【golang】简单http|`cheetach -i simplehttp`|[项目结构](./docs/http.md)|暂未创建|

### 安装

```bash
apt-get update 
apt-get install curl -y
# linux x86
version=0.0.5
release_package=https://github.com/go-cheetah/cheetach/releases/download/${version}/cheetach-linux-amd64
curl -o cheetach ${release_package}
./cheetach -h
```

## gitee和github仓库地址

> github: https://github.com/go-cheetah/cheetach.git  
> gitee: https://gitee.com/agocan/code-generator.git  