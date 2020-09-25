package main

import "fmt"

/**
	算法描述
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
 */

/**
	利用栈实现--当前柱子没有超过之前出现的最大的一直往后直到出现 。盛水容量回溯
 */
func trap(height []int) int {
	//声明栈--记录
	var stack []int

	//声明sum
	sum := 0
	for i := 0;i < len(height);i ++ {
		//判断是否为top >=
		for len(stack) != 0 && height[stack[len(stack) - 1]] <= height[i]{ //当前比栈顶大
			//出栈
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack) - 1]

			//获取
			h := minHeight(height[i],height[left])
			w := i - left - 1
			sum += w * (h - height[top])
		}

		//写入
		stack = append(stack,i)
	}
	//返回
	return sum
}
func minHeight(a,b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}



func main (){
	arr := []int{4,2,3}

	fmt.Println(trap(arr))
}
