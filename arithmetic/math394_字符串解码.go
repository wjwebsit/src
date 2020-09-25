package main

import "fmt"

/**
	算法描述：
	给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
 */
//初始化数字map
var nums = map[byte]int{
	'0':0,
	'1':1,
	'2':2,
	'3':3,
	'4':4,
	'5':5,
	'6':6,
	'7':7,
	'8':8,
	'9':9,
}
func decodeString(s string) string {
	//判断
	if len(s) == 0 {
		return ""
	}

	//声明返回
	str := ""

	//声明括号
	fg := 0


	//循环字符串
	i := 0
	for i < len(s) {
		//判断是否位数字
		if nums[s[i]] >= 1 {
			//继续寻找数字---并转化成整数
			num := nums[s[i]]
			j := i + 1
			for j < len(s) && s[j] != '[' {
				num = num * 10 + nums[s[j]]
				//num *= 10 + nums[s[j]] ---这么写结果不对
				j++
			}
			fmt.Println(num)

			//此时是s[j]为'['---fg遇到左括号+1 ,右括号-1
			fg ++
			i = j

			//寻找最右边的]
			for fg != 0 {
				j ++
				if s[j] == '[' {
					fg ++
				}
				if s[j] == ']' {
					fg --
				}
			}

			//获取字符串
			res := decodeString(string(s[i + 1:j]))

			//重复解码
			for k := 1;k <= num; k ++ {
				str += res
			}

			//重置i
			i = j + 1

		} else {
			//字符串拼凑
			str += string(s[i])
			i++
		}

	}

	//返回
	return str
}
func main() {
	s :="2[20[bc]31[xy]]xd4[rt]"
	fmt.Println(decodeString(s))

}