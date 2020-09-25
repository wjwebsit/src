package main

import "fmt"

/**
	算法描述
	编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
说明:
元音字母不包含字母"y"。

元音字母【a,o,e,i,u】
 */

func reverseVowels(s string) string {
	//判断
	if len(s) == 0 {
		return s
	}
	//获取字节数组
	sArr := []byte(s)

	//声明元音hash
	//声明元音hash
	mp := map[byte]bool {
		'a' : true,
		'A' : true,
		'o' : true,
		'O' : true,
		'e' : true,
		'E' : true,
		'i' : true,
		'I' : true,
		'u' : true,
		'U' : true,
	}

	//利用双指针
	low,high := 0,len(sArr) - 1
	for low <= high {
		//从前面找到一个元音
		for low <= high && !mp[sArr[low]] {
			low ++
		}

		//从后面找到一个元音
		for low <= high && !mp[sArr[high]] {
			high --
		}

		//交换
		if low < high {
			sArr[low] = sArr[low] ^ sArr[high]
			sArr[high] = sArr[low]^sArr[high]
			sArr[low] = sArr[low] ^ sArr[high]
		}

		//指针移动
		low++
		high--
	}

	//返回
	return string(sArr)

}
func main() {


}