package main

import "sort"

/**
	算法描述
	在一个整数数组中，“峰”是大于或等于相邻整数的元素，相应地，“谷”是小于或等于相邻整数的元素。例如，在数组{5, 8, 6, 2, 3, 4, 6}中，{8, 6}是峰， {5, 2}是谷。现在给定一个整数数组，将该数组按峰与谷的交替顺序排序。

示例:

输入: [5, 3, 1, 2, 3]
输出: [5, 1, 3, 2, 3]
提示：

nums.length <= 10000
 */
func wiggleSort(nums []int)  {
	//判断
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	//声明临时数组
	data := make([]int,len(nums))
	copy(data,nums)

	//排序
	sort.Ints(data)

	//构造
	s,e := 0,len(data) - 1
	i := 0
	for s <= e {
		nums[i] = data[e]
		i ++
		if s != e {
			nums[i] = data[s]
			i++
		}
		s ++
		e --
	}
}