
# 代码脚手架

一键生成一个代码目录架构。

## 工具使用

使用`./cheetah --help`查看帮助

|功能|命令|介绍|demo|
|---|---|---|---|
|【golang】基于gin的mvc框架|`cheetah -i mvcgorm`|[项目结构](./docs/mvc.md)|[demo地址](https://github.com/AgoCan/mvc-demo)|
|【ansible】基于shell|`cheetah -i ansible`|[项目结构](./docs/ansible.md)|暂未创建|
|【golang】命令行方式|`cheetah -i command`|[项目结构](./docs/command.md)|暂未创建|
|【gitbook】文档框架|`cheetah -i gitbook`|[项目结构](./docs/gitbook.md)|暂未创建|
|【mdbook】文档框架|`cheetah -i mdbook`|[项目结构](./docs/mdbook.md)|暂未创建|
|【golang】简单函数|`cheetah -i simple`|[项目结构](./docs/simple.md)|暂未创建|
|【golang】简单http|`cheetah -i simplehttp`|[项目结构](./docs/http.md)|暂未创建|

### 安装

```bash
apt-get update 
apt-get install curl -y
# linux x86
version=0.0.5
release_package=https://github.com/go-cheetah/cheetah/releases/download/${version}/cheetah-linux-amd64
curl -o cheetah ${release_package}
./cheetah -h
```

## gitee和github仓库地址

> github: https://github.com/go-cheetah/cheetah.git  
> gitee: https://gitee.com/agocan/code-generator.git  