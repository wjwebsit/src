package main

import (
	"fmt"
	"math"
	"sort"
)

/**
	算法描述
	给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

 

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
 */
/**
	完全背包问题
 */
func coinChange(coins []int, amount int) int {
	//判断
	if len(coins) == 0 ||amount < 0 {
		return - 1
	}

	//排序--用于剪枝
	sort.Ints(coins)

	var dp = make([]int,amount + 1)
	dp[0] = 0

	for i := 1; i <= amount; i ++ {
		//找零
		min := math.MaxInt64
		for j := 0; j < len(coins) && coins[j] <= i; j ++ {
			//判断
			if dp[i - coins[j]] == math.MaxInt64 {//找不开
				continue
			}
			//比较
			min = minCoin(min, 1 + dp[i - coins[j]])

		}
		//写入
		dp[i] = min
	}
	//返回
	if  dp[amount] == math.MaxInt64 {
		return  -1
	}
	return dp[amount]

}
func minCoin(a,b int) int{
	if a < b {
		return  a
	}
	return  b
}
func main() {
	coins:= []int{474,83,404,3}
	fmt.Println(coinChange(coins,264))
}
