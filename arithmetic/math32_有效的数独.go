package main

import "encoding/xml"

/**
	算法描述
	判断一个9x9的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字  1-9 在每一行只能出现一次。
数字  1-9 在每一列只能出现一次。
数字  1-9 在每一个以粗实线分隔的  3x3 宫内只能出现一次。

数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true
示例 2:

输入:
[
[".",".",".",".","5",".",".","1","."],
[".","4",".","3",".",".",".",".","."],
[".",".",".",".",".","3",".",".","1"],
["8",".",".",".",".",".",".","2","."],
[".",".","2",".","7",".",".",".","."],
[".","1","5",".",".",".",".",".","."],
[".",".",".",".",".","2",".",".","."],
[".","2",".","9",".",".",".",".","."],
[".",".","4",".",".",".",".",".","."]


]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的
 */

func isValidSudoku(board [][]byte) bool {
	//行
	for row := 0 ; row < len(board); row ++ {
		//遍历每一行的每一项
		mp := map[byte]int {}
		for hight := 0; hight < len(board[row]);hight ++ {
			//判断是否为.
			if board[row][hight] == '.' {
				continue
			}

			if !inMapInt2(board[row][hight],mp) {
				mp[board[row][hight]] = 0
			}  else {
				return false

			}
		}
	}

	//列
	for hight := 0 ; hight < len(board[0]); hight ++ {
		//遍历每一行的每一项
		mp := map[byte]int {}
		for row := 0; row < len(board);row ++ {
			//判断是否为.
			if board[row][hight] == '.' {
				continue
			}

			if !inMapInt2(board[row][hight],mp) {
				mp[board[row][hight]] = 0
			}  else {
				return false

			}
		}
	}

	// 3x3
	for i := 2; i< len(board); i += 3 {
		for j := 2; j < len(board[i]);j += 3 {
			mp := map[byte]int {}
			for start := i - 2; start <= i; start ++ {
				for end := j - 2; end <= j; end ++   {
					if board[start][end] == '.' {
						continue
					}
					if !inMapInt2(board[start][end],mp) {
						mp[board[start][end]] = 0
					}  else {
						return false

					}
				}
			}

		}

	}
	return true

}

func inMapInt2(key byte, mp map[byte]int) bool {
	for k,_ := range mp {
		if k == key {
			return true
		}
	}
	return false
}
