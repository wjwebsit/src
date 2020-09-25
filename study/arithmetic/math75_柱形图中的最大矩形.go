package main

import (
	"fmt"
)

/**
	算法描述:
	给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
示例:

输入: [2,1,5,6,2,3]
输出: 10
 */

func largestRectangleArea(heights []int) int {
	//定义面积
	maxArea := 0

	//定义单调栈
	var stack []int
	heights = append(heights, -1)

	//遍历柱子
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] { //出栈--
			//获取栈顶
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//h
			h := heights[top]
			var w int
			if len(stack) == 0 {
				w = i
			} else {
				w = i - stack[len(stack) - 1] - 1
			}



			//计算面积
			maxArea = maxHeightArea(maxArea, h*w)

		}

		//入栈---栈为空或者当前元素大于栈顶
		stack = append(stack, i)
	}






	//返回
	return maxArea
}
func maxHeightArea(a,b int)int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	arr := []int{2,1,2}

	fmt.Println(largestRectangleArea(arr))
}