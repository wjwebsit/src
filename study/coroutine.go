package main

import "fmt"

func main() {
	p:= myAdd(echo(1,50),echo(51,100))

	for v := range p{
		fmt.Println(v)
	}
}

func myAdd(ch1,ch2 chan int) chan int{
	var ch = make(chan int)
	go func() {
		ch <- <-ch1 + <-ch2
		defer close(ch)
	}()
	return ch
}

func echo(s,e int) chan int{
	rt := make(chan int)
	go func(s,e int) {
		sum := 0
		for s <= e {
			sum += s
			s++
		}

		rt <- sum
		defer close(rt)

	}(s,e)

	//返回
	return rt



}

