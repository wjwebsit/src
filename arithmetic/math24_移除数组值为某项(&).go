package main

import "fmt"

/**
	算法描述
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:

给定 nums = [3,2,2,3], val = 3,

函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,1,2,2,3,0,4,2], val = 2,

函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。

注意这五个元素可为任意顺序。

你不需要考虑数组中超出新长度后面的元素。
 */

func removeElement(nums []int, val int) int {
	//判断特殊
	if len(nums) == 0{
		return 0
	}

	//首尾双指针类似快排序
	begin,end := 0,len(nums) - 1
	if begin == end && nums[begin] == val{
		return 0
	}

	count := 0

	for begin < end {
		//寻找val
		for begin< end  && nums[begin] != val {
			begin ++
		}

		for begin < end && nums[end] == val && nums[begin] == val{
			count++
			end --
		}

		//交换
		if begin <= end && nums[begin] == val {
			//交换
			nums[begin],nums[end] = nums[end],nums[begin]
			count ++
			end--
		}

	}
	fmt.Println(nums)
	//返回
	return len(nums) - count
}

func main() {
	nums := []int{3}
	fmt.Println(removeElement(nums,3))
}