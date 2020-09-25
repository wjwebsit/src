package main

import (
	"fmt"
	"os"
)

func main() {
	//获取当前目录
	fmt.Println(os.Getwd())

    //改变当前目录
   // os.Chdir("~/Documents/")

   //创建目录
   os.Mkdir("test",os.ModePerm)

   //创建多级目录
   //os.MkdirAll("test/std/fsd",os.ModePerm)

   //删除文件或目录
   cur ,_ := os.Getwd()
  // os.Remove(cur + "/test/std/fsd") //删除了fsd
  // fmt.Println(os.Remove(cur + "/test/")) //如果目录不为空(有子文件或目录)则无法删除
  //删除所有
  //os.RemoveAll(cur + "/test/")

  //重命名文件
  os.Rename(cur + "/test",cur + "/hh")

   //返回临时目录
   fmt.Println(os.TempDir())
}
