package main

import (
	"bytes"
	"strconv"
	"strings"
)

/**
	算法描述
	给定一个整数，将其转化为7进制，并以字符串形式输出。

示例 1:

输入: 100
输出: "202"
示例 2:

输入: -7
输出: "-10"
 */
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var buffer []string
	offset := 1
	if num < 0 {
		num *= -1
		offset = -1

	}
	for num > 0 {
		//获取位数
		digit := num % 7
		buffer = append([]string{strconv.Itoa(digit)},buffer...)
		num /= 7
	}

	rtStr := strings.Join(buffer,"")
	if offset == -1 {
		return "-" + rtStr
	}
	return rtStr
}
