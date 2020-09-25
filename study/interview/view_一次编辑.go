package main
/**
	算法描述
	字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

 

示例 1:

输入:
first = "pale"
second = "ple"
输出: True
 

示例 2:

输入:
first = "pales"
second = "pal"
输出: False
 */
func oneEditAway(first string, second string) bool {
	//获取长度
	m,n := len(first),len(second)

	//构造dp
	var dp = make([][]int,n + 1)

	//初始化dp--并初始化第一列
	for i := 0 ; i < len(dp);i ++ {
		dp[i] = make([]int,m + 1)
		if i == 0 {
			dp[0][0] = 0
		} else {
			dp[i][0] = dp[i - 1][0] + 1
		}

	}

	//初始化第一行
	for i := 1 ; i < m + 1; i ++ {
		dp[0][i] = dp[0][i - 1] + 1
	}


	//中间部分
	for i := 1; i < n + 1; i ++ {
		for j := 1; j < m + 1; j ++ {
			min := minABC4(dp[i - 1][j - 1],dp[i - 1][j],dp[i][j - 1])
			if second[i] == first[j] {
				dp[i][j] = min
			} else {
				dp[i][j] = min + 1
			}

		}
	}

	//返回
	return dp[n][m] <= 1
}

func minABC4(a,b,c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return  min
}