package config

import (
	"strings"
)

// 如果变量都是通过flag获取的话，可以直接使用 ansible下面的Extra
// 如果跟这里一样的话，有非flag的，则需要写成函数，在初始化后进行重新赋值
var (
	Plugins           []string
	TreeBaToken       = ""            // flag.String("baidu", "", "gitbook: baidu tongji token")
	GaToken           = ""            // flag.String("google", "", "gitbook: google tongji token")
	ExtraPlugins      = ""            // flag.String("extra-plugin", "", `gitbook: use "," split extra plugins.`)
	Keywords          = "keywords"    // flag.String("keywords", "keywords", "gitbook: keywords")
	Description       = "description" // flag.String("description", "description", "gitbook: description")
	Author            = "author"      // flag.String("author", "author", "gitbook: author")
	SidebarTitle      = ""            // flag.String("sidebar-title", "", "gitbook: sidebarTitle")
	SidebarLink       = ""            // flag.String("sidebar-link", "", "gitbook: sidebarLink")
	NpmInstallPlugins []string
)

func NewGitbook() {
	// "ga" , "3-ba"
	Plugins = []string{"-highlight", "toggle-chapters", "codeblock-filename", "splitter", "-search",
		"-lunr", "search-pro", "theme-default", "prism", "prism-themes", "theme-comscore",
		"include", "favicon", "anchors", "tbfed-pagefooter", "hide-element"}
	ExtraPluginsSlice := strings.Split(ExtraPlugins, ",")
	if ExtraPlugins != "" {
		Plugins = append(Plugins, ExtraPluginsSlice...)
	}
	NpmInstallPlugins = make([]string, len(Plugins))
	for i, v := range Plugins {
		NpmInstallPlugins[i] = strings.TrimLeft(v, "-")
	}
}
