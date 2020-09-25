package main

import "fmt"

/**
	算法描述:

难度
简单

31

收藏

分享
切换为英文
关注
反馈
泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。



示例 1：

输入：n = 4
输出：4
解释：
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
示例 2：

输入：n = 25
输出：1389537


提示：

0 <= n <= 37
答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
 */
func tribonacci(n int) int {
	//采用动态规划
	var dp []int
	dp = append(dp,0) // dp[0] = 0
	dp = append(dp,1) //dp[1] = 1
	dp = append(dp,1) //dp[2] = 1
	if n >= 3 {
		for k := 3; k <= n; k ++ {
			dp = append(dp,dp[k - 1] + dp[k - 2] + dp[k - 3])
		}

	}

	//返回
	return dp[n]
}

func main()  {
	fmt.Println(tribonacci(25))
}