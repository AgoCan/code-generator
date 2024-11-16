
# 代码脚手架

一键生成一个代码目录架构。

## 工具使用

使用`./code-generator --help`查看帮助

|功能|命令|介绍|demo|
|---|---|---|---|
|【golang】基于gin的mvc框架|`code-generator -i mvcgorm`|[项目结构](./docs/mvc.md)|[demo地址](https://github.com/AgoCan/mvc-demo)|
|【ansible】基于shell|`code-generator -i ansible`|[项目结构](./docs/ansible.md)|暂未创建|
|【golang】命令行方式|`code-generator -i command`|[项目结构](./docs/command.md)|暂未创建|
|【gitbook】文档框架|`code-generator -i gitbook`|[项目结构](./docs/gitbook.md)|暂未创建|
|【mdbook】文档框架|`code-generator -i mdbook`|[项目结构](./docs/mdbook.md)|暂未创建|
|【golang】简单函数|`code-generator -i simple`|[项目结构](./docs/simple.md)|暂未创建|
|【golang】简单http|`code-generator -i simplehttp`|[项目结构](./docs/http.md)|暂未创建|

### 安装

```bash
apt-get update 
apt-get install curl -y
# linux x86
version=0.0.4
release_package=https://github.com/AgoCan/code-generator/releases/download/${version}/code-generator-linux-amd64
curl -o code-generator ${release_package}
./code-generator -h
```

## gitee和github仓库地址

> github: https://github.com/AgoCan/code-generator.git  
> gitee: https://gitee.com/agocan/code-generator.git  