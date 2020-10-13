package main

import "fmt"

/**
	描述：串行 A -> B -> C  A->B之间一个channels B->C之间一个channels

    Channels也可以用于将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的输入。
	这种串行的channels就是所谓的管道

 */

func main() {
	//定义2个并发信号
	num1 := make(chan int)
	num2 := make(chan int)

	//获取乘积B
	go func () {
		for x := range num1 {
			num2 <- x * x
		}
		close(num2)
	}()

	//生成数字A
	go func() {
		for i := 1; i <= 3; i ++ {
			num1 <- i
		}
		close(num1)
	}()

	//输出信号C
	for x := range num2 {
		fmt.Println(x)
	}

	//不用关闭num2 因为是读取

	//B等待A
	// for range 具有等待的功能
}
