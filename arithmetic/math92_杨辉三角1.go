package main

import "fmt"

/**
	杨辉三角 不会百度
 */
func generate(numRows int) [][]int {
	//采用动态规划
	var dp [][]int

	//判断参数
	if numRows <= 0 {
		return dp
	}

	//初始化
	for i:= 0; i < numRows; i++ {
		dp = append(dp,[]int{})
	}
	dp[0] = append(dp[0],1)

	//判断行数
	if numRows >= 2 {
		//从第2行开始
		for i := 1; i < numRows; i ++ {
			//列
			for j := 0; j <= i; j++ {
				//判断是否为特殊
				if j == 0 || j == i {
					dp[i] = append(dp[i],1)
					continue
				}

				//否则
				dp[i] = append(dp[i],dp[i - 1][j - 1] + dp[i - 1][j])
			}
		}
	}

	//返回
	return dp

}
func main() {
	fmt.Println(generate(5))
}