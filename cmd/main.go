package main

import (
	"fmt"
	"github.com/yazzyk/gcli/internal/prjectName"
	"github.com/yazzyk/gcli/internal/useHttp"
)

func main() {
	projectName, err := prjectName.Run()
	if err != nil {
		fmt.Printf("请输入项目名称 %v\n", err)
		return
	}
	fmt.Printf("项目名称: %q\n", projectName)
	useHTTP, err := useHttp.Run()
	if err != nil {
		fmt.Println("请输入y或n" + err.Error())
		return
	}
	fmt.Printf("http: %v \n", useHTTP)

}
