package main

import "fmt"

/**
算法描述
	累加数是一个字符串，组成它的数字可以形成累加序列。

一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。

给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。

说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。

示例 1:

输入: "112358"
输出: true
解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
示例 2:

输入: "199100199"
输出: true
解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
进阶:
你如何处理一个溢出的过大的整数输入?
 */
func isAdditiveNumber(num string) bool {
	//必须包含3个数
	if len(num) < 3 {
		return false
	}

	//累加值
	var num1,num2 string

	//是否是累加数
	fg := false


	//开始遍历
	for i := 0 ; i <= len(num) /2 && !fg; i ++ {
		//如果第一位数为0只能取0不能继续取
		if num[i] == 0 && i > 0 {
			break
		}
		//num1 = string(num[0: i + 1])
		for j := i + 1; j < len(num) && !fg; j ++ {
			//同上
			if num[i + 1] == '0' && j > i + 1 {
				break
			}
			num1 = string(num[0: i + 1])
			num2 = string(num[i + 1: j + 1])

			k := j + 1
			res := addStr(num1,num2)

			for k + len(res) <= len(num) && string(num[k: k + len(res)]) == res {
				num1 = num2
				num2 = res
				res = addStr(num1,num2)
				k = k + len(num2)

				if k == len(num) {
					fg =  true
				}

			}
		}

	}

	return fg
}

func addStr(s1,s2 string) string {
	var  nums = map[int]byte{
		0:'0',
		1:'1',
		2:'2',
		3:'3',
		4:'4',
		5:'5',
		6:'6',
		7:'7',
		8:'8',
		9:'9',
	}
	//字符串相加
	var result []byte

	//声明进位
	carry := 0

	//双指针
	i,j := len(s1) - 1,len(s2) - 1

	for i >= 0 && j >= 0 {
		//计算位数
		sum := carry + int(s1[i] - '0') + int(s2[j] - '0')

		//进位
		carry = sum / 10

		//写入
		result = append([]byte{nums[sum % 10]},result...)

		//next
		i--
		j--

	}

	//计算最后
	if i >= 0 {
		for i >= 0 {
			//计算位数
			sum := carry + int(s1[i] - '0')

			//进位
			carry = sum / 10

			//写入
			result = append([]byte{nums[sum % 10]},result...)

			//next
			i--


		}


	}

	if j >= 0 {
		for  j >= 0 {
			//计算位数
			sum := carry + int(s2[j] - '0')

			//进位
			carry = sum / 10

			//写入
			result = append([]byte{nums[sum % 10]},result...)

			//next
			j--

		}
	}
	if carry > 0 {
		result = append([]byte{nums[carry]},result...)
	}

	//返回
	return string(result[0:])


}
func main() {
	s := "198019823962"
	fmt.Println(isAdditiveNumber(s))
}