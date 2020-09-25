package main

import "sort"

/**
	算法描述:
	给定一个包含n个整数的数组  nums，判断  nums 中是否存在三个元素a，b，c，使得  a + b + c = 0？找出所有满足条件且不重复的三元组。

	注意：答案中不可以包含重复的三元组。

	例如，给定数组nums = [-1,0,1,2，-1，-4]，

	满足要求的三元组集合为：
	[
	  [-1,0,1]，
	  [-1，-1,2]
	]
 */

func main() {

}
/**
双指针解法---利用了排序
 */
func threeSum(nums []int) [][]int {
	//定义结果集
	var result [][]int

	//判断长度是否大于3
	if len(nums) < 3 {
		return result
	}

	//升序排列---关键
	sort.Ints(nums)

	//开始遍历i,j,k 三个同时
	for i := 0; i <= len(nums)-3; i++ {
		//第二个
		j := i + 1

		//第三个
		k := len(nums) - 1
		//[j:k] 构成的区间来找 num[j]+num[k] = -num[i]

		//判断后一个和之前遍历过如果相等跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//[j:k] 区间未结束
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 { //相等存放
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				//判断后面的j 和前面的是否相等相等则跳过
				for j < k && nums[j] == nums[j-1] {
					j++
				}

			} else if nums[i]+nums[j]+nums[k] > 0 { //超过值则k-- 最右边的值大往左值小
				k--
			} else {//反之则右移
				j++
			}
		}
	}

	return result
}
