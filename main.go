package main

import (
	"filemanager"
	"fmt"
	"os"
	"path/filepath"
)

var ppath string
var next string

func main() {
	// filemanager.GetZipFile("./haha.rar", "./h2")
	fmt.Print("输入初始目录：")
	fmt.Scan(&ppath)
	var err error
	if ppath != "" {
		ppath, err = filepath.Abs(ppath)
		if err != nil {
			panic(err)
		}
		serfDirectory()
	}

}

func serfDirectory() {
	if ppath != "" {
		p := filemanager.GetFileList(ppath)
		fmt.Println("")
		for Name, IsDir := range p.Content {
			if IsDir {

			}
			fmt.Println(Name, IsDir)
		}
	label0:
		fmt.Println("当前目录：" + p.Path_abs)
		fmt.Print("输入(cd进入目录/uz解压文件/exit退出):")
		var cmd string
		fmt.Scanln(&cmd, &next)
		fmt.Println(next)
		switch cmd {
		case "cd":
			cdDirectory(next)
			goto label0
		case "uz":
			fmt.Println("解压文件" + next)
			filemanager.GetZipFile(ppath+"/"+next, "./")
			goto label0
		case "exit":
			os.Exit(1)
			break
		default:
			fmt.Println("请输入正确代码")
			goto label0
		}
		// cdDirectory(next)
	}
}
func cdDirectory(next string) {
	_, err := filepath.Abs(ppath + "/" + next)
	if err != nil {
		panic(err)
	}
	ppath = ppath + "/" + next
	serfDirectory()
}
