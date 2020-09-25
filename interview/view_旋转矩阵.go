package main
/**
	算法描述
	给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？

 

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
 */
func rotate(matrix [][]int)  {
	//获取行和列
	rows,cols :=  len(matrix),len(matrix[0])

	//获取圈数
	n := (rows + 1) / 2

	//原址旋转
	for i := 0; i < n; i ++ {
		if i == rows - 1 - i {//中间只有一个元素
			break
		}

		for col := i ; col < cols - 1; col ++ {
			//记录位置
			a := matrix[i][col]
			b := matrix[col][cols - 1 - i]
			c := matrix[rows - 1 - i][cols - 1 - col]
			d := matrix[rows - 1 - col][i]

			//交换
			matrix[i][col] = d
			matrix[col][cols - 1 - i] = a
			matrix[rows - 1 - i][cols - 1 - col] = b
			matrix[rows - 1 - col][i] = c
		}
	}
}