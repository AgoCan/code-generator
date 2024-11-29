package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/go-cheetah/cheetach/config"
	"github.com/go-cheetah/cheetach/generator"
	"github.com/go-cheetah/cheetach/generator/ansible"
	"github.com/go-cheetah/cheetach/generator/command"
	"github.com/go-cheetah/cheetach/generator/gitbook"
	"github.com/go-cheetah/cheetach/generator/http"
	"github.com/go-cheetah/cheetach/generator/mdbook"
	"github.com/go-cheetah/cheetach/generator/mvc"
	"github.com/go-cheetah/cheetach/generator/simple"
)

var Version = "0.0.5"

func run(files map[string]string, dirs []string, extra map[string]interface{}) {

	opt := generator.Option{
		AbsProjectPath: config.AbsProjectPath,
		Title:          *config.Title,
		Extra:          extra,
	}
	opt.AbsProjectPath = path.Join(*config.ProjectPath, *config.Title)
	var dirGen generator.DirGenerator
	dirGen.Dirs = dirs
	var fileGen generator.FileGenerator
	fileGen.Files = files
	// 注册，按顺序来进行添加，先创建目录，再创建文件
	generator.Register(&dirGen)
	generator.Register(&fileGen)
	generator.RunGenerator(&opt)
}

func version() {
	fmt.Println(Version)
}

func main() {
	config.DefaultConfig()
	if *config.Version {
		version()
		return
	}
	generator.Init()
	switch {
	case *config.Item == "ansible":
		run(ansible.Files, ansible.Dirs, ansible.Extra)
	case *config.Item == "gitbook":
		config.NewGitbook()
		run(gitbook.Files, gitbook.Dirs, gitbook.GetExtra())
	case *config.Item == "mdbook":
		run(mdbook.Files, mdbook.Dirs, mdbook.Extra)
	case *config.Item == "simple":
		run(simple.Files, simple.Dirs, simple.Extra)
	case *config.Item == "http":
		run(http.Files, http.Dirs, http.Extra)
	case *config.Item == "mvcgorm":
		run(mvc.GormFiles, mvc.GormDirs, mvc.GormExtra)
	case *config.Item == "command":
		run(command.Files, command.Dirs, command.Extra)
	case *config.Item == "":
		flag.PrintDefaults()
	default:
		fmt.Printf("还不支持%v生成器\n", *config.Item)
	}
}
