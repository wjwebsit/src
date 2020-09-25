package main

import "fmt"

/**
	算法描述
	给定一个整数数组 nums，返回区间和在 [lower, upper] 之间的个数，包含 lower 和 upper。
区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。

说明:
最直观的算法复杂度是 O(n2) ，请在此基础上优化你的算法。

示例:

输入: nums = [-2,5,-1], lower = -2, upper = 2,
输出: 3
解释: 3个区间分别是: [0,0], [2,2], [0,2]，它们表示的和分别为: -2, -1, 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-of-range-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	//统计
	count := 0

	//定义hash
	hash := make(map[int]int)
	hash[0] = 1 //代表自身
	sum := 0

	for i := 0; i < len(nums); i ++ {
		//计算dp
		sum += nums[i]

		//判断--hash存在区间值---采用二分查找
		for k,v := range hash {
			if sum -upper <= k && sum -lower >= k {
				count += v
			}
		}
		hash[sum] ++

	}

	return count

}
func main() {
	nums := []int{-1,1}
	fmt.Println(countRangeSum(nums,0,0))
}