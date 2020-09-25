package main

import "fmt"

/**
	算法描述
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
 */
func moveZeroes(nums []int)  {
	//判断数组
	if len(nums) == 0 {
		return
	}

	//寻找第一个0的位置
	var index int
	for index = 0; index < len(nums);index ++ {
		if nums[index] == 0 {
			break
		}
	}

	//声明队列
	count := 1
	for i := index + 1; i < len(nums); {
		//判断是否为0
		if nums[i] == 0 {
			//入队
			i++
			count ++
			continue
		}

		//双指针
		var j  = i
		for  j < len(nums) && nums[j] != 0 {
			//左移count位
			nums[j - count] = nums[j]
			j++
		}

		//重置i
		i = j
	}

	//最后置0
	l := len(nums) - 1
	for i := 0; i < count; i ++ {
		nums[l - i] = 0
	}
	fmt.Println(nums)
}
func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
}