package main

import "fmt"

/**
	匿名函数（程序中一般只执行一次）func (形参数列表) 返回值 {方法体}（实参存在直接调用）
 */

 type myFunc func(int) int //定义函数类型参数

 func main() {
 	//匿名函数的用法1 ---定义完成直接执行
 	func(a int){
 		fmt.Print(a)
	}(2)

 	result := func(a int) int {
 		return a * a
 	}(2)

 	fmt.Println(result)

 	//匿名函数2-----赋给函数变量

 	//声明匿名函数并赋值给f
 	f := func(a int) int {
 		return a * a
	}
 	//调用匿名函数
 	fmt.Println(f(2))

 	//匿名函数--作为回调函数的参数
 	result = method(2,func(a int)int{
 		return a + a
	})

 	fmt.Println(result)

 	result = method(3,func(a int) int {
 		return a << 2
	})

 	fmt.Println(result)



 }

 /**
 定义闭包函数method
  */
 func method(a int,f myFunc) int {
 	return f(a)
 }
