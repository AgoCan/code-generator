package main

import (
	"fmt"
	"os"
	"path"

	"github.com/agocan/code-generator/config"
	"github.com/agocan/code-generator/generator"
	"github.com/agocan/code-generator/generator/ansible"
)

func run(files map[string]string, dirs []string) {

	opt := generator.Option{
		AbsProjectPath: config.AbsProjectPath,
		Title:          *config.Title,
		Dirs:           dirs,
	}
	opt.AbsProjectPath = path.Join(*config.ProjectPath, *config.Title)
	var dirGen generator.DirGenerator
	err := dirGen.Run(&opt)
	if err != nil {
		fmt.Printf("create dirs err: %v", err)
		os.Exit(1)
	}
	var fileGen generator.FileGenerator
	fileGen.Files = files
	// 注册
	generator.Register("files", &fileGen)
	generator.RunGenerator(&opt)
}

func main() {
	config.DefaultConfig()
	generator.Init()
	if *config.Item == "ansible" {
		run(ansible.Files, ansible.Dirs)
	} else {
		fmt.Printf("还不支持%v生成器\n", *config.Item)
	}
}
