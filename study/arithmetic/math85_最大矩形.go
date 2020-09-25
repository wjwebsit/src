package main


/**
	算法描述:
	给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例:

输入:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]

 */
func maximalRectangle(matrix [][]byte) int {
	//判断
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return  0
	}

	//获取行和列
	rows,cols := len(matrix),len(matrix[0])

	//高度数组
	heights := make([]int,cols + 1)
	max := 0

	//遍历矩阵
	for row := 0; row < rows;row ++ {
		for col := 0; col < cols;col ++ {
			if matrix[row][col] == '1' {
				//构造矩形
				heights[col] = heights[col] + int(matrix[row][col] - '0')
			} else {
				//构造矩形
				heights[col] = 0
			}

		}
		//求最大面积
		area := getMaxArea(heights,cols)
		if area > max {
			max = area
		}
	}

	//返回
	return max
}
/**
	利用单调栈求解矩形的最大面积
 */
func getMaxArea(heights []int,cols int) int {
	//声明栈
	var stack []int
	heights[cols] = -1

	//返回max
	var max = 0

	//遍历高度
	for i := 0; i < len(heights); i ++ {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] { //当前比栈顶小--出栈
			//出栈
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			//声明宽度
			var width int

			//判断栈是否为空
			if len(stack) == 0 { //表示 0-i之间为最小
				width = i
			} else {
				width = i - stack[len(stack) - 1] - 1
			}

			//求解面积
			area := heights[top] * width
			if area > max {
				max = area
			}
		}

		//入栈
		stack = append(stack,i)

	}

	//返回
	return max

}