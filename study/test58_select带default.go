package main

import (
	"fmt"
)

/**
	如果channel没有满则会等待 满了会default,没有缓存的直接进default
    select中的每一个case 都是针对channel的IO操作，读或写，当有多个case同时就绪（可读可写）时，select会随机选择一个执行
	select{}将会永远阻塞；通常select与for并用，这样select就可以循环不断地评估每一个case，随机选取一个case执行。
 */

func main() {
	//测试没有缓存的
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch1 <- sum
	}()
	select {
	case x := <-ch1:
		fmt.Println(x)
	default:
		fmt.Println("没有缓存的直接default")
	}

	//测试带缓存的---读default
	ch2 := make(chan int, 2)

	for i := 0; i < cap(ch2); i++ {
		go func() {
			//defer close(ch2)
			sum := 0
			for i := 0; i < 10; i++ {
				sum += i
			}
			ch2 <- sum
		}()
	}
	defer close(ch2)

	//便利
	for i:= 0; i < 2;{
		select {
		case x := <-ch2:
			fmt.Println(x)
			i ++ //来了一个
		default:
			fmt.Println("带缓存的default")
		}

	}

	//测试带缓存的写入---
	ch3 := make(chan int,2)
	fg := 0
	for {
		if fg == 1 {
			break
		}
		select {
		case ch3 <- 2:
			fmt.Println(len(ch3))
		default:
			fg = 1
			fmt.Println("服务器繁忙，请稍后再试！")
			close(ch3)
			break

		}
	}
	for x := range ch3 {
		fmt.Println(x)
	}

}
