package main

import "fmt"

/**
	可变参数的函数类型
	func 函数(形参1 ...(3ge) 数据类型) 返回值{} 形参1就是数据类型的一个切片

	访问 1 正常函数访问
    访问2  func(对应数据类型的切片...) 切片的0，1，2.....分别为函数的第1，2，3......参数

 */

 func main() {

 	//调用方法1
 	fmt.Print(makeSum(1,2,3))

 	//调用方法2 切片...
 	arg := []int{1,2,3}
 	fmt.Print(makeSum(arg...))

 	//扩大n倍
 	fmt.Println(makeSum2(2,arg...))

 }

 func makeSum(a ...int) int {
 	//求和
 	sum := 0
 	for _,value := range a {
 		sum += value
	}

 	//返回
 	return sum

 }
 /**
 	参数列表有一个为可变函数，且那个参数必须在最后
  */
 func makeSum2(n int,a ...int)int {
	 //求和
	 sum := 0
	 for _,value := range a {
		 sum += value
	 }

	 return n * sum
 }
