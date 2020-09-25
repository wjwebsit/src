package main

/**
	算法描述
	给定一个整数 n，返回 n! 结果尾数中零的数量。

示例 1:

输入: 3
输出: 0
解释: 3! = 6, 尾数中没有零。
示例 2:

输入: 5
输出: 1
解释: 5! = 120, 尾数中有 1 个零.

 */
func trailingZeroes(n int) int {
	//声明返回
	count := 0

	//统计5的出现的次数
	for n > 0 {
		count += n / 5
		n = n / 5
	}

	return  count
}