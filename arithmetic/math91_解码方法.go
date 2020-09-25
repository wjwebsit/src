package main

import (
	"fmt"
	"strconv"
)

/**
	一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1:

输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2:

输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
 */
 /**
 	DFS + 记忆搜索
  */
func numDecodings(s string) int {
	//定义缓存
	cache := make(map[string]int)
	return numDecodingHS(s,0,cache)
}

func numDecodingHS(s string,start int,cache map[string]int) int {
	//判断
	if start == len(s) {
		return  1
	}

	//判断缓存
	if _,ok := cache[string(s[start:])];ok {
		return cache[string(s[start:])]
	}

	//声明返回值
	sum := 0

	//开始---因为字母只能涉及到26 故i < start + 2 即可
	for i := start; i < len(s) && i < start + 2; i ++ {
		//如果为0
		if s[start] == '0' {//无法解析
			break
		}
		//判断
		if i == start {//必定存在
			sum += numDecodingHS(s,i + 1,cache)
			continue
		}

		//2位时
		if num ,_ := strconv.Atoi(string(s[start : i + 1]));num > 26 {
			break
		}

		//继续
		sum += numDecodingHS(s,i + 1,cache)
	}

	//写入cache
	cache[string(s[start:])] = sum

	//返回
	return sum
}
/**
	采用动态规划----倒叙s
 */
 func numDecodingsDP(s string) int {
 	//声明dp
 	if len(s) == 0 {
 		return 0
	}
 	var dp  = make([]int,len(s) + 1)

 	//获取长度
 	n := len(s)

 	dp[n] = 1
 	if s[n - 1] != '0' {
		dp[n-1] = 1
	} else {
		dp[n - 1] = 0
	}
 	for i := n - 2 ; i >= 0; i -- {
 		//判断
 		if s[i] == '0' {
 			dp[i] = 0
 			continue
		}
 		if num,_ := strconv.Atoi(string(s[i: i + 2])); num > 26 {
 			dp[i] = dp[i + 1]
		} else {
			dp[i] = dp[i + 1] + dp[i + 2]
		}
	}

 	//返会
 	return dp[0]
 }


func main() {
	s := "226"
	fmt.Println(numDecodingsDP(s))
}