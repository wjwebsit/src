package main

import "fmt"

/**
算法描述
给定一个字符串 (s) 和一个字符模式 (p)。实现支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符。
'*' 匹配零个或多个前面的元素。
匹配应该覆盖整个字符串 (s) ，而不是部分字符串。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: '*' 代表可匹配零个或多个前面的元素, 即可以匹配 'a' 。因此, 重复 'a' 一次, 字符串可变为 "aa"。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个('*')任意字符('.')。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 'c' 可以不被重复, 'a' 可以被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
 */
func isMatch(s string, p string) bool {
	//初始化缓存dp[i][j]表示p[j:] 是否能匹配到 s[i:]
	var dp = make (map[int]map[int]bool,len(p))

	//return
	return isMatchHS(s,p,0,0,&dp)
}

func isMatchHS(s,p string,i,j int, dp *map[int]map[int]bool) bool {
	//判断
	if  j >= len(p) { //p结束，s是否结束
		if i >= len(s) { //表示匹配了
			return true
		}

		//否则表示没有匹配
		return false
	}

	//判断缓存
	if _,ok := (*dp)[i][j];ok {
		return (*dp)[i][j]
	}

	//开头是否匹配
	first := false

	//判断开头是否匹配
	if i < len(s) && (p[j] == s[i] || p[j] == '.')  {
		first = true

	}

	//判断p的长度
	if  j + 1 < len(p) && p[j + 1] == '*'{
		//去掉或者重读 p[j]
		first = isMatchHS(s,p,i,j + 2,dp) || (first && isMatchHS(s,p,i + 1,j ,dp))

	} else {
		first = first && isMatchHS(s,p,i + 1,j + 1,dp)
	}


	//写入缓存
	if (*dp)[i] == nil {
		(*dp)[i] = make(map[int]bool)
	}
	(*dp)[i][j] = first

	//返回
	return first

}

