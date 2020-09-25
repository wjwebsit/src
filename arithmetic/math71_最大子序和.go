package main

import "fmt"

/**
	算法描述
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 */

func maxSubArray(nums []int) int {
	var dp []int
	dp = append(dp,nums[0])
	max := nums[0]
	for i := 1;i < len(nums);i++ {
		tmp := maxSub(nums[i],nums[i] + dp[i - 1])
		dp = append(dp,tmp)
		max = maxSub(tmp,max)
	}

	return max
}

func maxSub(a,b int ) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	arr := []int {-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(arr))
}
