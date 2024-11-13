package config

import (
	"flag"
)

// 直接写配置成配置文件，不多使用额外的配置文件，例如不使用config.yaml

var (
	Title       = flag.String("n", "demo", "Common: Project name.")
	ProjectPath = flag.String("p", ".", "Common: Project path.")
	Version     = flag.Bool("v", false, "Print version.")
	Item        = flag.String("i",
		"",
		"Common: Project item. Like 'ansible'、'gitbook'、'mdbook'、'simple'、'simplecobra'、'simplehttp'、'mvcgorm'.",
	)
	AbsProjectPath string
)

func DefaultConfig() {
	flag.Parse()
}
