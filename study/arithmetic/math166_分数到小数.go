package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
	算法描述
	给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。

如果小数部分为循环小数，则将循环的部分括在括号内。

示例 1:

输入: numerator = 1, denominator = 2
输出: "0.5"
示例 2:

输入: numerator = 2, denominator = 1
输出: "2"
示例 3:

输入: numerator = 2, denominator = 3
输出: "0.(6)"

 */
func fractionToDecimal(numerator int, denominator int) string {
	//判断分子
	if numerator == 0 {
		return "0"
	}
	//是否为负数
	fg := false

	//是否为负数
	if (denominator < 0 && numerator > 0) || (numerator < 0 && denominator > 0) {
		fg = true
	}

	//转为正整数
	denominator = int(math.Abs(float64(denominator)))
	numerator = int(math.Abs(float64(numerator)))

	//获取整数部分
	inter := numerator / denominator

	//小数部分
	deci := numerator % denominator

	//判断小数
	if deci == 0 {
		if fg == true {
			return "-" + strconv.Itoa(inter)
		} else {
			return strconv.Itoa(inter)
		}
	}


	//定义stack---记录小数的除数顺序hash 记录商
	var result []int
	mp := make(map[int]int)
	for true {
		//获取当前小数写入
		result = append(result,deci)

		//获取商
		trade := (deci * 10) / denominator

		//写入hash
		mp[deci] = trade

		//获取下一次的除数
		deci = deci * 10 - denominator * trade

		//判断是否循环
		if deci == 0 || mp[deci] > 0 {
			break
		}

	}

	//判断有无循环
	//构造小数部分--字符串
	var deciStr = ""

	if deci == 0 {
		for _,v := range result {
			deciStr += strconv.Itoa(mp[v])
		}

	} else {//表示存在循环
		for _,v := range  result {
			//判断是否等于最后的deci
			if v == deci {
				deciStr += "(" + strconv.Itoa(mp[v])
			} else {
				deciStr += strconv.Itoa(mp[v])
			}

		}

		//最后的
		deciStr += ")"
	}

	//返回
	if !fg {
		return strconv.Itoa(inter) + "." + deciStr
	}

	//负数
	return  "-" + strconv.Itoa(inter) + "." + deciStr

}
func main() {
	fmt.Println(fractionToDecimal(-4,2))


}