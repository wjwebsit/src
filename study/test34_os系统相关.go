package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(1)
	//获取主机名
	fmt.Println(os.Hostname()) //xupengxiangdeMacBook-Pro.local,nil

	//获取用户ID
	fmt.Println(os.Getuid()) //501

	//获取有效的用户ID
	fmt.Println(os.Geteuid()) //501

	//获取组ID
	fmt.Println(os.Getgid()) //20

	//获取有效的组ID
	fmt.Println(os.Getegid()) // 20

	//获取所有的组
	fmt.Println(os.Getgroups())

	//获取环境变量的值
	fmt.Println(os.Getenv("GOPATH")) ///Users/xupengxiang/go

	//设置环境变量变量的值
	fmt.Println(os.Setenv("TEST","1"))

	//获取所有的环境变量
	fmt.Println(strings.Join(os.Environ(),"\r\n"))
	//os.Clearenv(); ---删除当前的go系统的继承的环境变量

	//结束进程==1秒后
	os.Exit(1)

}



