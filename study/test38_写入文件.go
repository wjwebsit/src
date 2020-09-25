package main

import (
	"io/ioutil"
	"os"
)

func main() {
	//获取当前目录
	curDir,_ := os.Getwd()
	fileName := curDir + "/src/write.go"

	//创建文件--需要主要的是如果存在会删掉新建一个空的文件
	os.Create(curDir + "/src/write.go")

	//只读方式打开文件方式--如果不存在会出错
	file,_:= os.Open(fileName)
	file.Close()

	//读写方式打开---默认为覆盖写而且只是覆盖当前写入的长度
	file,_ = os.OpenFile(fileName,os.O_RDWR,os.ModePerm)
	file.WriteString("package zyb \r\n")
	file.Write([]byte{'h','e','l','l','o'})
	file.Close() //文件关闭以后就开始覆盖了--文件指针从0开始了

	//追加写入貌似只有这一种翻案
	file,_ = os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	file.WriteString("fff")

	//清空文件
	file.Truncate(0)

	//#####应为覆盖只是覆盖当前未关闭的文件的内容故覆盖写之前先清空一下

	/***利用ioutil***/
	//覆盖写
	ioutil.WriteFile(fileName,[]byte("sddff"),os.ModePerm)
	ioutil.WriteFile(fileName,[]byte("s"),os.ModePerm)
}
