package generator

import (
	"fmt"
	"os"
	"text/template"
)

// Option 参数保存
type Option struct {
	// AbsProjectPath 项目路径+项目名称
	AbsProjectPath string
	Title          string
}

// GeneratorMgr 生成器管理
type GeneratorMgr struct {
	genMap map[string]Generator
}

var genMgr *GeneratorMgr

// Generator 生成器接口
type Generator interface {
	// Run 生成器run
	Run(opt *Option) error
}

// init 初始化
func Init() {
	genMgr = &GeneratorMgr{
		genMap: make(map[string]Generator),
	}
}

// Register 把生成器都注册到map中，然后轮询执行
func Register(name string, gen Generator) (err error) {
	_, ok := genMgr.genMap[name]
	if ok {
		err = fmt.Errorf("genrator %v exits", name)
		return err
	}
	genMgr.genMap[name] = gen
	return nil
}

// writeFile 使用模版文件直接写入文件
func writeFile(tmpl, filePath string, opt *Option) (err error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	// t, err := template.ParseFiles(tmpl)
	t, err := template.New(filePath).Parse(tmpl)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	err = t.Execute(file, &opt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// RunGenerator 运行所有已经注册的生成器
func RunGenerator(opt *Option) (err error) {
	for _, gen := range genMgr.genMap {

		err = gen.Run(opt)
		if err != nil {
			fmt.Printf("err: %v", err)
			return err
		}
	}
	return nil
}
