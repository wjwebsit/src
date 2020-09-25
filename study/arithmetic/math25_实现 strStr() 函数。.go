package main

/**
	算法描述：
	给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
 */

/**
	典型的KMP模式匹配法
 */
func strStr(haystack string, needle string) int {
	//判断是否为空
	if len(needle) == 0 {
		return 0
	}

	//定义下标
	i,j,rtIndex := 0,0,-1

	//获取next
	next := getStrNextArray(needle)

	//遍历条件为都未达到瓶颈
	for i < len(haystack) && j < len(needle) {
		//判断是否相等
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++

		} else {

			j = next[j]

		}

	}

	//判断j 最后是否走完 走完说明完全匹配
	if j == len(needle) {
		rtIndex = i - j //最开始位置
	}

	//返回结果
	return rtIndex
}
/**
	寻找next数组
 */
func getStrNextArray(s string) []int {
	//声明next数组
	var next []int

	//初始化
	for i := 0; i < len(s);i++ {
		next = append(next,0)
	}

	//next数组0为-1
	next[0] = -1


	//定义开始和结束
	begin,end := -1,0

	for end < len(s) - 1  {

		if begin == -1 || s[begin] == s[end] {
			begin++
			end++
			//next[end] = begin 原始的
			/***改进**/
			if s[begin] != s[end] {
				next[end] = begin
			} else {
				next[end] = next[begin]
			}

		} else {
			begin = next[begin]
		}

	}

	return next

}
