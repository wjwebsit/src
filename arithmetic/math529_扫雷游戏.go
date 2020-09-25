package main
/**
	算法描述：
	让我们一起来玩扫雷游戏！

给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。

现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：

如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的方块都应该被递归地揭露。
如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
如果在此次点击中，若无更多方块可被揭露，则返回面板。
 */
func updateBoard(board [][]byte, click []int) [][]byte {
	//获取点击坐标
	i,j := click[0],click[1]

	//判断是否点击过
	if board[i][j] != 'E' && board[i][j] != 'M' {
		return board
	}

	//判断点击是否为地雷
	if board[i][j] == 'M' {
		//修改爆炸的雷
		board[i][j] = 'X'

		//返回
		return board
	}

	//获取表格的行和列
	row,col := len(board),len(board[0])

	//点击空白---采用队列
	var queue [][]int
	queue = append(queue,click)

	//声明队列缓存
	var cache [50][50]int
	cache[i][j] = 1

	for len(queue) > 0 {
		//出队
		cur := queue[0]
		queue = queue[1:]

		//点击的坐标
		i,j := cur[0],cur[1]

		//声明四周雷的数目
		sum := 0

		//记录位置
		var address [][]int

		//左上角
		if (i - 1) >= 0 && (j - 1) >= 0   {
			if board[i - 1][j - 1] == 'M' {
				sum ++
			} else if board[i - 1][j - 1] == 'E' && cache[i - 1][j - 1] == 0 {//未点击的相邻
				address = append(address,[]int{i - 1,j - 1})

			}
		}

		//上
		if (i - 1) >= 0  {
			if board[i - 1][j] == 'M' {
				sum ++
			} else if board[i - 1][j] == 'E' && cache[i - 1][j] == 0{//未点击的相邻
				address = append(address,[]int{i - 1,j})

			}
		}

		//右上角
		if (i - 1) >= 0 && (j + 1) < col {
			if board[i - 1][j + 1] == 'M' {
				sum ++
			} else if board[i - 1][j + 1] == 'E' && cache[i - 1][j + 1] == 0 {//未点击的相邻
				address = append(address,[]int{i - 1,j + 1})

			}
		}

		//左
		if j - 1 >= 0 {
			if board[i][j - 1]  == 'M' {
				sum ++
			} else if board[i][j - 1] == 'E' && cache[i][j - 1] == 0{//未点击的相邻
				address = append(address,[]int{i,j - 1})

			}

		}

		//右
		if j + 1 < col {
			if board[i][j + 1]  == 'M' {
				sum ++
			} else if board[i][j + 1] == 'E' && cache[i][j + 1] == 0{//未点击的相邻
				address = append(address,[]int{i,j + 1})

			}

		}
		//左下角
		if (i + 1) < row && (j - 1) >= 0   {
			if board[i + 1][j - 1] == 'M' {
				sum ++
			}else if board[i + 1][j - 1] == 'E' && cache[i + 1][j - 1] == 0{//未点击的相邻
				address = append(address,[]int{i + 1,j - 1})

			}
		}

		//下
		if (i + 1) < row  {
			if board[i + 1][j] == 'M' {
				sum ++
			} else if board[i + 1][j] == 'E' && cache[i + 1][j] == 0 {//未点击的相邻
				address = append(address,[]int{i + 1,j})
			}
		}

		//右下角
		if (i + 1) < row  && (j + 1) < col {
			if board[i + 1][j + 1] == 'M' {
				sum ++
			}else if board[i + 1][j + 1] == 'E' && cache[i + 1][j + 1] == 0 {//未点击的相邻
				address = append(address,[]int{i + 1,j + 1})
			}
		}

		//判断
		if sum > 0 {
			//写入
			board[i][j] = byte(sum) + '0'

		} else {
			//表示挖出的空白
			board[i][j] = 'B'

			//串连
			if len(address) > 0 {
				for _,addr := range address {
					if cache[addr[0]][addr[1]] == 0 {
						queue = append(queue,addr)
						cache[addr[0]][addr[1]] = 1
					}

				}
			}
		}
	}

	//返回
	return board

}