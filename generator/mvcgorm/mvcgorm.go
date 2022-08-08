package mvcgorm

// mvc服务，基于gorm

import (
	_ "embed"
)

var (
	//go:embed templates/config.tmpl
	configContent string
	//go:embed templates/configoption.tmpl
	configOptionContent string
	//go:embed templates/configyaml.tmpl
	configYamlContent string
	//go:embed templates/docker-file.tmpl
	dockerfileContent string
	//go:embed templates/idgenerator.tmpl
	idGenerateContent string
	//go:embed templates/main.tmpl
	mainContent string
	//go:embed templates/middlewarelog.tmpl
	middlewarelogContent string
	//go:embed templates/mod.tmpl
	modContent string
	//go:embed templates/model.tmpl
	modelContent string
	//go:embed templates/readme.tmpl
	readmeContent string
	//go:embed templates/router.tmpl
	routerContent string
)

var Files = map[string]string{
	"main.go":               mainContent,
	"routers/router.go":     routerContent,
	"middleware/log.go":     middlewarelogContent,
	"config/config.go":      configContent,
	"config/config.yaml":    configYamlContent,
	"models/model.go":       modelContent,
	"Dockerfile":            dockerfileContent,
	"README.md":             readmeContent,
	"config/option.go":      configOptionContent,
	"utils/generator/id.go": idGenerateContent,
	"go.mod":                modContent,
}

var Dirs = []string{
	"handlers",
	"routers",
	"models",
	"templates",
	"utils/response",
	"utils/generator",
	"config",
	"log",
	"middleware",
}

var Extra = map[string]interface{}{}
