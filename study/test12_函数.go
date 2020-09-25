package main

import (
	"fmt"
)

/**
	函数是指可以执行的一段代码段
	func 函数名(参数列表)（返回值列表） {代码段}
	go语言函数可以返回多个返回值
	返回值不需要使用时，可以用_来忽略
	可变函数（*****）
 */

func main () {
	//测试可变函数
	xuxiao3(3,4,5)
	xuxiao3(3,4,5,6)
}

//返回值为1个
func xuxiao1()int {
	return 1
}
//返回值为多个
func xuxiao2(a int,b string)(int,string) {
	return a,b
}

//可变函数
func xuxiao3(arg ...int) {
	//参数列表是一个[]int的切片
	fmt.Print(arg)

}


