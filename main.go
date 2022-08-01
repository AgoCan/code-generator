package main

import (
	"fmt"

	"github.com/agocan/code-generator/config"
	"github.com/agocan/code-generator/generator"
	"github.com/agocan/code-generator/generator/ansible"
)

func main() {
	config.DefaultConfig()
	generator.Init()
	if *config.Item == "ansible" {
		ansible.Run()
	} else {
		fmt.Printf("还不支持%v\n", *config.Item)
	}
	//Option.AbsProjectPath = path.Join(*config.ProjectPath, "")

}
