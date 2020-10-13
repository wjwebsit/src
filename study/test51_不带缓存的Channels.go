package main

import "fmt"

/**
不带缓存的Channels:不带缓存的channel将导致发送的协程阻塞，直到另一个协程在相同的协程上执行接受操作
当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作。

基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。因为这个原因，无缓存Channels有时候也被称为同步Channels。
当通过一个无缓存Channels发送数据时，接收者收到数据发生在再次唤醒唤醒发送者goroutine之前（译注：happens before，这是Go语言并发内存模型的一个关键术语！）。
*/

/**
	并发编程：并发（不是一起执行）发生前（y执行时 x已经完毕，y结束前需要x的参数，y时间步骤太长，x可能在中间环节才执行）
    我们说x事件在y事件之前发生 x-> y（happens before），我们并不是说x事件在时间上比y时间更早；我们要表达的意思是要保证在此之前的事件都已经完成了，例如在此之前的更新某些变量的操作已经完成，你可以放心依赖这些已完成的事件了。
	x事件既不是在y事件之前发生也不是在y事件之后发生，我们就说x事件和y事件是并发的x事件和y事件就一定是同时发生的，我们只是不能确定这两个事件发生的先后顺序

 */
func main() {
	ch := make(chan struct{})
	/*go func() {
		fmt.Println("123")
		ch <- struct{}{}
	}()*/

	go test_channel(&ch)

	<-ch
	close(ch)
	fmt.Println(456)

}
func test_channel(c *chan struct{}) {
	fmt.Println("hello word")
	*c <- struct{}{}
}