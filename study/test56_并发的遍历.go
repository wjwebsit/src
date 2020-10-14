package main

import (
	"fmt"
	"strings"
	"sync"
)

func main()  {
	//定义信号
	c := make(chan []string)
	go func() {
		c<- []string{"hello","word","xuxiao"}
	}()

	//获取
	str := <- c
	close(c)
	var ch = make(chan string,len(str))

	go func(s []string) {
		ch <- strings.Join(str,"-")

	}(str)
	go func(s []string) {
		ch <- strings.Join(str,",")

	}(str)
	go func(s []string) {
		ch <- strings.Join(str,"&&")

	}(str)

	//等待----缺点如果某一个协程中断会阻塞
	for i := 0; i < len(str);i++ {
		fmt.Println(<-ch)
	}
	close(ch)

	//----利用线程池
	var wg sync.WaitGroup
	wg.Add(3)
	ch2 := make(chan string,3)

	go func(s []string) {
		defer wg.Done()
		//ch2 <- strings.Join(str,"-")

	}(str)
	go func(s []string) {
		defer wg.Done()
		ch2 <- strings.Join(str,",")

	}(str)
	go func(s []string) {
		defer wg.Done()
		ch2 <- strings.Join(str,"&&")

	}(str)
	go func() {
		wg.Wait()
		close(ch2)
	}()

	//获取---注释一个信号写入也不会失败
	for x := range ch2 {
		fmt.Println(x)
	}
}
