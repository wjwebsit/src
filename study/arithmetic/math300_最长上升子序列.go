package main

import "fmt"

/**
	算法描述
	给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func lengthOfLIS(nums []int) int {
	//判断
	if len(nums) == 0 {
		return 0
	}

	//采用动态规划----倒叙统计后面比(i,len(nums))大的
	var dp  = make([][]int,len(nums))

	//统计
	dp[len(nums) - 1] = []int{}
	hash := make(map[int]int)
	hash[len(nums) - 1] = 1
	rtMax := 1

	//统计i后边有多少比i值大的
	for i := len(nums) - 2; i >= 0 ; i -- {
		max := 1
		for j := i + 1; j < len(nums); j ++ {
			if nums[j] > nums[i] {
				//判断
				tmpMax := 1 + hash[j]
				if tmpMax > max {
					max = tmpMax
				}
			}

		}
		hash[i] = max
		if rtMax < max {
			 rtMax = max
		}
	}

	//计算子序列
	/*hash := make(map[int]int)
	rtMax := 1

	for i := len(nums) - 1;i >= 0 ; i -- {
		//判断
		if len(dp[i]) == 0 {
			hash[i] = 1
		} else {
			//遍历
			max := 1
			for j := len(dp[i]) - 1;j >= 0;j -- {
				tmpMax := 1 + hash[dp[i][j]]
				if tmpMax > max{
					max = tmpMax
				}
			}
			hash[i] = max

		}
		if hash[i] > rtMax {
			rtMax = hash[i]
		}

	}*/

	//返回
	return rtMax

}
func main() {
	nums := []int{10,9,2,5,3,7,101,18}

	fmt.Println(lengthOfLIS(nums))
}