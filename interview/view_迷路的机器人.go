package main
/**
	算法描述
	返回一条可行的路径，路径由经过的网格的行号和列号组成。左上角为 0 行 0 列。如果没有可行的路径，返回空数组。

示例 1:

输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: [[0,0],[0,1],[0,2],[1,2],[2,2]]
解释:
输入中标粗的位置即为输出表示的路径，即
0行0列（左上角） -> 0行1列 -> 0行2列 -> 1行2列 -> 2行2列（右下角）
说明：r 和 c 的值均不超过 100。

 */
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	//判断
	if len(obstacleGrid) == 0 {
		return nil
	}
	//获取行和列
	r,c := len(obstacleGrid),len(obstacleGrid[0])
	if c == 0 {
		return nil
	}

	//判断左上角
	if obstacleGrid[0][0] == 1 {
		return nil
	}

	//声明路径
	var path [101][101][][]int

	//初始化
	path[0][0] = append(path[0][0],[]int{0,0})

	//第一行
	for i := 0 ; i < c ; i ++ {
		if obstacleGrid[0][i] != 1 {//没有障碍
			path[0][i] = append(path[0][i],path[0][i - 1]...)
			path[0][i] = append(path[0][i],[]int{0,i})

		} else {//存在障碍物
			break
		}
	}

	//第一列
	for i := 0 ; i < r ; i ++ {
		if obstacleGrid[r][0] != 1 {//没有障碍
			path[r][0] = append(path[r][0],path[r - 1][0]...)
			path[r][0] = append(path[r][0],[]int{i,0})

		} else {//存在障碍物
			break
		}
	}
	//定义是否到达
	isOk := false

	//中间部分
	for i := 1; i < r && !isOk; i ++ {
		for j := 1; j < c && !isOk; j ++ {
			//判断
			if obstacleGrid[i][j] == 1 {//当前为障碍物
				continue
			}

			//判断
			if len(path[i - 1][j]) == 0 && len(path[i][j - 1]) == 0 {
				//到达不了此位置
				continue
			}

			//判断
			if len(path[i - 1][j]) > 0 {
				path[i][j] = append(path[i][j],path[i - 1][j]...)
				path[i][j] = append(path[r][0],[]int{i,j})
			} else {
				path[i][j] = append(path[i][j],path[i][j - 1]...)
				path[i][j] = append(path[r][0],[]int{i,j})
			}

			//判断
			if i == r - 1 && j == c - 1 {
				isOk = true
				break
			}

		}
	}
	//返回
	return path[r - 1][c - 1]
}