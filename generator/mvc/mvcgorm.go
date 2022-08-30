package mvc

// mvc 服务，基于sqlx

import (
	_ "embed"
)

var (
	//go:embed templates/internal/pkg/database/gorm.tmpl
	gormDatabaseContent string
)

var GormFiles = map[string]string{
	"cmd/app/main.go":                        mainContent,
	"cmd/app/app/server.go":                  appServerContent,
	"cmd/app/app/options/options.go":         appOptionsContent,
	"config/config.yaml":                     configYamlContent,
	"internal/config/config.go":              configContent,
	"internal/handler/health/handler.go":     handlerContent,
	"internal/handler/health/service.go":     serviceContent,
	"internal/pkg/database/database.go":      gormDatabaseContent,
	"internal/pkg/generator/id.go":           idGenerateContent,
	"internal/pkg/middleware/cors/cors.go":   corsContent,
	"internal/pkg/middleware/cors/config.go": corsConfigContent,
	"internal/pkg/middleware/cors/utils.go":  corsUtilsContent,
	"internal/pkg/middleware/log/log.go":     middlewarelogContent,
	"internal/pkg/response/common.go":        responseCommonContent,
	"internal/pkg/response/response.go":      responseContent,
	"internal/server/server.go":              serverContent,
	"internal/server/router.go":              routerContent,
	"internal/server/migrate.go":             migrateContent,
	"pkg/pkg.go":                             pkgContent,
	"README.md":                              readmeContent,
	"go.mod":                                 modContent,
	".gitignore":                             gitignoreContent,
	"build/Dockerfile":                       dockerfileContent,
}

var GormDirs = []string{
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
	"internal/server",
	"internal/router",
	"pkg",
}

var GormExtra = map[string]interface{}{}
