package main

import (
	"fmt"
	"strconv"
)

/**
算法问题描述
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例1：

输入： 121
 输出： true
示例2：

输入： -121
 输出： false
 解释：从左向右读，为-121。从右向左读，为121-。因此它不是一个回文数。
示例3：

输入： 10
 输出： false
 解释：从右向左读，为01。因此它不是一个回文数。
 */

func main() {
	fmt.Print(isPalindrome(-121))
}
/**
两端不相等肯定不是，如果相等一次从两端往中心延申
算法和求KMP模式匹配算法的求next数组雷同
 */
func isPalindrome(x int) bool {
	//将整数转换成字符串
	s := strconv.Itoa(x)

	//定义起始
	begin, end := 0,len(s) - 1

	//定义fg,默认为是回文数
	fg := true

	//判断
	for begin < end {//中心递归未相遇过遍历
		if s[begin] == s[end] {
			begin++
			end--
		} else {
			//不是回文数
			fg = false
			break
		}


	}

	return fg


}
