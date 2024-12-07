package command

import "embed"

//go:embed templates/build/build.sh.tmpl
//go:embed templates/cmd/app/app/options/options.go.tmpl
//go:embed templates/cmd/app/app/server.go.tmpl
//go:embed templates/cmd/app/main.go.tmpl
//go:embed templates/config/config.yaml.tmpl
//go:embed templates/docs/README.md.tmpl
//go:embed templates/internal/command/command.go.tmpl
//go:embed templates/internal/config/common.go.tmpl
//go:embed templates/internal/config/config.go.tmpl
//go:embed templates/internal/log/log.go.tmpl
//go:embed templates/README.md.tmpl
//go:embed templates/go.mod.tmpl
var fEmbed embed.FS

var Files = map[string]string{
	"build/build.sh":                 readFileGetString("templates/build/build.sh.tmpl"),
	"cmd/app/app/options/options.go": readFileGetString("templates/cmd/app/app/options/options.go.tmpl"),
	"cmd/app/app/server.go":          readFileGetString("templates/cmd/app/app/server.go.tmpl"),
	"cmd/app/main.go":                readFileGetString("templates/cmd/app/main.go.tmpl"),
	"config/config.yaml":             readFileGetString("templates/config/config.yaml.tmpl"),
	"docs/README.md":                 readFileGetString("templates/docs/README.md.tmpl"),
	"internal/command/command.go":    readFileGetString("templates/internal/command/command.go.tmpl"),
	"internal/config/config.go":      readFileGetString("templates/internal/config/config.go.tmpl"),
	"internal/config/common.go":      readFileGetString("templates/internal/config/common.go.tmpl"),
	"internal/log/log.go":            readFileGetString("templates/internal/log/log.go.tmpl"),
	"README.md":                      readFileGetString("templates/README.md.tmpl"),
	"go.mod":                         readFileGetString("templates/go.mod.tmpl"),
}

var Dirs = []string{
	"build",
	"cmd/app/app/options",
	"config",
	"docs",
	"internal/config",
	"internal/command",
	"internal/log",
}

func readFileGetString(fileName string) string {
	data, err := fEmbed.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

var Extra = map[string]interface{}{}
