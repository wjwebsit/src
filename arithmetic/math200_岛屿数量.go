package main
/**
	算法描述:
	给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

 

示例 1:

输入:
11110
11010
11000
00000
输出: 1
示例 2:

输入:
11000
11000
00100
00011
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 */
/**
	采用回溯算法
 */
func numIslands(grid [][]byte) int {
	//判断
	if len(grid) == 0 || len(grid[0]) == 0 {
		//返回
		return 0
	}

	//获取岛屿数组的行和列
	rows,cols := len(grid),len(grid[0])

	//声明缓存数组
	var cache = make([][]int,rows)

	//初始化
	for i := 0; i < rows; i ++ {
		cache[i] = make([]int,cols)
	}

	//声明返回
	sum := 0

	//循环行
	for row := 0; row < rows; row ++ {
		//循环列
		for col := 0; col < cols;col ++ {
			//判断是否为'0'
			if grid[row][col] == '0' {
				continue
			}

			//表示为‘1’
			if cache[row][col] == 1 {//存在缓存中不做计算
				continue
			}

			//获取岛屿
			getLand(grid,rows,cols,row,col,&cache)
			sum ++

		}

	}

	//返回结果
	return sum
}
/**
	利用回溯算法 求解相连的岛屿并写入缓存
 */
func getLand(board [][]byte,rows,cols int,row,col int,cache *[][]int) {
	//判断缓存
	if (*cache)[row][col] == 1 {
		return
	}

	//写入缓存
	(*cache)[row][col] = 1


	//判断左边相邻
	if col - 1 >= 0 && board[row][col - 1] == '1' {
		//dfs
		getLand(board,rows,cols,row,col - 1,cache)
	}

	//判断右边相邻
	if col + 1 < cols && board[row][col + 1] == '1' {
		//dfs
		getLand(board,rows,cols,row,col + 1,cache)
	}

	//判断上相邻
	if row - 1 >= 0 && board[row - 1][col] == '1' {
		//dfs
		getLand(board,rows,cols,row - 1,col,cache)
	}

	//判断下相邻
	if row + 1 < rows && board[row + 1][col] == '1' {
		//dfs
		getLand(board,rows,cols,row + 1,col,cache)
	}

}