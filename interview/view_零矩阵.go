package main
/**
	算法描述
	编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
示例 1：

输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：

输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
 */
func setZeroes(matrix [][]int)  {
	//获取行和列
	rows,cols := len(matrix),len(matrix[0])

	//判断第一行和第一列是否需要重置0
	zeroRow,zeroCol := false,false

	//循环
	for col := 0 ; col < cols; col ++ {
		if matrix[0][col] == 0 {
			zeroRow = true
			break
		}
	}

	//循环
	for row := 0 ; row < rows; row ++ {
		if matrix[row][0] == 0 {
			zeroCol = true
			break
		}
	}

	//中间
	for row := 1; row < rows;row ++ {
		for col := 1; col <  cols;col ++ {
			//判断
			if matrix[row][col] == 0 {
				matrix[row][0] = 0 //行
				matrix[0][col] = 0 //列
			}

		}
	}

	//设置为0
	for row := 1; row < rows;row ++ {
		for col := 1; col <  cols;col ++ {
			//判断
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}

		}
	}

	//设置第一行和第一列
	if zeroRow == true {
		//循环
		for col := 0 ; col < cols; col ++ {
			matrix[0][col] = 0

		}
	}
	if zeroCol == true {
		//循环
		for row := 0 ; row < rows; row ++ {
			 matrix[row][0] = 0
		}
	}


}