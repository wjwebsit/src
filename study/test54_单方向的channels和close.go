package main

import "fmt"

/**
	随着程序的增长，人们习惯将大的函数拆分成小的函数。
	当一个channel作为一个函数参数时，它一般总是被专门用于发送或者接收。
    Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel。
    类型chan<- int表示一个只发送int的channel，只能发送不能接收。
    相反，类型<-chan int表示一个只接收int的channel，只能接收不能发送。（箭头<-和关键字chan的相对位置表明了channel的方向。）这种限制将在编译期检测。
 */

/**
	chan <- int //	只能发送赋值给它
	<- chan int //	只能接收，不能发送
 */

type suGGer struct {
}
func (s suGGer) getNum1(ch chan<-int) {
	for i := 1; i <= 3; i ++ {
		ch <- i
	}
	close(ch)
}
func (s suGGer) getNum2(ch1 <-chan int,ch2 chan <- int) {
	for x := range ch1 {
		ch2 <- x * x
	}
	close(ch2)
}
func (s suGGer) print(ch <-chan int) {
	for x := range ch {
		fmt.Println(x)
	}

	//close(ch) 关闭了只读的 重复关闭
}

func main() {
	ch1,ch2 := make(chan int),make(chan int)
	var s suGGer
	go s.getNum1(ch1)
	go s.getNum2(ch1,ch2)
	s.print(ch2)

	//close函数---关闭只写的 chan <- Type 或 双向的
	//关闭nil或重复关闭会报错
	var ch3 chan int
	fmt.Println(ch3 == nil)
	//close(ch3) //关闭了nil
	var ch4 = make(<-chan int)
	<-ch4

	//只读取的无法关闭
	close(ch4)

}


