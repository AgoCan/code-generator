package grpc

import (
	"embed"

	"github.com/go-cheetah/cheetah/generator/gen"
)

//go:embed templates/api/helloworld.proto.tmpl
//go:embed templates/cmd/server/options/options.go.tmpl
//go:embed templates/cmd/server/server.go.tmpl
//go:embed templates/cmd/main.go.tmpl
//go:embed templates/conf/conf.yaml.tmpl
//go:embed templates/internal/conf/common.go.tmpl
//go:embed templates/internal/conf/conf.go.tmpl
//go:embed templates/internal/data/data.go.tmpl
//go:embed templates/internal/service/helloworld.go.tmpl
//go:embed templates/internal/server.go.tmpl
//go:embed templates/pb/helloworld/helloworld_grpc.pb.go.tmpl
//go:embed templates/pb/helloworld/helloworld.pb.go.tmpl
//go:embed templates/README.md.tmpl
//go:embed templates/go.mod.tmpl
var fEmbed embed.FS

func GetFiles() (files map[string]string) {
	return gen.GetFiles(fEmbed)
}

var Dirs = []string{
	"api",
	"cmd/server/options",
	"conf",
	"docs",
	"internal/conf",
	"internal/data",
	"internal/service",
	"pb/helloworld",
}

var Extra = map[string]interface{}{}
