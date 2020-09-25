package main

import "strconv"

/**
	算法描述
	给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
 */
func addStrings(num1 string, num2 string) string {
	//判断特殊情况
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}

	//获取长度
	len1 := len(num1)
	len2 := len(num2)

	//前位补0
	if len1 < len2 {
		str := ""
		for i := len1 + 1; i <= len2 ;i ++ {
			str += "0"
		}
		num1 = str + num1
	}
	if len2  < len1 {
		str := ""
		for i := len2 + 1; i <= len1;i ++ {
			str += "0"
		}
		num2 = str + num2
	}

	//声明进位
	carry := 0

	//声明结果
	result := ""

	//开始相加
	for i := len(num1) - 1; i >= 0 ; i -- {
		//获取进位
		sum := carry + int(num1[i] - '0') + int(num2[i] - '0')

		//进位
		carry = sum / 10

		//当前
		result = strconv.Itoa(sum % 10) + result

	}

	//判断
	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}

	//返回
	return result

}