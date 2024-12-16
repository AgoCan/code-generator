package grpc

import "embed"

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

var Files = map[string]string{
	"api/helloworld.proto":                readFileGetString("templates/api/helloworld.proto.tmpl"),
	"cmd/server/options/options.go":       readFileGetString("templates/cmd/server/options/options.go.tmpl"),
	"cmd/server/server.go":                readFileGetString("templates/cmd/server/server.go.tmpl"),
	"cmd/main.go":                         readFileGetString("templates/cmd/main.go.tmpl"),
	"conf/conf.yaml":                      readFileGetString("templates/conf/conf.yaml.tmpl"),
	"internal/conf/common.go":             readFileGetString("templates/internal/conf/common.go.tmpl"),
	"internal/conf/conf.go":               readFileGetString("templates/internal/conf/conf.go.tmpl"),
	"internal/data/data.go":               readFileGetString("templates/internal/data/data.go.tmpl"),
	"internal/service/helloworld.go":      readFileGetString("templates/internal/service/helloworld.go.tmpl"),
	"internal/server.go":                  readFileGetString("templates/internal/server.go.tmpl"),
	"pb/helloworld/helloworld_grpc.pb.go": readFileGetString("templates/pb/helloworld/helloworld_grpc.pb.go.tmpl"),
	"pb/helloworld/helloworld.pb.go":      readFileGetString("templates/pb/helloworld/helloworld.pb.go.tmpl"),
	"README.md":                           readFileGetString("templates/README.md.tmpl"),
	"go.mod":                              readFileGetString("templates/go.mod.tmpl"),
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

func readFileGetString(fileName string) string {
	data, err := fEmbed.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

var Extra = map[string]interface{}{}
