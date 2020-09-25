package main

import "fmt"

/**
	算法描述：
	给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指由字母组成，但不包含任何空格的字符串。

示例:

输入: "Hello World    "
输出: 5
 */

 /**
 	主要是最后的空白
  */
func lengthOfLastWord(s string) int {
	//声明返回
	rtIndex := 0

	//判断
	if len(s) == 0 {
		return rtIndex
	}

	//判断末尾
	end := true   //是否为最后一个单词的末尾
	for j := len(s) - 1; j >= 0;{
		//寻找末尾第一个非" "
		for j >= 0 && s[j] == ' ' && end   {
			j --
		}
		//末尾找到
		if end {
			end = false
		} else { //计数
			if s[j] != ' ' {
				rtIndex ++
				j --
			} else {
					break
			}

		}

	}

	//返回
	return rtIndex
}

func main() {
	fmt.Println(lengthOfLastWord("a "))

}
