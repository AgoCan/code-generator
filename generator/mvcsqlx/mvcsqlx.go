package mvcsqlx

// mvc 服务，基于sqlx

import (
	_ "embed"
)

var (
	//go:embed templates/healthapi.tmpl
	apiContent string
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
	//go:embed templates/modelping.tmpl
	modelPingContent string
	//go:embed templates/models.tmpl
	modelContent string
	//go:embed templates/readme.tmpl
	readmeContent string
	//go:embed templates/router.tmpl
	routerContent string
	//go:embed templates/healthservice.tmpl
	serviceContent string
	//go:embed templates/response.tmpl
	responseContent string
	//go:embed templates/responsecommon.tmpl
	responseCommonContent string
	//go:embed templates/responseping.tmpl
	responsePingContent string
)

var Files = map[string]string{
	"main.go":                  mainContent,
	"routers/router.go":        routerContent,
	"api/health.go":            apiContent,
	"middleware/log/log.go":    middlewarelogContent,
	"middleware/cors/cors.go":  corsContent,
	"config/config.go":         configContent,
	"config/config.yaml":       configYamlContent,
	"service/health/health.go": serviceContent,
	"models/model.go":          modelContent,
	"models/ping.go":           modelPingContent,
	"go.mod":                   modContent,
	"Dockerfile":               dockerfileContent,
	"README.md":                readmeContent,
	"config/option.go":         configOptionContent,
	"utils/generator/id.go":    idGenerateContent,
	"response/response.go":     responseContent,
	"response/common.go":       responseCommonContent,
	"response/ping.go":         responsePingContent,
	".gitignore":               gitignoreContent,
}

var Dirs = []string{
	"api",
	"service/health",
	"routers",
	"models",
	"templates",
	"response",
	"utils/generator",
	"config",
	"log",
	"middleware/log",
	"middleware/cors",
}

var Extra = map[string]interface{}{}
