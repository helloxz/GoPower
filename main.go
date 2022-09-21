package main

import (
	"fmt"
	"http_shutdown/cli"
	"http_shutdown/controller"
	"http_shutdown/router"
	"os"
	"time"
)

func main() {
	//获取命令行参数
	args := os.Args
	//获取切片长度
	args_cap := cap(args)

	//如果没有任何参数，则直接运行
	if args_cap == 1 {
		//获取当前时间
		now := time.Now()
		//当前时间传递
		controller.GetStartTime(now)
		router.Start()
	} else if args_cap == 2 {
		//如果参数长度为2，根据长度决定后续执行步骤

		//获取版本信息
		if args[1] == "version" {
			cli.Get_Version()
			os.Exit(0)
		} else if args[1] == "start" {
			//运行程序

			//获取当前时间
			now := time.Now()
			//当前时间传递
			controller.GetStartTime(now)
			router.Start()
		} else {
			fmt.Printf("Parameter is incorrect!\n")
		}
	}

}
