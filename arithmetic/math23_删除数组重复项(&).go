package main

import "fmt"

/**
	算法描述:
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。
 */

 func main() {
 	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))
 }
/**
	消除重复项my
 */
func removeDuplicates(nums []int) int {
	//判断长度
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	length := len(nums) //长度

	for i := 0;i < length; i++{
		j := i + 1

		for j < length && nums[j] == nums[i] {
			j ++
		}

		//判断矢量
		move := j - i - 1

		//左移move
		for j < length {
			 nums[j - move] = nums[j]
			j ++
		}

		fmt.Println(nums)

		length -= move
	}

	return length
}
/**
	优化
 */
func removeDuplicates1(nums []int) int {
	//判断长度
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	length := len(nums) //长度

	i := 0  //基于栈实现的

	for j := 1; j < length;j ++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1


}