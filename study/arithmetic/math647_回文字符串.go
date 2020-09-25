package main

import "fmt"

/**
	算法描述:
	给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。

示例 1:

输入: "abc"
输出: 3
解释: 三个回文子串: "a", "b", "c".
示例 2:

输入: "aaa"
输出: 6
说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".
注意:

输入的字符串长度不会超过1000。
 */
 /**
 	s(i,j) = s[i]==s[j] && s[i + 1][j - 1] 为回文子串
  */
func countSubstrings(s string) int {
	//判断
	if len(s) == 0 {
		return 0
	}

	//采用动态规划----优化空间
	var dp  = make(map[string]bool)
	dp[""] = true

	//统计
	count := 0

	//初始化
	for i := 0; i < len(s); i ++ {
		dp[string(s[i: i+1])] = true
		count ++
		if i + 1 < len(s) {
			if s[i + 1] == s[i] {
				dp[string(s[i:i+2])] = true
				count ++
			}  else {
				dp[string(s[i:i+2])] = false
			}
		}
	}

	//dp
	for i := len(s) - 1;i >= 0 ; i -- {
		for j := i + 1; i < len(s); i ++ {
			//判断
			if s[i] == s[j] && dp[s[i+1:j]] == true {
				count ++
				dp[string(s[i:j])] = true
			} else {
				dp[string(s[i:j])] = false
			}
		}
	}


	//返回
	return count
}
func main() {
	fmt.Println(countSubstrings("aaa"))
}