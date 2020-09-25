package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	//"io/ioutil"
	"os"
)
/**
	逐行读取文件
 */
func line(fileName string) []string  {
	//打开文件
	file,err := os.Open(fileName)

	//判断
	if err != nil {
		return nil
	}

	//关闭文件
	defer file.Close()
	var result []string
	//助航扫描
	scanner := bufio.NewScanner(file)

	//逐行扫描
	for scanner.Scan() {
		result = append(result,scanner.Text())
	}

	//返回
	return result
}

func file_get_content(fileName string) string {
	//读取整个文件
	file,err := os.Open(fileName)
	if err != nil {
		return  ""
	}

	//获取文件大小
	fileInfo,_ := file.Stat()
	if fileInfo.Size() == 0 {
		return ""
	}

	//声明缓冲区
	var buffer bytes.Buffer

	//分段读取避免过长
	ll := 2000
	for {
		bt := make([]byte, ll)
		_,err := file.Read(bt)

		//写入结果
		buffer.Write(bt)
		if err == io.EOF {
			break
		}
	}

	//返回
	return buffer.String()
}

func main() {
	//逐行扫描
	curDir,_ := os.Getwd()
	line(curDir + "/Login.go")

	//读取所有信息1
	fmt.Println(file_get_content(curDir + "/Document.go"))

	//读取所有文件内容2--ioutil.ReadFile(文件名称) --byte数组返回
	f, _ := ioutil.ReadFile(curDir + "/Document.go")
	fmt.Println(string(f))

	//读取所有文件内容3--io.util.ReadAll(io.File)
	file,_ := os.Open(curDir + "/Document.go")
	f ,_ = ioutil.ReadAll(file)
	fmt.Println(string(f))

	//读取所有内容4---先从file读取到reader
	reader := bufio.NewReader(file)
	ll := 2000
	buf := make([]byte,ll)
	var buffer bytes.Buffer
	for {
		_,err := reader.Read(buf)

		//写入
		buffer.Write(buf)
		if err == io.EOF {
			break
		}

	}
}