package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/go-cheetah/cheetah/config"
	"github.com/go-cheetah/cheetah/generator"
	"github.com/go-cheetah/cheetah/generator/ansible"
	"github.com/go-cheetah/cheetah/generator/command"
	"github.com/go-cheetah/cheetah/generator/gitbook"
	"github.com/go-cheetah/cheetah/generator/grpc-go"
	"github.com/go-cheetah/cheetah/generator/http"
	"github.com/go-cheetah/cheetah/generator/mdbook"
	"github.com/go-cheetah/cheetah/generator/mvc"
	"github.com/go-cheetah/cheetah/generator/simple"
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
		run(ansible.GetFiles(), ansible.Dirs, ansible.Extra)
	case *config.Item == "gitbook":
		config.NewGitbook()
		run(gitbook.Files, gitbook.Dirs, gitbook.GetExtra())
	case *config.Item == "mdbook":
		run(mdbook.Files, mdbook.Dirs, mdbook.Extra)
	case *config.Item == "simple":
		run(simple.Files, simple.Dirs, simple.Extra)
	case *config.Item == "http":
		run(http.Files, http.Dirs, http.Extra)
	case *config.Item == "mvc":
		run(mvc.GetFiles(), mvc.GormDirs, mvc.GormExtra)
	case *config.Item == "command":
		run(command.GetFiles(), command.Dirs, command.Extra)
	case *config.Item == "grpc":
		run(grpc.GetFiles(), grpc.Dirs, grpc.Extra)
	case *config.Item == "":
		flag.PrintDefaults()
	default:
		fmt.Printf("还不支持%v生成器\n", *config.Item)
	}
}
