package main

import (
	"fmt"
	"time"
)

/**
	在go程序中每一个并发的单元叫做goroutine（协程）可以当作线程来处理
	利用 go 关键字来开启

 */
func main() {
	n := 45
	//开启协程
	go spinner(100 * time.Millisecond)

	fmt.Println(fio(n))
}
func fio(n int) int{
	if n < 2 {
		return n
	}
	return fio(n - 1) + fio(n - 2)
}
/***
	动画效果
 */
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}