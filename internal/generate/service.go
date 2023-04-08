package generate

import (
	"fmt"
	"os"
)

type Generate struct {
	ProjectName string
	UseHTTP     bool
}

func (this *Generate) Run() {
	if err := this.folder(); err != nil {
		fmt.Printf(err.Error())
		return
	}

}

func (this *Generate) folder() (err error) {

	if err = os.Mkdir(this.ProjectName, os.ModePerm); err != nil {
		fmt.Println("创建项目目录错误:" + err.Error())
		return
	}

	return
}

func (this *Generate) cloneTemplate() (err error) {

	return
}
