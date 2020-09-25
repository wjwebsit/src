package main


/**
题目描述：
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 */
func minimumTotal(triangle [][]int) int {
	//判断
	if len(triangle) < 0 {
		return 0
	}

	//获取行数
	rows := len(triangle)

	//采用动态规划 ---dp表示当前行的j列位置的最小值
	var dp []int

	//初始化
	dp = triangle[rows - 1]  //自底向上最后一行即为dp[rows - 1]

	//遍历 倒数第二行开始
	for row := rows - 2;row >= 0; row -- {
		//每一行有 row +1 列
		for col := 0; col <= row;col ++ {
			//当前可以触及的下一层
			dp[col] = triangle[row][col] + minInt1(dp[col],dp[col + 1])
		}
	}

	//返回
	return dp[0]
}
//比较2个整数
func minInt1(a,b int) int{
	if a < b {
		return a
	} else {
		return b
	}
}