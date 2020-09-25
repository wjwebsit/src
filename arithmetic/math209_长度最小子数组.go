package main

import "fmt"

/**
	算法描述
	给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。

示例: 

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。

 */
func minSubArrayLen(s int, nums []int) int {
	//判断
	if len(nums) == 0 {
		return 0
	}

	//采用栈---递增的
	var stack []int
	stack = append(stack,nums[0])
	sum := nums[0]
	if sum >= s {
		return  1
	}

	//定义返回值
	rt := len(nums) + 1

	for i := 1; i < len(nums); i ++ {
		sum += nums[i]
		//判断
		if sum >= s {
			//获取差
			diff := sum - s

			//从栈找到小于diff的最开始的那个
			j := len(stack) - 1
			for j >= 0 && stack[j] > diff {//改为二分查找---就好了(logn)
				j --
			}

			//判断
			var distance int
			if j >= 0 {
				distance = i - j

			} else {
				distance = i + 1
			}

			if distance < rt {
				rt = distance
			}

		}
		stack = append(stack,sum)
	}
	//没有找到
	if rt == len(nums) + 1 {
		rt = 0
	}
	//返回
	return rt


}
func main() {
	nums := []int{1,1}
	fmt.Println(minSubArrayLen(7,nums))
}