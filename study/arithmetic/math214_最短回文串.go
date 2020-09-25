package main
/**
	算法描述
	给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。

示例 1:

输入: "aacecaaa"
输出: "aaacecaaa"
示例 2:

输入: "abcd"
输出: "dcbabcd"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func shortestPalindrome(s string) string {
	//判断
	if len(s) == 0 {
		return ""
	}

	//获取反转字符串
	s2 := ""
	for i := len(s) - 1;i >= 0 ; i -- {
		s2 += string(s[i])
	}

	//倒叙s
	var i int
	for i = len(s) - 1;i >= 0 ;i ++ {
		//判断s[0,i]是否为回文串
		if string(s[0: i + 1]) == s2[0:i+1] {
			break
		}

	}

	return string(s2[0:i]) + s

}