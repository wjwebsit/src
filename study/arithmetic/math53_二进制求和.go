package main

import (
	"fmt"
	"strconv"
)

/**
	算法描述
	给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
 */

func addBinary(a string, b string) string {
	//获取长度
	lengthA  := len(a)
	lengthB := len(b)

	//比较取最大,最小
	length := lengthA
	min := lengthA

	if lengthA < lengthB {
		length = lengthB
	} else {
		min = lengthB
	}

	//声明map
	mp := map[byte]int{}
	mp['0'] = 0
	mp['1'] = 1

	//高位补位0
	str := ""
	for i := 1;i <= length - min;i++ {
		str += "0"
	}
	if length > lengthA {
		a = str + a
	} else {
		b = str + b
	}

	//进位
	carry := 0

	//末尾
	index := length - 1

	//声明结果
	result := ""
	for index >= 0 {
		tmp := carry + mp[a[index]] + mp[b[index]]

		//判断
		if tmp >= 2 {
			//进位
			carry = 1
			tmp = tmp % 2
		} else {//进位归0
			carry = 0
		}

		result = strconv.Itoa(tmp) + result
		index --

	}

	//判断末尾
	if carry == 1 {
		result = "1" + result
	}

	//返回
	return result

}

//测试
func main() {
	fmt.Println(addBinary("1010","1011"))
}