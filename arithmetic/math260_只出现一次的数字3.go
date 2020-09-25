package main
/**
	算法描述
	给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

示例 :

输入: [1,2,1,3,2,5]
输出: [3,5]
注意：

结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

 */
func singleNumber3(nums []int) []int {
	//计算不同数字的异或
	result := 0
	for i := 0 ; i < len(nums); i ++ {
		result ^= nums[i]
	}

	//获取最右边不同的位
	right := 1
	for (right & result) == 0 {
		right <<= 1
	}

	//分组
	num1 , num2 := 0, 0

	for _,v := range nums {
		if v & right == 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	return []int{num1,num2}
}