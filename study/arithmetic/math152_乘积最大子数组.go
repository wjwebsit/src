package main

import "fmt"

/**
	算法描述:
	给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字）。

 

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 */
func maxProduct(nums []int) int {
	//定义最大和最小dp
	var dp_Max  = make([]int,len(nums) + 1)
	var dp_Min  = make([]int,len(nums) + 1)

	//初始化
	dp_Max[0] = 1
	dp_Min[0] = 1

	for i := 1; i <= len(nums);i++ {
		if nums[i - 1] >= 0 {
			//判断当前值
			dp_Max[i] = maxInt1(dp_Max[i-1]*nums[i-1], nums[i-1])
			dp_Min[i] = minInt(dp_Min[i-1]*nums[i-1], nums[i-1])
		} else {
			dp_Max[i] = maxInt1(dp_Min[i-1]*nums[i-1], nums[i-1])
			dp_Min[i] = minInt(dp_Max[i-1]*nums[i-1], nums[i-1])
		}
	}

	//最大值
	max := dp_Max[1]
	for i := 2; i < len(dp_Max); i ++ {
		if max < dp_Max[i] {
			max = dp_Max[i]
		}
	}

	//返回
	return max
}
func maxInt1(a,b int) int {
	if a > b {
		return  a
	}
	return b
}
func minInt(a,b int) int {
	if a < b {
		return  a
	}
	return  b
}

func main() {
	nums := []int{-2,0,-1,9,-2}
	fmt.Println(maxProduct(nums))
}
