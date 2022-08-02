package simplecobra

// 简单的应用，并加上cobra做命令行传参

import (
	_ "embed"
)

var (
	//go:embed templates/app.go.tmpl
	appContent string
	//go:embed templates/config.go.tmpl
	configContent string
	//go:embed templates/config.yaml.tmpl
	configYamlContent string
	//go:embed templates/go.mod.tmpl
	modContent string
	//go:embed templates/log.go.tmpl
	logContent string
	//go:embed templates/main.go.tmpl
	mainContent string
	//go:embed templates/cmdroot.go.tmpl
	cmdRootContent string
	//go:embed templates/cmdversion.go.tmpl
	cmdVersionContent string
)

var Files = map[string]string{
	"app/app.go":             appContent,
	"app/config.go":          configContent,
	"go.mod":                 modContent,
	"config/config.yaml":     configYamlContent,
	"app/log.go":             logContent,
	"main.go":                mainContent,
	"cmd/root.go":            cmdRootContent,
	"cmd/version/version.go": cmdVersionContent,
}

var Dirs = []string{
	"app",
	"config",
	"cmd",
	"cmd/version",
}

var Extra = map[string]interface{}{}
