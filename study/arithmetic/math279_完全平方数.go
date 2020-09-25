package main
/**
	算法描述:
	给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.
 */
/**
	完全背包问题
 */
func numSquares(n int) int {
	//构造平方序列
	var squares []int
	for i := 1;i * i <= n; i++ {
		squares = append(squares,i * i)
	}
	//声明dp
	var dp  = make([]int,n + 1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i ++ {
		//可以决策的
		dp[i] = i //最大就是i
		for j := 0;j < len(squares) && squares[j] <= i; j ++ {
			dp[i] = minNum(dp[i],1 + dp[i - squares[j]])

		}

	}
	//返回
	return  dp[n]
}
func minNum(a,b int ) int{
	if a < b {
		return a
	}
	return b
}
