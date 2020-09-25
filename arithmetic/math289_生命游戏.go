package main

import "fmt"

/**
	算法描述
	根据 百度百科 ，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态：1 即为活细胞（live），或 0 即为死细胞（dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
根据当前状态，写一个函数来计算面板上所有细胞的下一个（一次更新后的）状态。下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。

 

示例：

输入：
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
输出：
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]
 

进阶：

你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？
 */
/**原址实现
	用数字3表示 原来是0改为1
	用数字2表示 原来是1改为0
	其它未改变
 */
func gameOfLife(board [][]int)  {
	//获取行
	rows := len(board)

	//判断
	if rows == 0 {
		return
	}

	//获取列
	cols := len(board[0])

	//判断
	if cols == 0 {
		return
	}

	//遍历
	for row := 0; row < rows; row ++ {
		for col := 0; col < cols;col ++ {
			//获取
			r := col + 1
			l := col - 1
			up := row - 1
			down := row + 1

			//统计活细胞数目
			count := 0

			//左上角
			if l >= 0 && up >= 0 && (board[up][l] == 1 || board[up][l] == 2) {
				count ++
			}

			//上
			if up >= 0 && (board[up][col] == 1 || board[up][col] == 2) {
				count ++
			}

			//右上角
			if r < cols && up >= 0 && (board[up][r] == 1 || board[up][r] == 2) {
				count ++
			}

			//左
			if l >= 0 && (board[row][l] == 1 || board[row][l] == 2){
				count ++
			}

			//右
			if r < cols && (board[row][r] == 1 || board[row][r] == 2) {
				count ++
			}

			//左下角
			if down < rows && l >= 0 && (board[down][l] == 1 || board[down][l] == 2) {
				count ++
			}

			//下
			if down < rows && (board[down][col] == 1 || board[down][col] == 2) {
				count ++
			}

			//右下角
			if down < rows && r < cols && (board[down][r] == 1 || board[down][r] == 2) {
				count ++
			}

			//判断
			if count < 2 {
				if board[row][col] == 1 {//活细胞死亡
					board[row][col] = 2
				}

			} else if count <= 3 {//2到3个细胞存活不处理
				if count == 3 && board[row][col] == 0 {//死细胞复活
					board[row][col] = 3
				}
			} else {
				if board[row][col] == 1 {//活细胞死亡
					board[row][col] = 2
				}
			}

		}

	}

	//处理原址
	for row := 0; row < rows;row ++ {
		for col := 0 ; col < cols;col ++ {
			if board[row][col] >= 2 {
				board[row][col] -= 2
			}
		}
	}
	fmt.Println(board)
}