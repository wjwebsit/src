package main

import "fmt"

/**
	算法描述:

 */
func calculateMinimumHP(dungeon [][]int) int {
	//采用动态规划
	var dp [][]int //所需hp
	//获取行和列
	rows,cols := len(dungeon),len(dungeon[0])


	//初始化dp 和 current
	for row := 0; row < rows;row ++ {
		dp = append(dp,make([]int,cols))
	}

	//从右上角往左上角dp
	if dungeon[rows - 1][cols - 1] < 0 { //表示可以到达当前右下角需要的体力
		dp[rows - 1][cols - 1] =  -1 * dungeon[rows - 1][cols - 1] + 1

	} else {
		dp[rows - 1][cols - 1] = 1

	}

	//初始化最后一行和最后一列
	for col := cols - 2; col >= 0;col-- {
		hp := dp[rows - 1][col + 1] - dungeon[rows - 1][col]
		if hp <= 0 {
			dp[rows - 1][col] = 1
		} else {
			dp[rows - 1][col] = hp
		}
	}
	for row := rows - 2;row >= 0; row -- {
		dp[row][cols - 1] = dp[row + 1][cols - 1] - dungeon[row][cols - 1]
		if dp[row][cols - 1] <= 0 {
			dp[row][cols - 1] = 1
		}
	}


	//中间部分
	for row := rows - 2;row >= 0;row -- {
		for col := cols - 2;col >= 0; col -- {
			//获取左下角和右下角的最小hp
			hp := minHp(dp[row + 1][col],dp[row][col + 1]) - dungeon[row][col]

			//判断
			if hp <= 0 {
				hp = 1
			}
			dp[row][col] = hp
		}
	}

	//返回
	return dp[0][0]
}
func minHp (a,b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	dungen := [][]int{
		[]int{1,-3,3},
		[]int{0,-2,0},
		[]int{-3,-3,-3},
	}
	fmt.Println(calculateMinimumHP(dungen))
}