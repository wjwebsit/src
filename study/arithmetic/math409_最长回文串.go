package main
/**
	算法描述:
	给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 */
func longestPalindrome(s string) int {
	//判断是否为空
	if len(s) == 0 {
		return 0
	}

	//构造mp---统计出现次数
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] ++
	}

	//申明返回
	sum := 0
	for _,v := range mp {
		//判断是否为偶数
		if v & 1 == 0 {
			sum += v
		} else {
			sum += v - 1
		}
	}

	//返回
	return sum + 1

}