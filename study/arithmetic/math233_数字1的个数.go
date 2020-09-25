package main

import "fmt"

/**
	算法描述
	给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。

示例:

输入: 13
输出: 6
解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。

 */
/**
	数位dp
 */
func countDigitOne(n int) int {
	//判断
	if n <= 0 {
		return 0
	}

	//获取digit
	var digit []int
	for n > 0 {
		digit = append(digit,n % 10)
		n /= 10
	}
	dp  := [20][20]int{}
	return dfsCount(len(digit) - 1,true,0,digit,&dp)

}

func dfsCount(cur int,limit bool,pre int ,digit []int,cache *[20][20]int) int {
	//判断
	if cur < 0 {
		return pre
	}

	//判断缓存
	if (*cache)[cur][pre] > 0 && !limit {
		return (*cache)[cur][pre]
	}

	//获取
	max := 9
	if limit == true {
		max = digit[cur]
	}

	//统计
	ans := 0

	//遍历
	for i := 0; i <= max;i ++ {
		//判断
		if i == 1 {
			ans += dfsCount(cur - 1,limit && i == digit[cur],pre + 1,digit,cache)

		} else {
			ans += dfsCount(cur - 1,limit && i == digit[cur],pre,digit,cache)
		}
	}
	//存入cache
	if !limit {
		(*cache)[cur][pre] = ans

	}
	//返回
	return ans
}

func main() {
	fmt.Println(countDigitOne(200))
}