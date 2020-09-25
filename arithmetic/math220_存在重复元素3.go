package main

import (
	"math"
	"sort"
)

/**
	算法描述
	给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j 之间的差的绝对值最大为 ķ。

示例 1:

输入: nums = [1,2,3,1], k = 3, t = 0
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1, t = 2
输出: true
示例 3:

输入: nums = [1,5,9,1,5,9], k = 2, t = 3
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	//判断特殊
	if len(nums) == 0 {
		return false
	}

	//采用滑动窗口
	left,right := 0,0

	for right < len(nums) {
		//判断
		for i := left;i <= right;i++ {
			//判断
			if math.Abs(float64(nums[right] - nums[i])) <= float64(t) {
				return true
			}

		}

		//判断窗口大小
		if right - left == k {
			left ++
		}
		right ++

	}
	return false

}