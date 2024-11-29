package gitbook

import (
	_ "embed"

	"github.com/go-cheetah/cheetach/config"
)

var (
	//go:embed templates/book.json.tmpl
	bookJsonContent string
	//go:embed templates/Dockerfile.tmpl
	dockerfileContent string
	//go:embed templates/SUMMARY.md.tmpl
	summaryContent string
	//go:embed templates/README.md.tmpl
	readmeContent string
	//go:embed templates/website.css.tmpl
	websiteContent string
)

var Files = map[string]string{
	"book.json":          bookJsonContent,
	"Dockerfile":         dockerfileContent,
	"SUMMARY.md":         summaryContent,
	"README.md":          readmeContent,
	"styles/website.css": websiteContent,
}

var Dirs = []string{
	"styles",
	"assets",
}

func GetExtra() map[string]interface{} {

	return map[string]interface{}{
		"NpmInstallPlugins": config.NpmInstallPlugins,
		"Plugins":           config.Plugins,
		"TreeBaToken":       config.TreeBaToken,
		"GaToken":           config.GaToken,
		"ExtraPlugins":      config.ExtraPlugins,
		"Keywords":          config.Keywords,
		"Description":       config.Description,
		"Author":            config.Author,
		"SidebarTitle":      config.SidebarTitle,
		"SidebarLink":       config.SidebarLink,
	}
}
