package main
/**
	算法描述
	给定M×N矩阵，每一行、每一列都按升序排列，请编写代码找出某元素。

示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。
 */
func searchMatrix(matrix [][]int, target int) bool {
	//判断
	if len(matrix) == 0 {
		return false
	}
	//获取矩阵的行和列
	row,col := len(matrix),len(matrix[0])

	//左下角坐标
	i,j := row - 1, 0
	for i >= 0 && j < col {
		//判断
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
				i--
		} else {
			 j++
		}
	}

	return false

}