package main

import (
	"fmt"
	"os"
)

/**
	错误类型：
	文件的错误类型有
	package os
	func IsExist(err error) bool //创建时如果文件存在则
	func IsNotExist(err error) bool //打开文件时如果不存在
	func IsPermission(err error) bool //没有权限打开
 */

func  main() {
	//获取当前目录
	curDir,_ := os.Getwd()
	fmt.Println(curDir)

	//判断不存在
	_,err := os.OpenFile(curDir +"/xuxiao.txt",os.O_RDWR|os.O_APPEND,os.ModePerm)
	fmt.Println(os.IsNotExist(err)) //true---文件不存在

	//创建文件
	file,_ := os.Create(curDir + "/xuxiao.txt")
	file.Close()
	_,err = os.Create(curDir + "/xuxiao.txt")
	fmt.Println(os.IsExist(err)) //false
	os.Remove(curDir + "/xuxiao.txt")

	file,_ = os.OpenFile(curDir + "/xuxiao.txt",os.O_RDONLY,os.ModePerm) //只读
	_,err = file.WriteString("hello")
	fmt.Println(os.IsPermission(err)) //false ---没有权限false



}