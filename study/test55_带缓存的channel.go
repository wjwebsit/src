package main

import (
	"fmt"
	"sync"
)

/**
带缓存的Channel内部持有一个元素队列。队列的最大容量是在调用make函数创建channel时通过第二个参数指定的。
如果制定为0则相等于不带缓存的，当缓存满了会发生阻塞
带缓存的channel当作同一个goroutine中的队列使用 不可取
 */

func main() {
	//带缓存
	ch := make(chan int ,3)
	ch <- 1
	ch <- 2
	ch <- 3
	//ch <- 4此时阻塞
	//写完之后要关闭否则报错 --如果调用为for range
	close(ch)

	for x := range ch {
		fmt.Println(x) // 1,2,3
	}


	//获取长度
	fmt.Println(len(ch)) // 0

	//获取容量
	fmt.Println(cap(ch)) // 3

	//1件事3个并发取最快的
	var w sync.WaitGroup
	w.Add(3)
	ch1 := make(chan float64,3)

	go  func () {
		defer w.Done()
		sum := 0
		for i := 0; i < 10; i ++ {
			sum += i
		}
		//没有关闭
		ch1 <- float64(sum)


	}()
	go func() {
		defer w.Done()
		sum := 0
		for i := 0; i < 100000; i ++ {
			sum += i
		}

		//没有关闭
		ch1 <- float64(sum)


	}()

	go func() {
		defer w.Done()
		sum := 0
		for i := 0; i < 10000000; i ++ {
			sum += i
		}

		//没有关闭
		ch1 <- float64(sum)

	}()
	go func(){
		w.Wait()
		close(ch1)
	}()

	for x := range ch1 {
		fmt.Println(x)
	}

	//并发执行
	fm := sum1()
	fmt.Println(fm)



}
func sum1() []int {
	//获取所有
	ch2 := make(chan int,3)
	var w sync.WaitGroup
	w.Add(3)
	go  func () {
		defer w.Done()
		sum := 0
		for i := 0; i < 10; i ++ {
			sum += i
		}
		ch2 <- sum

	}()
	go func() {
		defer w.Done()
		sum := 0
		for i := 0; i < 100000; i ++ {
			sum += i
		}
		ch2 <- sum


	}()

	go func() {
		defer w.Done()
		sum := 0
		for i := 0; i < 10000000; i ++ {
			sum += i
		}
		ch2 <- sum

	}()
	go func() {
		w.Wait()
		defer close(ch2)
	}()

	var result []int
	for x := range ch2 {
		result = append(result,x)
	}

	return result
}
