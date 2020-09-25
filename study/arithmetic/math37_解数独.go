package main
/**
	算法描述:
	一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。
 */
func solveSudoku(board [][]byte)  {
	//判断
	if len(board) == 0 {
		return
	}

	//构造每一行需要的填充的数组
	rows,cols := make([]map[int]int,9),make([]map[int]int,9)
	square := make([][]map[int]int,3)
	//初始化
	/*for i := 0; i < 9; i ++ {
		rows = append(rows, map[int]int{})
		cols = append(cols, map[int]int{})


	}*/
	for i := 0; i < 3;i++ {
		tmp := []map[int]int{}
		for j := 0; j < 3;j++ {
			tmp = append(tmp,map[int]int{})
		}
		square = append(square,tmp)

	}
	for i := 0; i < 9; i ++ {//行
		//列
		for j := 0; j < 9; j ++ {
			//判断
			if board[i][j] != '.' {
				//写入行
				rows[i][int(board[i][j] - '0')] = 1

				//写入列
				cols[j][int(board[i][j] - '0')] = 1

				//写入3X3
				square[i / 3][j / 3][int(board[i][j] - '0')] = 1
			}
		}
	}

	//解数独
	solveSudokuHS(0,0,rows,cols,square,board)
}
func solveSudokuHS(row,col int,rows,cols []map[int]int,square [][]map[int]int,board [][]byte) bool {
	//判断终止条件
	if col == 9 {
		col = 0
		row ++
		if row == 9 {
			return true
		}
	}
	res := false
	if board[row][col] == '.' {
		//填充
		for k := 1; k <= 9; k ++ {
			if rows[row][k] == 1 || cols[col][k] == 1 || square[row / 3][col / 3][k] == 1 {
				continue
			}
			rows[row][k] = 1
			cols[col][k] = 1
			square[row / 3][col / 3][k] = 1
			board[row][col] = byte(k + '0')

			res = solveSudokuHS(row,col + 1,rows,cols,square,board)
			if res == false {//回溯
				rows[row][k] = 0
				cols[col][k] = 0
				square[row / 3][col / 3][k] = 0
				board[row][col] = '.'
			}

		}

	} else {
		return solveSudokuHS(row,col + 1,rows,cols,square,board)
	}
	
	//return 
	return res
}
func main() {

}