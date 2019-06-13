package main

import (
	"GoLecture/go_crontab_lecture/master"
	"flag"
	"fmt"
	"runtime"
	"time"
)

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func main()  {

	var (
		err error
	)
	//初始化线程
	initEnv()

	initArgs()

	//
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	if err = master.InitJobMgr(); err!=nil {
		goto ERR
	}

	if err = master.InitApiServer(); err != nil {
		goto ERR

	}

	for {
		time.Sleep(1* time.Second)
	}
	return

	ERR:
		fmt.Println(err)
}