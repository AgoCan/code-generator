package ansible

import (
	_ "embed"
)

var (
	//go:embed templates/README.md.tmpl
	readmeContent string
	//go:embed templates/_common.sh.tmpl
	commonContent string
	//go:embed templates/_gen_inventory_file.sh.tmpl
	genInventoryFileContent string
	//go:embed templates/all.yaml.tmpl
	allContent string
	//go:embed templates/config.ini.tmpl
	configIniContent string
	//go:embed templates/config.sh.tmpl
	configShContent string
	//go:embed templates/set_config.sh.tmpl
	setConfigShContent string
	//go:embed templates/start.sh.tmpl
	startShContent string
	//go:embed templates/tasks.yaml.tmpl
	tasksContent string
)

var Files = map[string]string{
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

var Dirs = []string{
	"config",
	"playbooks",
	"utils",
	"playbooks/roles/base/tasks",
}

var Extra = map[string]interface{}{}
