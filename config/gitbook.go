package config

import (
	"flag"
	"strings"
)

var (
	Plugins           []string
	TreeBaToken       = flag.String("t", "", "baidu tongji token")
	GaToken           = flag.String("g", "", "google tongji token")
	ExtraPlugins      = flag.String("e", "", `use "," split extra plugins.`)
	Keywords          = flag.String("k", "keywords", "keywords")
	Description       = flag.String("d", "description", "description")
	Author            = flag.String("a", "author", "author")
	SidebarTitle      = flag.String("st", "", "sidebarTitle")
	SidebarLink       = flag.String("sl", "", "sidebarLink")
	NpmInstallPlugins []string
)

func NewGitbook() {
	// "ga" , "3-ba"
	Plugins = []string{"-highlight", "toggle-chapters", "codeblock-filename", "sectionx", "splitter", "-search",
		"-lunr", "search-pro", "theme-default", "prism", "prism-themes", "theme-comscore",
		"include", "favicon", "anchors", "tbfed-pagefooter", "hide-element"}
	ExtraPluginsSlice := strings.Split(*ExtraPlugins, ",")
	if *ExtraPlugins != "" {
		Plugins = append(Plugins, ExtraPluginsSlice...)
	}
	NpmInstallPlugins = make([]string, len(Plugins))
	for i, v := range Plugins {
		NpmInstallPlugins[i] = strings.TrimLeft(v, "-")
	}
}
