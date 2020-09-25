package main

import (
	"fmt"
	"strings"
)

/**
	算法描述:
	给定一个字符串，逐个翻转字符串中的每个单词。

 

示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

 */
func reverseWords(s string) string {
	//声明反转后的字符串
	i := len(s) - 1
	var word []string

	//从后往前
	for i >= 0{
		//判断i是否为空
		if s[i] != ' ' {//i位置为单词结束
			//双指针
			j := i - 1
			for j >= 0 && s[j] != ' ' {
				j --
			}

			//截取
			word =  append(word,string(s[j + 1: i + 1]))


			//j此时为空格 i = j - 1
			i = j - 1

		} else {
			i --
		}


	}

	//返回
	return strings.Join(word," ")
}
func main() {
	s := ""
	fmt.Println(reverseWords(s))
}