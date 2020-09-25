package main

import "fmt"

/**
	算法描述
	给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

 

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
 

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
 */
func productExceptSelf(nums []int) []int {
	//声明返回值
	var result = make([]int,len(nums))

	//判断
	if len(nums) < 2 {
		return result
	}
	result[0] = 1
	//构造左区间
	for i := 1; i< len(nums); i ++ {
		result[i] = nums[i] * result[i - 1]
	}

	//构造返回数组
	r := 1
	for i := len(nums) - 1; i > 0 ; i -- {
		result[i] = r * result[i - 1]
		r *= nums[i]
	}
	result[0] = r

	//返回
	return result


}
func main() {
	nums := []int{1,2,3,4}
	fmt.Println(productExceptSelf(nums))
}