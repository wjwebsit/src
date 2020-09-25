package main

import (
	"fmt"
	"strconv"
)

/**
	算法描述
	定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10n 。

示例:

输入: 2
输出: 91
解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-numbers-with-unique-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func countNumbersWithUniqueDigits(n int) int {
	var dp = make(map[string]int)

	//位数
	var digit = make([]int,n + 1)
	digit[len(digit) - 1] = 1

	//记录之前填过的数字
	mp := make(map[int]int)
	for i := 0; i <= n; i ++ {
		mp[i] = -1
	}



	return dfsCountNumbers(n,true,true,digit,mp,dp)

}
/**
	自顶向下动态规划
 */
func dfsCountNumbers(cur int ,limit bool,zero bool,digit []int,cache map[int]int,dp map[string]int) int {
	//判断位数
	if cur < 0 {
		return 1
	}
	if !limit && !zero {
		str := strconv.Itoa(cur)
		for i := 0; i < 10;i ++ {
			if cache[i] == 1 {
				str += strconv.Itoa(i)
			}
		}
		if _,ok := dp[str];ok {
			return dp[str]
		}
	}

	//判断界限
	max := 9
	if limit == true {
		max = digit[cur]
	}
	ans := 0
	//寻找
	for i := 0 ;i <= max; i ++ {
		//不相同
		if cache[i] == 1 && i != 0{
			continue
		}

		//去除前导0
		if i == 0 && !zero && cache[i] == 1 {
			continue
		}

		//统计
		cache[i] = 1
		ans += dfsCountNumbers(cur - 1,limit && i == digit[cur],zero && i == 0,digit,cache,dp)

		//回溯
		cache[i] = -1
	}

	//写入缓存
	if !limit && !zero {
		str := strconv.Itoa(cur)
		for i := 0; i < 10;i ++ {
			if cache[i] == 1 {
				str += strconv.Itoa(i)
			}
		}

		//写入
		dp[str] = ans
	}
	//返回
	return ans


}
func main() {
	fmt.Println(countNumbersWithUniqueDigits(4))
}