package main

import "fmt"

/**
	题目描述
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

 */
func isPalindrome1(s string) bool {
	//需要字符过滤和大小写 可以利用散列 一招解决 即 mp['a'] = 1 mp['A'] = 1 其它的为字母+1 mp[其它字符] = 0
	mp := map[byte]int{
		'a':1,
		'A':1,
		'B':2,
		'b':2,
		'C':3,
		'c':3,
		'D':4,
		'd':4,
		'E':5,
		'e':5,
		'F':6,
		'f':6,
		'G':7,
		'g':7,
		'H':8,
		'h':8,
		'I':9,
		'i':9,
		'J':10,
		'j':10,
		'K':11,
		'k':11,
		'L':12,
		'l':12,
		'M':13,
		'm':13,
		'N':14,
		'n':14,
		'O':15,
		'o':15,
		'P':16,
		'p':16,
		'Q':17,
		'q':17,
		'R':18,
		'r':18,
		'S':19,
		's':19,
		'T':20,
		't':20,
		'U':21,
		'u':21,
		'V':22,
		'v':22,
		'W':23,
		'w':23,
		'X':24,
		'x':24,
		'Y':25,
		'y':25,
		'Z':26,
		'z':26,
		'0':27,
		'1':28,
		'2':29,
		'3':30,
		'4':31,
		'5':32,
		'6':33,
		'7':34,
		'8':35,
		'9':40,

	}

	//定义返回值默是
	isPal := true

	//利用双指针
	start,end := 0,len(s) - 1

	//循环
	for start <= end {
		//判断start是否为字母
		if mp[s[start]] == 0 {
			start ++
			continue
		}

		//判断end是否为字母
		if mp[s[end]] == 0 {
			end --
			continue
		}

		//比较
		if mp[s[start]] != mp[s[end]] {
			isPal = false
			break
		}

		//表示相等向中间推进
		start ++
		end --
	}

	//返回
	return isPal
}

func main() {
	s := "OP"
	fmt.Println(isPalindrome1(s))
}