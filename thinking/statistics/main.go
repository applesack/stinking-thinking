package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	conf := handleCommand()
	search(conf)
}

func handleCommand() *config {
	var (
		conf config
		lang string
	)
	flag.StringVar(&conf.path, "path", ".", "要搜索的路径")
	flag.StringVar(&lang, "lang", "",
		"选择要搜索的文件类型(后缀),如果有多个类型要搜索,多个类型之间用逗号隔开")
	flag.Parse()
	if len(lang) != 0 {
		conf.lang = strings.Split(lang, ",")
		if len(conf.lang) == 0 {
			fmt.Println("未选择要搜索的文件类型")
			os.Exit(0)
		}
	}
	return &conf
}

func search(conf *config) {
	file, err := os.Open(conf.path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(fmt.Sprintf("路径不存在%s", conf.path))
			os.Exit(0)
		} else {
			fmt.Println(fmt.Sprintf("路径无法访问%s, 原因%s", conf.path, err))
		}
	}

	ctrl := newPCB(conf)
	go handleFiles(ctrl)
	go recursive(file, ctrl)
}

func handleFiles(ctrl *pcb) {
	select {
	case f := <-ctrl.chanFile:
		go handleFile(f, ctrl)
	case <-ctrl.chanFileClose:
		close(ctrl.chanFile)
		close(ctrl.chanFileClose)
	}
}

func handleFile(f fb, ctrl *pcb) {
	fName := f.file.Name()
	for _, suffix := range ctrl.conf.lang {
		if strings.HasSuffix(fName, suffix) {
			return
		}
	}
}

func recursive(file *os.File, ctrl *pcb) {
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("读取文件错误, 路径")
		return
	}
	if stat.IsDir() {
		parseDir(file, ctrl)
		return
	}

	ctrl.chanFile <- newFB("", file)
}

func parseDir(file *os.File, ctrl *pcb) {
	dir, err := file.Readdirnames(0)
	if err != nil {
		return
	}

	for _, member := range dir {
		child, err := os.Open(member)
		if err != nil {
			continue
		}
		go recursive(child, ctrl)
	}
}
