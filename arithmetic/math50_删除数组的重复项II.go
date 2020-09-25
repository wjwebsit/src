package main

import "fmt"

/**
	算法描述
	给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定 nums = [1,1,1,2,2,3],

函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,1,2,3,3],

函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。

你不需要考虑数组中超出新长度后面的元素。
*/
func removeDuplicates2(nums []int) int {
	//判断
	if len(nums) == 0 {
		return 0
	}

	//双指针
	i := 0
	count := 1
	n := 2

	for j := 1; j < len(nums); j ++ {
		//判断
		if nums[j] != nums[i] {
			i ++
			nums[i] = nums[j]
			count = 1
			continue
		}

		//判断重复次数
		if count < n {
			i ++
			nums[i] = nums[j]
			count ++
		}

	}

	//返回
	return i + 1
}

/**
	基于栈
 */
func removeDuplicates3(nums []int,n int) []int {
	//判断
	if len(nums) <= n {
		return nums
	}

	if n == 0 {
		return []int{}
	}

	//声明栈
	var stack []int
	stack = append(stack,nums[0])

	//mp
	mp := make(map[int]int)
	mp[nums[0]]= 1

	for i := 1; i <len(nums); i ++ {
		//入栈
		if nums[i] != stack[len(stack) - 1] {
			stack = append(stack,nums[i])
			mp[nums[i]] ++
			continue
		}

		//相等但是没超过重复次数
		if mp[nums[i]] < n {
			stack = append(stack,nums[i])
			mp[nums[i]] ++
		}

	}

	//return
	return stack
}


//测试
func main() {
	arr := []int{1,1,1,2,2,3,3,3}

	fmt.Println(removeDuplicates2(arr))
}
