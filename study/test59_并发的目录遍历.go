package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	//获取当前目录
	root,_ := os.Getwd()

	//定义遍历目录
	dirs := make([]string,4)
	dirs[0] = root + "/src/data/"


	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan []string,40)
	go getFiles(dirs[0],&wg,ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch{
		fmt.Println(x)
	}
}
func getFiles(dir string,w *sync.WaitGroup,ch chan []string) {
	defer w.Done()

	//打开
	dirs,err := os.Open(dir)
	if err != nil {
		ch <- []string{}
		return
	}

	//读取目录
	files,err := dirs.Readdir(-1)
	if err != nil {
		ch <- []string{}
		return
	}

	var result []string
	for _,file := range files  {
		//判断是否为目录
		if file.IsDir() == true {
			//获取目录
			dirStr := dirs.Name() + "/" + file.Name()
			w.Add(1)
			go getFiles(dirStr,w,ch)

		}
		//写入
		result = append(result,file.Name())
	}
	ch <- result
}
