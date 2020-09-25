package main

import "fmt"

/**
	递归函数---自己调用自己  注意防止内存栈溢出
 */

 func main() {
 	//测试1+....n
 	sum := 0
 	dg1(100,&sum)

 	fmt.Print(sum)
 }
 /**
 递归求1...n的和
  */
 func dg1 (n int,sum *int) {
 	if n >=1 {
		(*sum) = (*sum) + n
 		dg1(n - 1,sum)
	}
 }
