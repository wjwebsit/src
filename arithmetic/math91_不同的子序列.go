package main

import "fmt"

/**
	题目描述：
	给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。

一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）

示例 1:

输入: S = "rabbbit", T = "rabbit"
输出: 3
解释:

如下图所示, 有 3 种可以从 S 中得到 "rabbit" 的方案。
(上箭头符号 ^ 表示选取的字母)

rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^
示例 2:

输入: S = "babgbag", T = "bag"
输出: 5
解释:

如下图所示, 有 5 种可以从 S 中得到 "bag" 的方案。
(上箭头符号 ^ 表示选取的字母)

babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
    ^^^
 */
func numDistinct(s string, t string) int {
	//判断
	if len(s) < len(t) {
		return 0
	}

	//采用动态规划
	var dp [][]int

	//初始化
	for i := 0; i <= len(t); i ++ {
		dp = append(dp,[]int{})
		dp[i] = append(dp[i],0)
	}
	dp[0] = []int{}
	for j := 0; j <= len(s); j ++ {
		dp[0] = append(dp[0],1)
	}

	//遍历
	for i := 1; i <= len(t); i ++ {
		for j := 1; j <= len(s);j++ {
			//判断
			if s[j - 1] == t[i - 1] {
				dp[i] = append(dp[i],dp[i][j - 1] + dp[i - 1][j - 1])
			} else {
				dp[i] = append(dp[i],dp[i][j - 1])
			}
		}
	}

	//返回
	return dp[len(t)][len(s)]
}
func main() {
	s,t := "rabbbit","rabbit"
	fmt.Println(numDistinct(s,t))
}