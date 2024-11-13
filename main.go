package main

import (
	"fmt"
	"path"

	"github.com/agocan/code-generator/config"
	"github.com/agocan/code-generator/generator"
	"github.com/agocan/code-generator/generator/ansible"
	"github.com/agocan/code-generator/generator/command"
	"github.com/agocan/code-generator/generator/gitbook"
	"github.com/agocan/code-generator/generator/mdbook"
	"github.com/agocan/code-generator/generator/mvc"
	"github.com/agocan/code-generator/generator/simple"
	"github.com/agocan/code-generator/generator/simplehttp"
)

var Version = "0.0.3"

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
	case *config.Item == "simplehttp":
		run(simplehttp.Files, simplehttp.Dirs, simplehttp.Extra)
	case *config.Item == "mvcgorm":
		run(mvc.GormFiles, mvc.GormDirs, mvc.GormExtra)
	case *config.Item == "command":
		run(command.Files, command.Dirs, command.Extra)
	default:
		fmt.Printf("还不支持%v生成器\n", *config.Item)
	}
}
