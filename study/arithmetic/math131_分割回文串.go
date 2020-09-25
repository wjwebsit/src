package main

import "fmt"

/**
	算法描述:
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func partition1(s string) [][]string {
	var dp = make([][]int,len(s))

	//长度为1，长度为2
	for i := 0 ; i < len(s);i ++ {
		//初始化dp[i]
		dp[i] = make([]int,len(s))
		dp[i][i] = 1

		//2个字符
		if i + 1 < len(s) && s[i + 1] == s[i] {
			dp[i][i + 1] = 1
		}
	}


	//利用动态规划确定所有的的回文串
	for i := len(s) - 3; i >= 0 ; i -- {
		for j := len(s) - 1; j > i + 2 ; j ++ {
			//判断
			if s[i] == s[j] && dp[i + 1][j - 1] == 1 {
				dp[i][j] = 1
			}

		}
	}


	//声明缓存
	cache := make([][][]string,len(s))

	return partitionHs(s,0,dp,&cache)

}

func partitionHs(s string,begin int,dp [][]int,cache *[][][]string) [][]string {
	//判断
	if begin >= len(s) {
		return [][]string{}
	}

	//判断缓存
	if len((*cache)[begin]) > 0 {
		return (*cache)[begin]
	}

	//声明结果
	var result [][]string

	//遍历
	for i := begin; i < len(s); i ++ {
		//判断是否为回文
		if dp[begin][i] == 1 {
			//获取后面的回文串
			tmp := partitionHs(s,i + 1,dp,cache)

			cur := string(s[begin:i + 1])
			//判断是否为空
			if len(tmp) == 0 {
				//当前就是
				result = append(result,[]string{cur})
			} else {

				//遍历并拼上当前
				for _,v := range tmp {
					//压入当前
					v = append([]string{cur},v...)

					//写入结果
					result = append(result,v)
				}

			}

		}

	}

	//写入缓存
	(*cache)[begin] = result

	//返回
	return result

}
/**
	判断是否为回文子串
 */
func isParam(str string,s,e int) bool {
	//双指针
	for s <= e {
		if str[s] != str[e] {
			return  false
		}
		s ++
		e --
	}
	return  true
}

func main()  {
	fmt.Println(partition1("aab"))
}