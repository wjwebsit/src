package main

import (
	"fmt"
)

/**
	算法描述:
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？
说明：m 和 n 的值均不超过 100。

示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
 */
func uniquePaths(m int, n int) int {
	//判断条件
	if m == 0 || n == 0  {
		return 0
	}

	//采用动态规划A、dp[m][n] = dp[m][n-1] + dp[m-1][n](m>1,n>1)B、m=1 dp[m][n] = 1 n=1 dp[m][n] = 1

	var dp [][]int

	//初始化
	for i := 0; i < m; i++ {
		dp = append(dp,[]int{})
	}

	//开始
	for i := 0; i < m; i ++ {
		for j := 0; j < n; j ++ {
			if i == 0 || j == 0 {
				dp[i] = append(dp[i],1)
				continue;
			}

			dp[i] = append(dp[i],dp[i - 1][j] + dp[i][j - 1])
		}
	}

	//返回
	return dp[m - 1][n - 1]
}


/**
	test
 */
 func main() {
	fmt.Println(uniquePaths(3,2))
 }



