package main

import "fmt"

/**
	算法描述
	给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

 

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2
 */
func majorityElement(nums []int) int {
	//获取长度
	n := len(nums)

	//统计
	c := n >> 1

	//构造map
	mp := make(map[int]int)
	rt := -1

	for i := 0; i < n;i++ {
		//统计
		if mp[nums[i]] + 1 > c {
			rt = nums[i]
			break
		} else {
			mp[nums[i]] ++
		}
	}

	//返回
	return  rt
}
func main() {
	nums := []int{2,2,1,1,1,2,2}
	fmt.Println(majorityElement(nums))
}