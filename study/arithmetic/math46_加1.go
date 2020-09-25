package main

import "fmt"

/**
	算法描述
	给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

 */
func plusOne(digits []int) []int {
	//判断
	if len(digits) == 0 {
		return digits
	}

	//定义进位
	carry := 0

	index := len(digits) -1

	//末尾+1
	digits[index] ++

	//判断
	for  index >= 0 && (carry + digits[index]) >=10 {
		//处理当前位数
		digits[index] = (carry + digits[index]) - 10

		//进位
		carry = 1

		//向前
		index --

	}

	//处理最后一位
	if index == -1 && carry == 1{
		//切片的unshift
		digits = append([]int{1},digits...)

	} else {//处理当前位
		digits[index] = carry + digits[index]
	}

	//返回
	return digits

}

func main() {
	arr := []int {9,9,9}

	fmt.Println(plusOne(arr))
}
