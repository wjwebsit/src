package main

import "fmt"

/**
	算法描述
	给定一个n  ×  n的二维矩阵表示一个图像。

将图像顺时针旋转90度。

说明：

必须你在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例1：

给定矩阵 =
[
  [1,2,3]，
  [4,5,6]
  [7,8,9]
]

原地旋转输入侧矩阵，使其变为：
[
  [7,4,1]
  [8,5,2]
  [9,6,3]
]
示例2：

给定矩阵 =
[
  [5,1,9,11]，
  [2,4,8,10]，
  [13,3,6,7]，
  [15,14,12,16]
]

原地旋转输入侧矩阵，使其变为：
[
  [15,13,​​2,5]，
  [14,1,3,1]，
  [12,6,8,9]，
  [16,7,10,11]
]
 */
func rotate1(matrix [][]int) {
	//判断
	if len(matrix) == 0 || len(matrix) == 1 {
		return

	}
	//获取行和列
	rows,cols := len(matrix) - 1 ,len(matrix) - 1

	//获取旋转圈数---
	n := (rows + 1) / 2


	//从外圈逐层向内层一圈的
	for i := 0; i < n;i ++ {
		for k := i;k < cols - i;k ++ {

			//依次获取交换位置值
			a := matrix[i][k]
			b := matrix[k][cols - i]

			c := matrix[rows - i][cols - k]
			d := matrix[rows - k][i]

			//交换
			matrix[i][k] = d
			matrix[k][cols - i] = a
			matrix[rows - i][cols - k] = b
			matrix[rows - k][i] = c

		}

	}

	fmt.Println(matrix)

}
func main()  {
	matrix := [][]int{[]int{5,1,9,11},[]int{2,4,8,10},[]int{13,3,6,7},[]int{15,14,12,16}}
	rotate1(matrix)
}