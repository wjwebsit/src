package main

import "fmt"

/**
	算法描述:
	给定一个Excel表格中的列名称，返回其相应的列序号。

例如，

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
示例 1:

输入: "A"
输出: 1
示例 2:

输入: "AB"
输出: 28
示例 3:

输入: "ZY"
输出: 701
 */
func titleToNumber(s string) int {
	//判断
	if len(s) == 0 {
		return 0
	}

	//声明返回
	result := 0
	carry := 1

	//倒叙遍历
	for i := len(s) - 1; i >= 0; i-- {
		//获取元素
		v := s[i]
		result += int(v - 'A' + 1) * carry
		carry *= 26

	}

	//返回结果
	return result
}

func main() {
	fmt.Println(titleToNumber("ZY"))
}