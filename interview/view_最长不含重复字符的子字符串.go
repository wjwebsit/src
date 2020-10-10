package main

import "fmt"

/**
	请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

 

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
提示：
s.length <= 40000
 */

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	//声明散列
	hash := make(map[byte]int)
	hash[s[0]] = 0
	max := 1
	i := 0
	for j := 1; j < len(s);j ++ {
		//判断是否重复
		if _,ok := hash[s[j]]; ok { //重复了
			//计算长度
			ll := j - i

			//获取长度
			if ll > max {
				max = ll
			}

			//更新左区间
			if hash[s[j]] + 1 > i {
				i = hash[s[j]] + 1
			}
		}

		//写入散列
		hash[s[j]] = j
	}

	//最后的
	if len(s) - i > max {
		max = len(s) - i
	}

	//返回
	return max
}
func main() {
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}