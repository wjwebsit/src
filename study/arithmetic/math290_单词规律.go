package main

import "strings"

/**
	单词规律
	给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false
说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。    

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-pattern
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func wordPattern(pattern string, str string) bool {
	if len(str) == 0 {
		return  false
	}
	if len(pattern) == 0 {
		return true
	}

	//分割
	strArr := strings.Split(str," ")

	//判断
	if len(strArr) != len(pattern) {
		return  false
	}

	//定义映射
	hash1 := make(map[byte]string)
	hash2 := make(map[string]byte)
	var i int
	for i = 0 ;i < len(pattern);i++ {
		//判断是否村子
		if _,ok := hash1[pattern[i]];ok {
			if hash1[pattern[i]] != strArr[i] {
				break
			}

		} else {//不能和之前的重复
			if _,ok := hash2[strArr[i]]; ok && hash2[strArr[i]] != pattern[i] {
				break
			}
			hash1[pattern[i]] = strArr[i]
			hash2[strArr[i]] = pattern[i]
		}

	}

	//返回
	return i == len(pattern)

}