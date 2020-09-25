package main

import "fmt"

/**
	算法描述
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
 */
func climbStairs(n int) int {
	//判断
	if n <= 0 {
		return  0
	}

	//定义决策
	choice := []int{1,2}
	rtResult := 0

	//添加缓存
	mp := map[int]int{}


	//回溯求解
	rtResult = f8(n,choice,mp)

	//返回
	return rtResult

}

/**
	回溯
 */
 func f8 (last int,choice []int,mp map[int]int) int {
	//判断
	if last < 0 {
		//返回0
		return 0
	}

	//判断
	if mp[last] > 0 {
		return mp[last]
	}

	//判断
	if last == 0 {
		return 1
	}


	//抉择
	sum := 0
	for _,v := range choice {
		//判断
		if v <= last {
			sum += f8(last - v,choice,mp)
		}
	}

	//写入缓存
	mp[last] = sum


	//返回
	return sum
 }
 /**
 	动态规划
  */
func climbStairs2(n int) int {
	//判断
	if n <= 0 {
		return  0
	}
	var dp []int
	dp = append(dp,1) //dp[0] = 1

	for i := 1;i <= n ; i ++ {
		sum := 0
		for j := 1; j <= i && j <=2 ; j ++ {
			sum += dp[i - j]
		}
		dp = append(dp,sum)
	}

	//返回
	return dp[n]

}

 //测试
 func main() {
 	fmt.Println(climbStairs(5))
 	fmt.Println(climbStairs2(5))

 }

