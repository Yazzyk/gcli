package prjectName

import (
	"github.com/manifoldco/promptui"
)

func Run() (result string, err error) {
	prompt := promptui.Prompt{
		Label: "项目名称(Project Name)",
	}
	result, err = prompt.Run()

	return
}
