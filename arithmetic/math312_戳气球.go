package main

import (
	"fmt"
)

/**
	算法描述
	有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。

求所能获得硬币的最大数量。

说明:

你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
示例:

输入: [3,1,5,8]
输出: 167
解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
通过次数13,453提交次数22,205

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/burst-balloons
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func maxCoins(nums []int) int {
	//判断
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)

	//前后补1
	nums = append([]int{1},nums...)
	nums = append(nums,1)

	//初始化---只有一个元素
	var dp = make([][]int,len(nums))
	for i := 0 ; i < len(nums); i ++ {
		dp[i] = make([]int,len(nums))
		dp[i][i] = 0
	}
	for l := 1; l <= n; l ++ {//步长
		for i := 1; i <= n + 1 - l; i ++ {//步长范围内区间起位置
			j := i + l - 1
			for k := i; k <= j; k ++ {
				dp[i][j] = maxAB(dp[i][j],dp[i][k - 1] + dp[k + 1][j] + nums[i - 1] * nums[k] * nums[j + 1])
			}


		}

	}
	//返回
	return dp[1][n]


}
func myMaxCoins(nums []int,s,e int,dp *[50][50]int) int {
	//判断越界
	if s > e {
		return  0
	}

	//判断缓存
	if (*dp)[s][e] > 0 {
		return (*dp)[s][e]
	}

	//定义max
	max := 0

	left,right := 1,1
	if s - 1 >= 0 {
		left = nums[s - 1]
	}
	if e + 1 < len(nums) {
		right = nums[e + 1]
	}

	//分割
	for i:= s; i <= e; i ++ {
		max = maxAB(max,myMaxCoins(nums,s,i - 1,dp) + myMaxCoins(nums,i + 1,e,dp) + left * nums[i] * right)

	}

	//写入缓存
	(*dp)[s][e] = max

	//返回
	return (*dp)[s][e]

}

func maxAB(a,b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	nums := []int{3,1,5,8}
	fmt.Println(maxCoins(nums))
}