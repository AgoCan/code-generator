package ansible

import (
	_ "embed"
	"fmt"
	"os"
	"path"

	"github.com/agocan/code-generator/config"
	"github.com/agocan/code-generator/generator"
)

var Option generator.Option

var (
	//go:embed templates/ansible/README.md.tmpl
	readmeContent string
	//go:embed templates/ansible/_common.sh.tmpl
	commonContent string
	//go:embed templates/ansible/_gen_inventory_file.sh.tmpl
	genInventoryFileContent string
	//go:embed templates/ansible/all.yaml.tmpl
	allContent string
	//go:embed templates/ansible/config.ini.tmpl
	configIniContent string
	//go:embed templates/ansible/config.sh.tmpl
	configShContent string
	//go:embed templates/ansible/set_config.sh.tmpl
	setConfigShContent string
	//go:embed templates/ansible/start.sh.tmpl
	startShContent string
	//go:embed templates/ansible/tasks.yaml.tmpl
	tasksContent string
)

var files = map[string]string{
	"README.md":                           readmeContent,
	"utils/_common.sh":                    commonContent,
	"utils/_gen_inventory_file.sh":        genInventoryFileContent,
	"playbooks/all.yml":                   allContent,
	"playbooks/roles/base/tasks/main.yml": tasksContent,
	"start.sh":                            startShContent,
	"config/config.ini":                   configIniContent,
	"config/config.sh":                    configShContent,
	"config/set_config.sh":                setConfigShContent,
}

func Run() {
	Option = generator.Option{
		AbsProjectPath: config.AbsProjectPath,
		Title:          *config.Title,
	}
	Option.AbsProjectPath = path.Join(*config.ProjectPath, *config.Title)
	var dirGen generator.DirGenerator
	err := dirGen.Run(&Option)
	if err != nil {
		fmt.Printf("create dirs err: %v", err)
		os.Exit(1)
	}
	var fileGen generator.FileGenerator
	fileGen.Files = files
	// 注册
	generator.Register("files", &fileGen)

	generator.RunGenerator(&Option)
}
