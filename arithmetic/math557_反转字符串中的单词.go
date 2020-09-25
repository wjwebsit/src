package main

import (
	"fmt"
	"strings"
)

/**
	给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc" 
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

 */
func reverseWords2(s string) string {
	//判断字符串
	if len(s) == 0 {
		return ""
	}

	//声明返回值
	str := []string{}

	//获取单词
	i := 0
	for i < len(s) {
		//判断是否非空
		if s[i] != ' ' {

			//构建单词
			word := string(s[i])


			//获取单词
			j := i + 1
			for j < len(s) && s[j] != ' ' {
				//反转单词
				word = string(s[j]) + word
				j++

			}

			//写入
			str = append(str,word)

			//继续 j此时为‘ ’
			i = j + 1


		} else {
			i ++
		}

	}

	//返回
	return strings.Join(str," ")

}
func main() {
	str := " Let's take LeetCode contest "
	fmt.Println(reverseWords2(str))
}