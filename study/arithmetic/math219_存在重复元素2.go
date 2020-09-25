package main
/**
	算法描述
	给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

 

示例 1:

输入: nums = [1,2,3,1], k = 3
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1
输出: true
示例 3:

输入: nums = [1,2,3,1,2,3], k = 2
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	//采用hash
	hash := make(map[int]int)

	//遍历
	for i := 0 ; i < len(nums); i ++ {
		//判断hash
		if _,ok := hash[nums[i]];ok {
			//计算
			if i - hash[nums[i]] <= k {
				return true
			}

		}

		//更新写入
		hash[nums[i]] = i

	}

	return false

}