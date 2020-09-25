package main
/**
	算法描述
	设计一个算法，算出 n 阶乘有多少个尾随零。

示例 1:

输入: 3
输出: 0
解释: 3! = 6, 尾数中没有零。
示例 2:

输入: 5
输出: 1
解释: 5! = 120, 尾数中有 1 个零
 */
func trailingZeroes(n int) int {
	//判断1 - n 有多少5
	if n < 5 {
		return 0
	}
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}