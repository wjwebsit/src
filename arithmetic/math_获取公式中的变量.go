package main

import (
	"fmt"
	"strings"
)

func main() {
	formula := "2+4/7";
	fmt.Println(getVariables(formula))
}
func getVariables(formula string) []string {
	//声明返回
	var result []string

	//判断公式
	if len(formula) == 0 {
		return  result
	}

	//去除空格
	formula = strings.Replace(formula," ","",-1)

	//构造操作符
	operator := map[byte]bool {
		'+':true,
		'-':true,
		'*':true,
		'/':true,
		'(':true,
		')':true,
		'@':true,//用作哨兵
	}

	//i定义
	i := 0

	//定义合法
	valid := true

	//定义变量是否存在
	exist := make(map[string]bool)

	for i < len(formula) {
		//是否为@
		if formula[i] == '@' {
			//j 从i + 1开始
			j := i + 1
			for j < len(formula) && operator[formula[j]] == false {
				j++
			}

			//判断结束是否为哨兵
			if j < len(formula) && formula[j] == '@' {
				//不合法
				valid = false
				break
			}

			//判断是否为5 + @ 这种情况
			if j - i == 1 {
				valid = false
				break
			}

			//写入结果集
			if exist[string(formula[i+1:j])] == false {
				result = append(result, string(formula[i+1:j]))
				exist[string(formula[i+1:j])] = true
			}

			//重置i为 j + 1
			i = j + 1

		} else {//不为@ i右移1位
			i++
		}

	}

	//判断是否合法
	if valid == false {
		return []string{}
	}

	//返回
	return result

}

func canJump2(nums []int) bool {
	//获取开始的跳跃距离
	max := -1
	i := 0
	for i <= max{
		//寻找最大
		tmpMax := i + nums[i]
		if tmpMax > max {
			max = tmpMax
		}
		i++
	}

	//判断
	return  max >= len(nums) - 1

}

