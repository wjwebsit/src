package main

import "fmt"

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
 */

func main() {
	str := "{[]}"
	fmt.Print(isValid(str))
}
/**
	判断是否合法
 */
func isValid(s string) bool {


	//定义返回 默认合法
	isValid := true

	//定义栈
	stack := []byte{}

	//定义map
	left,right := map[byte]int {},map[byte]int{}

	left['('] = 1
	right[')'] = -1
	left['['] = 2
	right[']'] = -2
	left['{'] = 3
	right['}'] = -3

	//遍历---当sum >0 时就表示false
	for i := 0; i < len(s);i++ {
		//判断左边
		if left[s[i]] > 0 {
			//进栈
			stack = append(stack,s[i])
			continue
		}

		//判断右边
		if right[s[i]] < 0{

			//判断栈元素是否为空
			if len(stack) == 0 {
				isValid = false
				break
			}

			//出栈
			preLeft := stack[len(stack) - 1]

			//判断是否相等
			if left[preLeft] + right[s[i]] != 0 {
				isValid = false
				break
			} else {
				stack = stack[:len(stack) - 1]
			}
		}

	}

	//判断栈是否出完
	if len(stack) != 0 {
		isValid = false
	}

	//返回结果
	return isValid

}
