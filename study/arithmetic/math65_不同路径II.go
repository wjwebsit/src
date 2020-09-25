package main

import "strconv"

/*
	算法描述
	一个机器人位于一个mxn网格的左上角（起始点在下图中标记为“开始”）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“完成”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
示例1：

输入：
 [
  [0,0,0]
  [0,1,0]
  [0,0,0]
]
输出： 2
 解释：
3x3网格的正中间有一个障碍物。
左上角从右下角到一教育学习语言文字2条不同的路径：
1.向右 - >向右 - >向下 - >向下
2.向下 - >向下 - >向右 - >向右
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	//判断条件
	if len(obstacleGrid) == 0{
		return 0
	}
	m ,n := len(obstacleGrid) - 1,len(obstacleGrid[0]) - 1

	//起止点有障碍物
	if obstacleGrid[m][n] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	//采用动态规划
	var dp [][]int

	//初始化
	for i := 0; i < m;i++ {
		dp = append(dp,[]int{})
	}



	//开始求解
	for i := 0; i < m; i ++ {
		for j := 0; j < n;j++ {
			//判断障碍物
			if obstacleGrid[i][j] == 1 {
				dp[i] = append(dp[i],0)
				continue
			}

			//都为0
			if i == 0 && j == 0 {
				dp[i] = append(dp[i],1)
				continue
			}

			//判断边界
			if i == 0 {
				dp[i] = append(dp[i],dp[i][j - 1])
				continue
			}

			if j == 0 {
				dp[i] =  append(dp[i],dp[i - 1][j])
				continue
			}

			//其它
			dp[i] = append(dp[i],dp[i - 1][j] + dp[i][j - 1])
		}
	}

	//返回
	return dp[m - 1][n -1]
}


