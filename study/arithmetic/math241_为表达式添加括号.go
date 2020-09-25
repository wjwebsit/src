package main

import (
	"sort"
	"strconv"
)

/**
	算法描述
	给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。

示例 1:

输入: "2-1-1"
输出: [0, 2]
解释:
((2-1)-1) = 0
(2-(1-1)) = 2
示例 2:

输入: "2*3-4*5"
输出: [-34, -14, -10, -10, 10]
解释:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/different-ways-to-add-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
//定义变量操作符
var operator = map[byte]func(int,int)int {
	'+' : func (a,b int) int {return a + b},
	'-' : func (a,b int) int {return a - b},
	'*' : func (a,b int) int {return a * b},
}


func diffWaysToCompute(input string) []int {
	//判断
	if len(input) == 0 {
		return []int{}
	}

	//声明缓存
	cache := make(map[string][]int)

	//计算
	result := myCompute(0,len(input) - 1,input,&cache)

	//排序
	sort.Ints(result)

	//返回
	return result


}
func myCompute(s,e int,input string,cache *map[string][]int) []int {
	//判断
	if s == e {
		return []int{int(input[s] - '0')}
	}

	//判断缓存
	if _,ok := (*cache)[string(input[s:e + 1])];ok {
		return (*cache)[string(input[s:e + 1])]
	}

	//返回值
	var result []int

	//获取操作符
	fg := false //默认没有操作符
	for i := s; i <= e;i ++ {
		//获取操作符
		if _,ok := operator[input[i]];ok {//表示为操作符
			fg = true	//标记更改

			//左边的值
			left := myCompute(s,i - 1,input,cache)

			//右边的值
			right := myCompute(i + 1,e,input,cache)

			//笛卡尔积
			for _,l := range left {
				for _,r := range right {
					result = append(result,operator[input[i]](l,r))
				}
			}

		}
	}

	//判断
	if !fg {//没有操作符表示位数字
		num,_ := strconv.Atoi(string(input[s:e + 1]))
		result = append(result,num)
	}

	//写入缓存
	(*cache)[string(input[s:e + 1])] = result

	//返回
	return result

}
