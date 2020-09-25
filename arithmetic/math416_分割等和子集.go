package main

import "fmt"

/**
	算法描述
	给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].
 

示例 2:

输入: [1, 2, 3, 5]

输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func canPartition(nums []int) bool {
	//判断数组
	if len(nums) == 0 {
		return false
	}

	//计算和
	sum := 0
	for i := 0; i < len(nums);i ++ {
		sum += nums[i]
	}

	//如果为奇数则不可以
	if sum & 1 == 1 {
		return false
	}

	//获取一半和
	sum = sum / 2

	//0-1背包
	//声明dp
	var dp  = make([]bool,sum + 1)

	for i := 1; i < len(nums); i ++ {
		for j := sum;j >= 1;j -- {
			if nums[i - 1] == j {
				dp[j] = true
			} else {
				if nums[i-1] < j {
					dp[j] = dp[j] || dp[j-nums[i-1]]
				}
			}
		}
	}

	return dp[sum]
}
func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}