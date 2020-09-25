package main

import "fmt"

/**
	算法描述:
	给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
 

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
 */
func minDistance(word1 string, word2 string) int {
	//采用动态规划
	var dp [][]int

	//获取两个单词的长度
	m,n := len(word1),len(word2)

	//初始化dp
	for i := 0; i <= n;i++ {
		dp = append(dp,make([]int,m + 1))
	}

	//初始化第一行第一列
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[0][i] = dp[0][i - 1] + 1
	}
	for j := 1;j <= n;j ++ {
		dp[j][0] = dp[j - 1][0] + 1
	}


	//中间部分
	for i := 1;i <= n;i ++ {
		for j := 1;j <= m;j++ {
			//判断是否相等
			if word1[j - 1] == word2[i - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] =  minDisT(dp[i - 1][j - 1],dp[i - 1][j],dp[i][j - 1]) + 1
			}

		}
	}

	//返回
	return dp[n][m]
}
func minDisT(a,b,c int) int {
	min := a;
	if a > b {
		min = b
	}
	if min > c {
		return c
	} else {
		return min
	}
}

func main() {
	 w1 := "intention"
	 w2 := "execution"
	fmt.Println(minDistance(w1,w2))
}