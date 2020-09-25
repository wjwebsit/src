package main

import "math"

/**
	算法描述:
	在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func maximalSquare(matrix [][]byte) int {
	//判断特殊情况
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	//获取矩阵的行和列
	rows,cols := len(matrix),len(matrix[0])

	//采用动态规划
	var dp = make([][]int,rows)

	//初始化
	for i := 0 ; i < rows; i ++ {
		dp[i] = make([]int,cols)
	}

	//定义最大面积
	max := 0

	//第一行第一列
	for col := 0;col < cols;col ++ {
		dp[0][col] = int(matrix[0][col] - '0')
		if dp[0][col] > max {
			max = dp[0][col]
		}

	}
	for row := 0;row < rows;row ++ {
		dp[row][0] = int(matrix[row][0] - '0')
		if dp[row][0] > max {
			max = dp[row][0]
		}
	}

	//中间部分
	for row := 1;row < rows;row ++ {
		for col := 1;col < cols;col ++ {
			//判断当前
			if matrix[row][col] == '0' {
				dp[row][col] = 0

			} else {//当前为1
				//获取上，左 和左上角的最小值
				min := minABC(dp[row - 1][col],dp[row][col - 1],dp[row - 1][col - 1])

				//开方
				min = int(math.Sqrt(float64(min)))

				//写入面积
				dp[row][col] = (min + 1) * (min + 1)

				//判断
				if dp[row][col] > max {
					max = dp[row][col]
				}
			}
		}
	}

	//返回最大面积
	return max
}

/**
	三数取最小
 */
func minABC(a,b,c int) int {
	tmp := a
	if tmp < b {
		tmp = b
	}
	if tmp < c {
		return c
	}
	return tmp
}