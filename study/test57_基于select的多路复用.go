package main

import (
	"fmt"
	"time"
)

/**
与 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。
 */

func main() {
	//获取执行速度最快的
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	//协程
	go func() {
		defer close(ch1)
		sum := 0
		for i := 1;i <= 100;i++ {
			sum += i
		}

		//写入信号
		ch1 <- sum

	}()

	//协程
	go func() {
		defer close(ch2)
		sum := 0
		for i := 1;i <= 100000;i++ {
			sum += i
		}

		//写入信号
		ch2 <- sum

	}()

	//协程
	go func() {
		defer close(ch3)
		sum := 0
		for i := 1;i <= 1000000;i++ {
			sum += i
		}
		fmt.Println(sum)

		//写入信号
		ch3 <- sum

	}()

	//多路复用获取执行速度最快的
	select {
	case x := <-ch1: //获取值
		fmt.Println(x)

	case x := <-ch2:
		fmt.Println(x)

	case x := <-ch3:
		fmt.Println(x)
	}

	//三个协程都会执行的
	time.Sleep(3 * time.Second)

	//利用select设置超时机制
	ch4 := make(chan int)
	ch5 := make(chan int)
	go func() {
		defer close(ch4)
		time.Sleep(time.Second * 1)
		ch4 <- 1

	}()
	select {
	case <- ch5:
		fmt.Println("ch4接收到了")
	case <-	ch4:
		fmt.Println("超时了")
	}

}