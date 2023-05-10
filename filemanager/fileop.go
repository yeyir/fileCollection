package filemanager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gen2brain/go-unarr"
)

type Pkgdir struct {
	Path_abs string
	Content  map[string]bool
}

func GetZipFile(src string, target string) (string, error) {
	fileinfo, err0 := os.Stat(src)
	if err0 != nil {
		panic(err0)
	}
	filename := strings.Split(fileinfo.Name(), ".")
	a, err := unarr.NewArchive(src)
	if err != nil {
		panic(err)
	}
	defer a.Close()
	// fmt.Println(a.List())
	now := time.Now()
	timeTag := now.Format("20060102150405")
	_, err2 := a.Extract(target + "/" + filename[0] + "-" + timeTag + "/")
	fmt.Println("解压目录:" + target + "/" + filename[0] + "-" + timeTag + "/")
	if err != nil {
		panic(err2)
	}
	// fmt.Println(c)
	return "ok", nil
}
func GetFileList(src string) Pkgdir {
	// fmt.Println("获取当前路径的文件")
	pp, err0 := filepath.Abs(src)
	if err0 != nil {
		panic(err0)
	}
	var pkgdir Pkgdir
	pkgdir.Path_abs = pp
	// fmt.Println(pp)
	fs, err := ioutil.ReadDir(src)
	if err != nil {
		panic(err)
	}
	cc := make(map[string]bool)
	pkgdir.Content = cc
	for _, info := range fs {
		// fmt.Println(info.IsDir(), info.Name())
		cc[info.Name()] = info.IsDir()
	}
	// fmt.Println(pkgdir)
	return pkgdir
}
