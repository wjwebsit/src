package main

import
	(
	"strconv"
	"strings"
	)

/**
	算法描述：
	给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 */
func solve(board [][]byte)  {
	//获取矩阵的行
	rows := len(board)

	//判断
	if rows == 0 || rows == 1 {
		return
	}

	//获取列
	cols := len(board[0])

	//判断
	if cols == 0 || cols == 1 {
		return
	}

	//初始化缓存 表示默认表示封闭
	var cache = make([][]bool,rows)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols;col++  {
			cache[row] = append(cache[row],true)
		}
	}

	//初始化4周
	for col := 0; col < cols;col ++ {
		//第一行
		if board[0][col] == 'O' {
			cache[0][col] = false
		}

		//最后一行
		if board[rows - 1][col] == 'O' {
			cache[rows - 1][col] = false
		}
	}
	for row := 0; row < rows ; row ++ {
		//第一列
		if board[row][0] == 'O' {
			cache[row][0] = false
		}

		//最后一列
		if board[row][cols - 1] == 'O' {
			cache[row][cols - 1] = false
		}
	}

	//中间部分
	for row := 1; row < rows; row ++ {
		for col := 1;col < cols;col++ {

			//如果为‘O’并且缓存中为true
			if board[row][col] == 'O' && cache[row][col] == true {
				//声明path
				path := make(map[string]bool)
				if solveValid(board,row,col,cache,&path) == true {
					//设置
					for key ,_ := range path {
						tmpRow,_ := strconv.Atoi(strings.Split(key,"-")[0])
						tmpCol,_ := strconv.Atoi(strings.Split(key,"-")[1])
						board[tmpRow][tmpCol] = 'X'

					}

				} else {//写入cache
					for key ,_ := range path {
						tmpRow,_ := strconv.Atoi(strings.Split(key,"-")[0])
						tmpCol,_ := strconv.Atoi(strings.Split(key,"-")[1])
						cache[tmpRow][tmpCol] = false

					}
				}

			}
		}
	}
}

func solveValid(board [][]byte,row,col int,cache [][]bool,path *map[string]bool) bool {
	//判断当前是否为‘X’
	if board[row][col] == 'X' {
		return true
	}


	//判断当前是否在cache中
	if cache[row][col] == false {
		return false
	}

	//判断是否是之前出现的
	if (*path)[strconv.Itoa(row) + "-" + strconv.Itoa(col)] {
		return true //取决于后面的故设为true

	} else {//增加记忆
		(*path)[strconv.Itoa(row) + "-" + strconv.Itoa(col)] = true
	}

	//获取上下左右
	return (solveValid(board,row - 1,col,cache,path) &&
		solveValid(board,row + 1,col,cache,path) &&
		solveValid(board,row,col - 1,cache,path) &&
		solveValid(board,row,col + 1,cache,path))
}