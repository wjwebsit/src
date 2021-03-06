package main
/**
算法描述:
	编写一个程序，找出第 n 个丑数。

丑数就是质因数只包含 2, 3, 5 的正整数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:  

1 是丑数。
n 不超过1690。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ugly-number-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func nthUglyNumber(n int) int {
	//采用dp
	p1,p2,p3 := 0,0,0
	var dp []int
	dp = append(dp,1)

	//构造丑数
	for i := 1; i < n ; i ++ {
		//取最小
		min := minABC2(dp[p1]*2,dp[p2] * 3,dp[p3] * 5)

		//写入
		dp = append(dp,min)
		if min == dp[p1] * 2 {
			p1 ++
		}
		if min == dp[p2] * 3 {
			p2 ++
		}
		if min == dp[p3] * 5 {
			p3 ++
		}
	}

	//返回
	return dp[n - 1]
}

func minABC2(a ,b,c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		return c
	}
	return  min
}