package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
	实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

示例 1:

输入: "1 + 1"
输出: 2
示例 2:

输入: " 2-1 + 2 "
输出: 3
示例 3:

输入: "(1+(4+5+2)-3)+(6+8)"
输出: 23
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。
 */

func calculate(s string) int {
	//获取后缀表达式
	formula := getFormula(s)

	//后缀表达式求值
	return myCalculate(formula)
}

func getFormula(s string)[]string {
	//声明操作符优先级
	operator := make(map[byte]int)
	operator['+'] = 1
	operator['-'] = 1
	operator['*'] = 2
	operator['/'] = 2
	operator['('] = 4
	operator[')'] = 4

	//去除空格
	s = strings.Replace(s," ","",-1)

	//声明栈
	var stack []byte

	//存储后缀表达式
	var formula []string

	//遍历s
	var i = 0
	for i < len(s) {
		//判断是否为数字
		if operator[s[i]] == 0 {
			j := i + 1
			for j < len(s) && operator[s[j]] == 0 {
				j ++
			}

			//存入
			formula = append(formula,string(s[i : j]))

			//赋值
			i = j
			continue
		}

		//判断符号栈是否为空
		if len(stack) == 0 || s[i] == '(' {
			stack = append(stack,s[i])
			i ++
			continue
		}

		//判断当前是否为右括号
		if s[i] == ')' {//出栈知道碰到'('
			for len(stack) > 0 && stack[len(stack) - 1] != '(' {
				//出栈并写入公式
				formula = append(formula,string(stack[len(stack) - 1]))

				//出栈
				stack = stack[0 :len(stack) - 1]
			}

			//左括号出栈
			stack = stack[0 :len(stack) - 1]

			//继续
			i++

		} else if operator[stack[len(stack) - 1]] >= operator[s[i]] {//栈顶元素的优先级比要入的操作符优先级低
			//出栈知道遇到一个比当前操作符优先级大的
			for len(stack) > 0 && operator[stack[len(stack) - 1]] >= operator[s[i]] && stack[len(stack) - 1] != '(' {
				//出栈并写入公式
				formula = append(formula,string(stack[len(stack) - 1]))

				//出栈
				stack = stack[0 :len(stack) - 1]
			}

			//当前入栈
			stack = append(stack,s[i])

			//继续
			i++

		} else {
			//入栈
			stack = append(stack,s[i])
			i++
		}
	}

	//剩余的写入
	for len(stack) > 0 {
		formula = append(formula,string(stack[len(stack) - 1]))
		stack = stack[0:len(stack) - 1]
	}

	//返回后缀表达式数组
	return formula

}
func myCalculate(formula []string) int {
	//后缀表达式
	funcHash := make(map[string]func(int,int)int)
	funcHash["+"] = func(a,b int) int {//+
		return a + b
	}
	funcHash["-"] = func(a,b int) int {//-
		return a - b
	}
	funcHash["*"] = func(a,b int) int {//*
		return a * b
	}
	funcHash["/"] = func(a,b int) int {///
		return a / b
	}

	//结果数组
	var nums []int

	//遍历后缀表达式
	for i := 0 ; i < len(formula);i++ {
		//判断
		if _,ok := funcHash[formula[i]]; !ok {//表示为数字

			//转为整数
			num,_ := strconv.Atoi(formula[i])

			//写入
			nums = append(nums,num)

		} else {
			//取数字
			num1 := nums[len(nums) - 1]
			num2 := nums[len(nums) - 2]

			//弹出
			nums = nums[:len(nums) - 1]

			//计算
			nums[len(nums) - 1] = funcHash[formula[i]](num2,num1)
		}
	}

	//返回
	return nums[0]
}
func main() {
	s := "2-1 + 2"
	fmt.Println(calculate(s))

}