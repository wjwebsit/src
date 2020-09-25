package main
/**
	算法描述
	给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

示例 1:

输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2:

输入:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
 */
func setZeroes(matrix [][]int)  {
	//判断
	if len(matrix) == 0 {
		return
	}

	//获取行和列
	rows,cols := len(matrix),len(matrix[0])
	fgRow := false
	fgCol := false
	for col := 0;col < cols;col ++ {
		if matrix[0][col] == 0 {
			fgRow = true
			break
		}
	}
	for row := 0; row < rows;row ++ {
		if matrix[row][0] == 0 {
			fgCol = true
			break
		}
	}


	//利用第一行和第一列来记录需要置零的行和列---从matrix[1][1]开始
	for row := 1;row < rows; row ++ {
		for col := 1; col < cols;col ++ {
			//判断是否为0
			if matrix[row][col] == 0 {
				matrix[0][col] = 0
				matrix[row][0] = 0

			}
		}
	}

	//置换
	for row := 1;row < rows;row ++ {
		for col := 1 ;col < cols;col ++ {
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}
		}

	}

	//最后修改第一行一列
	if fgRow == true {
		for col := 0;col < cols;col ++ {
			 matrix[0][col] = 0
		}
	}
	if fgCol == true {
		for row := 0; row < rows;row ++ {
			matrix[row][0] = 0
		}
	}

}