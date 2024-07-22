package mvc

// mvc 服务，基于sqlx

import (
	_ "embed"
)

var (
	//go:embed templates/build/docker-file.tmpl
	dockerfileContent string
	//go:embed templates/cmd/app/app/options/options.tmpl
	appOptionsContent string
	//go:embed templates/cmd/app/app/server.tmpl
	appServerContent string
	//go:embed templates/cmd/app/main.tmpl
	mainContent string
	//go:embed templates/config/config.tmpl
	configYamlContent string
	//go:embed templates/internal/config/config.tmpl
	configContent string
	//go:embed templates/internal/config/db.tmpl
	dbConfigContent string
	//go:embed templates/internal/config/log.tmpl
	logConfigContent string
	//go:embed templates/internal/config/server.tmpl
	serverConfigContent string
	//go:embed templates/internal/handler/health/handler.tmpl
	handlerContent string
	//go:embed templates/internal/handler/health/service.tmpl
	serviceContent string
	//go:embed templates/internal/pkg/database/gorm.tmpl
	gormDatabaseContent string
	//go:embed templates/internal/pkg/database/sqlite.tmpl
	sqliteDatabaseContent string
	//go:embed templates/internal/pkg/database/mysql.tmpl
	mysqlDatabaseContent string
	//go:embed templates/internal/pkg/generator/id.tmpl
	idGenerateContent string
	//go:embed templates/internal/pkg/middleware/cors/cors.tmpl
	corsContent string
	//go:embed templates/internal/pkg/middleware/cors/config.tmpl
	corsConfigContent string
	//go:embed templates/internal/pkg/middleware/cors/utils.tmpl
	corsUtilsContent string
	//go:embed templates/internal/pkg/middleware/log/log.tmpl
	middlewarelogContent string
	//go:embed templates/internal/pkg/response/common.tmpl
	responseCommonContent string
	//go:embed templates/internal/pkg/response/response.tmpl
	responseContent string
	//go:embed templates/internal/server/router.tmpl
	routerContent string
	//go:embed templates/internal/server/migrate.tmpl
	migrateContent string
	//go:embed templates/internal/server/server.tmpl
	serverContent string
	//go:embed templates/pkg/pkg.tmpl
	pkgContent string
	//go:embed templates/go.mod.tmpl
	modContent string
	//go:embed templates/gitignore.tmpl
	gitignoreContent string
	//go:embed templates/readme.tmpl
	readmeContent string

	//go:embed templates/internal/model/model.tmpl
	modelContent string
	//go:embed templates/internal/model/health/health.tmpl
	healthContent string
	//go:embed templates/internal/model/health/model.tmpl
	healthModelContent string
	//go:embed templates/internal/model/health/types.tmpl
	healthTypesContent string
)

var GormFiles = map[string]string{
	"cmd/app/main.go":                        mainContent,
	"cmd/app/app/server.go":                  appServerContent,
	"cmd/app/app/options/options.go":         appOptionsContent,
	"config/config.yaml":                     configYamlContent,
	"internal/config/config.go":              configContent,
	"internal/config/db.go":                  dbConfigContent,
	"internal/config/log.go":                 logConfigContent,
	"internal/config/server.go":              serverConfigContent,
	"internal/handler/health/handler.go":     handlerContent,
	"internal/handler/health/service.go":     serviceContent,
	"internal/model/model.go":                modelContent,
	"internal/model/health/health.go":        healthContent,
	"internal/model/health/model.go":         healthModelContent,
	"internal/model/health/types.go":         healthTypesContent,
	"internal/pkg/database/database.go":      gormDatabaseContent,
	"internal/pkg/database/mysql.go":         mysqlDatabaseContent,
	"internal/pkg/database/sqlite.go":        sqliteDatabaseContent,
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
	"internal/model/health",
	"internal/pkg/database",
	"internal/pkg/generator",
	"internal/pkg/middleware/cors",
	"internal/pkg/middleware/log",
	"internal/pkg/response",
	"internal/server",
	// "internal/router",
	"pkg",
}

var GormExtra = map[string]interface{}{}
