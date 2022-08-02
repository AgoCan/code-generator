package simplehttp

// 简单的http服务，例如做个metrics的任务
import (
	_ "embed"
)

var (
	//go:embed templates/go.mod.tmpl
	modContent string
	//go:embed templates/main.go.tmpl
	mainContent string
)

var Files = map[string]string{
	"main.go": mainContent,
	"go.mod":  modContent,
}

var Dirs = []string{}

var Extra = map[string]interface{}{}
