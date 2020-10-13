package main

import "fmt"

/**
随着程序的增长，人们习惯于将大的函数拆分为小的函数。我们前面的例子中使用了三个goroutine，然后用两个channels来连接它们，它们都是main函数的局部变量。将三个goroutine拆分为以下三个函数是自然的想法：
 */

/**
	单方向---对于函数而言
	<-chan int //只能获取数据
	chan <- int //只能写入
 */
type sugar struct {
	in,out chan int
}
func (s sugar) mkNum1() {
	for i := 1; i <= 3; i ++ {
		s.in <- i
	}
	close(s.in)
}
func (s sugar)mkNum2() {
	for x := range s.in {
		s.out <- x * x
	}
	close(s.out)
}
func (s sugar) print() {
	for x := range s.out {
		fmt.Println(x)
	}
	//close(s.out)
}

func main()  {
	var s sugar
	s.in = make(chan int)
	s.out = make(chan int)
	go s.mkNum1()
	go s.mkNum2()
	go s.print()

}

