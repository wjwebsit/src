package main

import "fmt"

/**
	算法描述
	给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:

输入: 2
输出: [0,1,1]
示例 2:

输入: 5
输出: [0,1,1,2,1,2]
进阶:

给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
要求算法的空间复杂度为O(n)。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
 */

func countBits(num int) []int {
	//采用动态规划
	var dp []int
	dp = append(dp,0)

	//判断
	if num <= 0 {
		return dp
	}

	//carry表示2的幂
	dp = append(dp,1)

	//开始遍历num
	carry := 2
	for i := 2; i <= num; {
		var j int
		m := carry / 2
		for j = i; j <= num && j <= carry;j++ {
			dp = append(dp, 1 + dp[j  % m])
		}
		i = j
		carry <<=1
	}

	//返回
	return dp
}

func main()  {
	fmt.Println(countBits(5))
}