package main

import (
	"fmt"
	"os"
)

/**
	文件相关：

 */
func main()  {
	//获取当前目录
	curDir, _ := os.Getwd()

	//在当前目录下创建files--如果存在会先删掉在生成
	os.Mkdir("files",os.ModePerm)

	//创建文件---(文件名)----返回打开后的文件指针
	f1,_ := os.Create(curDir + "/files/" + "xuxiao.go",)

	//字符串内容---长字符串
	content := `package main 
import "fmt"
func main() {
	fmt.Println("hello Word")
}
`
	//写入string----WriteString追加写
	f1.WriteString(content)


	//写入字符
	bt := []byte{'a','b','c'}
	fmt.Println(f1.Write(bt))  //返回写入的长度和错误

	//获取文件名称-----绝对路径
	fmt.Println(f1.Name())

	//获取文件的path info
	pathInfo , _ := f1.Stat()
	pathInfo.Name() //文件名称不包含目录
	fmt.Println(pathInfo.Name())
	pathInfo.Size() //byte计算
	fmt.Println(pathInfo.ModTime()) //修改时间
	pathInfo.IsDir() //是否是目录
	fmt.Println(pathInfo.Mode()) //Linux文件权限
	fmt.Println(pathInfo.Sys()) //底层数据来源

	//设置文件的权限
	fmt.Println(f1.Chmod(777))

	fmt.Println(f1.Fd())

	//关闭文件
	f1.Close()

	//打开目录
	dir,_ := os.Open(curDir + "/files/")

	//获取目录下的所有文件
	//files,_ := dir.Readdir(-1) // n > 第一个，n < 0 所有 返回*file--目录树下直接孩子
	fmt.Println(fileExist(curDir + "/files/"))

	//遍历文件目录的所有文件
	fmt.Println(getAllFiles(dir))

	//删除文件下的所有的文件和文件夹
	os.RemoveAll(dir.Name())

}
/**
	判断目录和文件是否存在
 */
func fileExist(fileName string) bool {
	//判断文件是否存在
	_,err := os.Stat(fileName)

	//一定存在
	if err == nil {
		return true
	}

	//判断
	if os.IsExist(err) {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	//不确定是否存在就不存在好了
	return false
}
/**
	遍历目录下的所有的文件
	注意点os.File.Name() 返回绝对路径
	os.File.Readdir(-1) //返回所有文件或文件夹当前的直接孩子 并且是 os.FileInfo 类型
	os.File.Name() + "/" + file.Name()来拼装自文件的目录 并打开传递
 */
func getAllFiles(dir *os.File) []string{
	//声明返回值
	var result []string

	//读取所有的文件或目录
	files, err := dir.Readdir(-1)

	//判断
	if err != nil {
		return result
	}

	//遍历目录
	for _,file := range files {
		if file.IsDir() == false {
			result = append(result,file.Name())
		} else {
			//读取文件
			ff,err := os.Open(dir.Name() +"/" + file.Name())
			if err == nil {
				result = append(result, getAllFiles(ff)...)
			}
		}
	}

	//返回
	return result
}
