package main

/**
	给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
说明:

你可以假设数组不可变。
会多次调用 sumRange 方法。
 */
type NumArray struct {
	sums []int
	data []int
}


func Constructor2(nums []int) NumArray {
	//遍历数组
	arr := new(NumArray)
	arr.sums = append(arr.sums,nums[0])
	for i := 1; i < len(nums);i ++ {
		arr.sums = append(arr.sums,arr.sums[i - 1] + nums[i])
	}

	arr.data = make([]int,len(nums))
	copy(arr.data,nums)

	//返回
	return *arr


}


func (this *NumArray) SumRange(i int, j int) int {
	//判断
	if i == j {
		return this.data[i]
	}

	//判断0
	if i == 0 {
		return this.sums[j]
	}

	return this.sums[j] - this.sums[i - 1]

}