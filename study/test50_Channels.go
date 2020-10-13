package main

import (
	"fmt"
	"reflect"
)

/**
	如果协程是并发体的话那么channels就是并发体之间的通信。一个channel是一个通信机制，他可以让一个
	协程给另外一个协程发型信息。每个channel都有一个特殊的类型，也就是channels可发送的数据类型。
	如 int 写为 channel int.

 */

/***

 */
func main() {
	//创建信号--
	ch := make(chan int,1)
	fmt.Print(reflect.TypeOf(ch).String() == "chan int")  //---类型为chan(1个空格)int

	//和map类似，channel也对应一个make创建的底层数据结构的引用
	ch <- 2 //赋值
	x := <-ch //输出取决于chan 类型
	fmt.Print(x == 2)

	ch2 := ch
	ch2 <- 3
	fmt.Println(<-ch) //3 传引用的

	//两个相同类型的channel可以使用==运算符比较。
	//如果两个channel引用的是相同的对象，那么比较的结果为真。
	//一个channel也可以和nil进行比较。
	ch3 := make(chan int,1)
	ch3 <- 3
	fmt.Println(ch3 == ch) //false即使类型和值都相同

	//比较来自同一引用值
	fmt.Println(ch2 == ch) //true

	//声明没有初始化
	var ch4 chan int //没有声明缓冲区
	fmt.Println(ch4 == nil) //true
	//ch4 <- 3 ---报错为nil

	ch5 := make(chan int,1)
	fmt.Println(ch5 == nil) //false
	ch5 <- 4
	ch4 = ch5 //打印出错误 因为ch4没有缓冲区
	//fmt.Println(<-ch5,<-ch4) //fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch5) //4

	//测试一下copy
	ch7 := make(chan int,2)
	ch8 := make(chan int,2)

	//报错了
	//copy(ch7,ch8)

	//关闭信号
	close(ch7)
	close(ch8)

}
