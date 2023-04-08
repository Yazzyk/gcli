package useHttp

import (
	"errors"
	"github.com/manifoldco/promptui"
	"strings"
)

func Run() (useHttp bool, err error) {
	prompt := promptui.Prompt{
		Label:    "是否使用HTTP服务?(y/n)",
		Validate: validate,
	}
	result, err := prompt.Run()
	return strings.ToLower(result) == "y", err
}

func validate(input string) error {
	if strings.ToLower(input) != "y" && strings.ToLower(input) != "n" {
		return errors.New("请输入y或n")
	}
	return nil
}
