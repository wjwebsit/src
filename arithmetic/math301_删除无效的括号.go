package main

import "fmt"

/**
	算法描述
	删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。

说明: 输入可能包含了除 ( 和 ) 以外的字符。

示例 1:

输入: "()())()"
输出: ["()()()", "(())()"]
示例 2:

输入: "(a)())()"
输出: ["(a)()()", "(a())()"]
示例 3:

输入: ")("
输出: [""]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-invalid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
 /**
 	BFS
  */
func removeInvalidParentheses(s string) []string {
	//声明返回结果
	var result []string

	//声明map
	mp := make(map[string]int)
	mp[s] = 1

	for len(mp) > 0 {
		//遍历
		for k,_ := range mp {
			if checkValid(k) {
				result = append(result,k)
			}

		}

		//判断
		if len(result) > 0 {
			break
		}

		//删除
		mp2 := make(map[string]int)
		for k,_ := range mp {

			for i := 0 ; i < len(k); i ++ {
				//
				if k[i] == '(' || k[i] == ')' {
					if i == 0 {
						mp2[string(k[1:])] = 1
					} else if i == len(k) - 1 {
						mp2[string(k[:len(k)-1])] = 1
					} else {
						mp2[string(k[:i])+string(k[i+1:])] = 1
					}
				}

			}
		}
		mp = mp2
		delete(mp,"")

	}
	if len(result) == 0 {
		result = append(result,"")
	}

	//返回
	return result
}
func checkValid(s string) bool {
	//利用
	if len(s) == 0 {
		return false
	}
	ans := 0
	for i := 0; i < len(s) ; i ++ {
		if s[i] == '(' {
			ans += 1
		} else if s[i] == ')' {
			ans -= 1
		}
		if ans < 0 {
			return false
		}
	}

	return  ans == 0
}