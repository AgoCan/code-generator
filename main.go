package main

import (
	"fmt"
	"path"

	"github.com/agocan/code-generator/config"
	"github.com/agocan/code-generator/generator"
	"github.com/agocan/code-generator/generator/ansible"
	"github.com/agocan/code-generator/generator/gitbook"
)

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

func main() {
	config.DefaultConfig()
	generator.Init()
	if *config.Item == "ansible" {
		run(ansible.Files, ansible.Dirs, ansible.Extra)
	} else if *config.Item == "gitbook" {
		config.NewGitbook()
		run(gitbook.Files, gitbook.Dirs, gitbook.GetExtra())
	} else {
		fmt.Printf("还不支持%v生成器\n", *config.Item)
	}
}
