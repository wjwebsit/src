package main

import (
	"fmt"
	"math"
)

/**
	算法描述：
	给定一个非负整数数组，你最初位于数组的第一个位置。

    数组中的每个元素代表你在该位置可以跳跃的最大长度。

   你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 */
/**
	dynamic 动态规划 带备忘录机制
 */
func jump(nums []int) int {
	//判断
	if len(nums) == 0 {
		return len(nums)
	}
	if len(nums) == 1 {
		return 0
	}
	//定义起始
	min := math.MaxInt64
	result := []int{}
	//初始化
	for i := 0; i < len(nums);i++ {
		result = append(result,0)
	}

	//跳跃
	min = f5 (nums,0,len(nums) -1,&result)

	//返回
	return -1 * min
}

func f5(nums []int,begin int,end int,result *[]int) int{
	//判断越界
	if begin >= end {//最后一次
		return 0
	}
	//判断缓存
	if len(*result) > 0 && (*result)[begin] != 0 {
		return (*result)[begin]
	}
	//定义最小
	min := math.MinInt64

	//获取当前可以跳跃的步骤
	length := nums[begin]
	//不能在跳
	if length <= 0 {
		return -1 * end + 1
	}

	//决策跳跃步数--如果最大就越界了则就是最小（贪心策略）
	for j := 1;j <= length;j++ {

		//处理子问题
		tmp := -1 + f5(nums,begin + j,end,result)

		//判断
		if tmp > min {
			min = tmp
		}

	}

	//写入缓存
	(*result)[begin] = min

	//返回
	return   min

}

func main() {
	arr := []int{2,3,1,1,4}

	fmt.Println(jump2(arr))
}

func jump2(nums []int) int {
	//判断
	if len(nums) == 0 {
		return len(nums)
	}
	if len(nums) == 1 {
		return 0
	}
	//步数
	var dp []int
	dp = append(dp,nums[0])
	step := 0

	for i := 1; i < len(nums); i++ {
		if dp[i - 1] < (nums[i] + i) {
			step ++
			dp = append(dp,nums[i] + i)
		} else {
			dp = append(dp,dp[i - 1])
		}

	}
	//返回
	return step


}
func maxJump2(a,b int ) int {
	if a > b {
		return a
	}
	return b
}
