package mvc

// mvc 服务，基于sqlx

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
	//go:embed templates/cors.tmpl
	corsContent string
	//go:embed templates/docker-file.tmpl
	dockerfileContent string
	//go:embed templates/gitignore.tmpl
	gitignoreContent string
	//go:embed templates/idgenerator.tmpl
	idGenerateContent string
	//go:embed templates/main.tmpl
	mainContent string
	//go:embed templates/middlewarelog.tmpl
	middlewarelogContent string
	//go:embed templates/mod.tmpl
	modContent string
	//go:embed templates/readme.tmpl
	readmeContent string
	//go:embed templates/router.tmpl
	routerContent string
	//go:embed templates/response.tmpl
	responseContent string
	//go:embed templates/responsecommon.tmpl
	responseCommonContent string
	//go:embed templates/responseping.tmpl
	responsePingContent string
	//go:embed templates/appoptions.tmpl
	appOptionsContent string
	//go:embed templates/appserver.tmpl
	appServerContent string
	//go:embed templates/database.tmpl
	databaseContent string
	//go:embed templates/handler.tmpl
	handlerContent string
	//go:embed templates/ping.tmpl
	pingContent string
	//go:embed templates/service.tmpl
	serviceContent string
	//go:embed templates/pkg.tmpl
	pkgContent string
)

var Files = map[string]string{
	"cmd/app/main.go":                      mainContent,
	"cmd/app/app/server.go":                appServerContent,
	"cmd/app/app/options/options.go":       appOptionsContent,
	"config/config.yaml":                   configYamlContent,
	"internal/config/config.go":            configContent,
	"internal/config/option.go":            configOptionContent,
	"internal/handler/health/handler.go":   handlerContent,
	"internal/handler/health/service.go":   serviceContent,
	"internal/pkg/database/database.go":    databaseContent,
	"internal/pkg/database/ping.go":        pingContent,
	"internal/pkg/generator/id.go":         idGenerateContent,
	"internal/pkg/middleware/cors/cors.go": corsContent,
	"internal/pkg/middleware/log/log.go":   middlewarelogContent,
	"internal/pkg/response/common.go":      responseCommonContent,
	"internal/pkg/response/ping.go":        responsePingContent,
	"internal/pkg/response/response.go":    responseContent,
	"internal/router/router.go":            routerContent,
	"pkg/pkg.go":                           pkgContent,
	"README.md":                            readmeContent,
	"go.mod":                               modContent,
	".gitignore":                           gitignoreContent,
	"build/Dockerfile":                     dockerfileContent,
}

var Dirs = []string{
	"api",
	"build",
	"cmd/app/app/options",
	"config",
	"docs",
	"internal/config",
	"internal/handler/health",
	"internal/pkg/database",
	"internal/pkg/generator",
	"internal/pkg/middleware/cors",
	"internal/pkg/middleware/log",
	"internal/pkg/response",
	"internal/router",
	"pkg",
}

var Extra = map[string]interface{}{}
