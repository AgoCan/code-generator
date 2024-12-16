# grpc-go

## 代码目录结构

> 三方库，根据习惯使用pkb目录，或者直接添加目录

```
.
├── api                                # 放proto文件，方便下载
│   └── helloworld.proto
├── cmd                                # 入口函数
│   ├── main.go                        # 主函数
│   └── server                         # 实际的入口内容
│       ├── options                    # 处理参数
│       │   └── options.go
│       └── server.go
├── conf                               # 配置文件
│   └── config.yaml
├── go.mod                             # go mod
├── internal                           # 内部函数
│   ├── conf                           # 处理配置函数
│   │   ├── common.go
│   │   └── conf.go
│   ├── data                           # 处理数据
│   │   └── data.go
│   ├── server.go                      # 启动grpc
│   └── service                        # 服务的处理
│       └── helloworld.go
├── pb                                 # 根据proto生成的文件，放置外层方便其他服务进行调用
│   └── helloworld                     # demo
│       ├── helloworld_grpc.pb.go
│       └── helloworld.pb.go
└── README.md                          # readme
```

## 补充

### grpc

安装grpc命令

```bash
package=protoc-29.0-linux-x86_64.zip
wget https://github.com/protocolbuffers/protobuf/releases/download/v29.0/${package}
mkdir protoc
unzip ${package} -d protoc
mv protoc/bin/protoc /usr/local/bin/protoc
```

### 准备好go环境

```bash
go_version=1.23.3
wget https://golang.google.cn/dl/go${go_version}.linux-amd64.tar.gz
tar xf go${go_version}.linux-amd64.tar.gz -C /usr/local
rm -f go${go_version}.linux-amd64.tar.gz

echo export GOPROXY=https://goproxy.cn >> /etc/profile
echo export GO111MODULE=on >> /etc/profile
echo export GOPATH=/root/go >> /etc/profile
echo export PATH=\$PATH:/usr/local/go/bin >> /etc/profile
echo export PATH=\$PATH:\$\(go env GOPATH\)/bin >> /etc/profile

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```
