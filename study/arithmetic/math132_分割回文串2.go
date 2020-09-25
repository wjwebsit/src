package main
/**
	算法描述
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回符合要求的最少分割次数。

示例:

输入: "aab"
输出: 1
解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
 */
func minCut(s string) int {
	//预处理获取回文串
	var dp = make([][]int,len(s))
	cache := make([]int,len(s))
	for i:= 0; i < len(s); i ++ {
		dp[i] = make([]int,len(s))
		dp[i][i] = 1
		cache[i] = -1
		if i + 1 < len(s) && s[i + 1] == s[i] {
			dp[i][i + 1] = 1
		}
	}
	for i := len(s) - 3; i >= 0;i -- {
		for j := len(s) - 1; j > i + 1; j -- {
			if s[i] == s[j] && dp[i + 1][j - 1] == 1 {
				dp[i][j] = 1
			}
		}
	}

	var split = make ([]int,len(s) + 1)
	split[len(s)] = 0
	split[len(s) - 1] = 0
	//改为自底向上
	for i := len(s) - 2; i >= 0; i -- {
		split[i] = len(s) - 1 - i
		//判断
		for j := i; j < len(s); j ++ {
			if dp[i][j] == 1 {
				split[i] = min(split[i], 1 + split[j + 1])
			}

		}
	}

	//返回
	return split[0]

}
func minCutHS(s string,begin int,dp [][]int,cache *[]int) int{
	//判断
	if begin == len(s) {
		return -1
	}
	//判断cache
	if (*cache)[begin] >= 0 {
		return (*cache)[begin]
	}

	//声明
	m := len(s)

	for i := begin; i < len(s);i++ {
		if dp[begin][i] == 1 {
			m = min(m,1 + minCutHS(s,i + 1,dp,cache))
		}

	}

	//写入
	(*cache)[begin] = m

	//返回
	return m

}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}