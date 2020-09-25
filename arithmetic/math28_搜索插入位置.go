package main

import "fmt"

/**
	算法描述
	给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
 */

func searchInsert(nums []int, target int) int {
	//判断长度
	if len(nums) == 0 {
		return 0
	}

	//判断1
	if len(nums) == 1 {
		if nums[0] >= target {
			return  0
		} else if nums[0] < target {
			return 1
		}
	}

	//折半查找索引
	begin,end := 0,len(nums) - 1
	index := -1

	for begin <= end {
		mid := (begin + end) / 2 //向下取整

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			begin = mid + 1
		} else {
			index = mid
			break
		}

	}

	//判断是否存在
	if index != -1 {
		return index
	}

	//不存在位置
	return end + 1
}

func main() {
	arr := []int{1,2,3,4,6}
	fmt.Println(searchInsert(arr,5))

}