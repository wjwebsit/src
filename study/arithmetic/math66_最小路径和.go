package main

import (
	"math"
)

/**
	算法描述:
	给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
 */
/**
	dynamic programming
 */
func minPathSum(grid [][]int) int {
	//判断
	if len(grid) == 0 {
		return 0
	}

	//或行和列
	rows,cols := len(grid) - 1,len(grid[0]) - 1

	//声明dp
	var dp [][]int
	for i := 0;i <=rows;i++ {
		dp = append(dp,[]int{})
	}
	dp[0] = append(dp[0],grid[0][0])

	//处理第一行和第一列
	for k := 1;k <= rows;k ++ {
		dp[k] = append(dp[k],dp[k - 1][0] + grid[k][0])
	}
	for col := 1;col <= cols;col++ {
		dp[0] = append(dp[0],dp[0][col - 1] + grid[0][col])
	}

	for i := 1; i <= rows;i++ {
		for j := 1; j <= cols;j++ {
			dp[i] = append(dp[i],min2(dp[i - 1][j],dp[i][j - 1]) + grid[i][j])
		}
	}

	//返回
	return dp[rows][cols]

}
/**
	比较大小
 */
func min2(a,b int ) int {
	if a < b {
		return a
	} else {
		return b
	}
}


