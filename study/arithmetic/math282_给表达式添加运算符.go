package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
	算法描述
	给定一个仅包含数字 0-9 的字符串和一个目标值，在数字之间添加二元运算符（不是一元）+、- 或 * ，返回所有能够得到目标值的表达式。

示例 1:

输入: num = "123", target = 6
输出: ["1+2+3", "1*2*3"]
示例 2:

输入: num = "232", target = 8
输出: ["2*3+2", "2+3*2"]
示例 3:

输入: num = "105", target = 5
输出: ["1*0+5","10-5"]
示例 4:

输入: num = "00", target = 0
输出: ["0+0", "0-0", "0*0"]
示例 5:

输入: num = "3456237490", target = 9191
输出: []
 */
func addOperators(num string, target int) []string {
	var cache = make(map[string][]string)
	formulas := myAddOperators(num,0,len(num) - 1,target,&cache)

	//计算结果
	if len(formulas) == 0 {
		return []string{}
	}

	var result []string

	//计算
	for _,formula := range formulas {
		if calculate(formula) == target {
			result  = append(result,formula)
		}

	}

	return result


}

func myAddOperators(num string,s,e,target int,cache *map[string][]string) []string{
	//判断
	if s > e {
		return []string{}
	}

	//判断
	if s == e {
		return []string{string(num[s:e + 1])}
	}

	//判断缓存
	if _,ok := (*cache)[string(num[s:e + 1])];ok {
		return (*cache)[string(num[s:e + 1])]
	}


	//公式
	var formula []string
	mp := make(map[string]int)

	//遍历
	for i := s; i < e ; i ++ {
		var right []string

		//左半部部分结果
		left := string(num[s: i + 1])
		//123 == 1,12,123 but 012= 不行
		if num[s] == '0' && i != s{
			break
		}

		//右半部分结果
		right = myAddOperators(num, i+1, e, target, cache)

		if len(right) > 0 {
			//构造表达式
			for _, r := range right {
				if mp[left+"+"+r] != 1 {
					formula = append(formula, left+"+"+r)
					mp[left+"+"+r] = 1
				}
				if mp[left+"-"+r] != 1 {
					formula = append(formula, left+"-"+r)
					mp[left+"-"+r] = 1
				}
				if mp[left+"*"+r] != 1 {
					formula = append(formula, left+"*"+r)
					mp[left+"*"+r] = 1
				}

			}

		} else {
			formula = append(formula,left)
		}

	}
	if num[s] != '0' {
		formula = append(formula, num[s:e + 1])
	}

	//写入缓存
	(*cache)[string(num[s:e + 1])] = formula


	//返回
	return formula
}
func calculate(s string) int {

	stack := make([]int, 0, len(s))
	op := make([]int,0, len(s))
	num := 0
	for i:=0; i<len(s); i++ {
		switch s[i] {
		case '1','2','3','4','5','6','7','8','9','0':
			num = 0
			for i < len(s) && s[i]>='0'&&s[i]<='9'{
				num = num*10 + int(s[i]-'0')
				i++
			}
			if len(op)>0 && op[len(op)-1] > 2 {
				if op[len(op)-1] == 3 {
					stack[len(stack)-1] *= num
				}else {
					stack[len(stack)-1] /= num
				}
				op = op[:len(op)-1]
			}else {
				stack = append(stack, num)
			}
			// 退一格i 上面自动i++
			i--
		case '+':
			op = append(op, 1)
		case '-':
			op = append(op, -1)
		case '*':
			op = append(op, 3)
		case '/':
			op = append(op, 4)
		default:
			// fmt.Println(s[i])
		}
	}
	// fmt.Println(stack, op)
	for len(op) > 0 {
		stack[1] = stack[0] + op[0]*stack[1]
		op = op[1:]
		stack = stack[1:]
	}

	return stack[0]
}


func main() {
	s := "123"
	fmt.Println(addOperators(s,6))

}