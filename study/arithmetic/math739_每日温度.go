package main

import "fmt"

/**
	算法描述
	根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
 */

func dailyTemperatures(T []int) []int {
	//判断
	if len(T) == 0 {
		return T
	}

	//声明单调栈（单调递减栈）
	var stack []int

	//第一个元素下标入栈
	stack = append(stack,0)

	//声明返回
	//result := make([]int,len(T))

	//循环
	for i := 1; i < len(T); i ++ {
		//判断当前和栈顶元素
		for len(stack) > 0 && T[stack[len(stack) - 1]] < T[i] {
			//栈顶出栈
			top := stack[len(stack) - 1]
			stack = stack[0 :len(stack) - 1]

			//修改天数
			T[top] = i - top

		}

		//入栈
		stack = append(stack,i)

	}

	//判断栈是否为空
	if len(stack) > 0 {
		for _,v := range stack {
			T[v] = 0
		}
	}

	//返回
	return T
}
func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}