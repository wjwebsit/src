package main

import "fmt"

/**
	算法描述：
给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。

规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。

请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。

 

提示：

输出坐标的顺序不重要
m 和 n 都小于150
 

示例：

 

给定下面的 5x5 矩阵:

  太平洋 ~   ~   ~   ~   ~
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * 大西洋

返回:

[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pacific-atlantic-water-flow
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func pacificAtlantic(matrix [][]int) [][]int {
	//声明返回值
	var result [][]int

	//获取行和列
	rows := len(matrix)
	if rows == 0 {
		return result
	}
	cols := len(matrix[0])
	if cols == 0 {
		return result
	}

	//定义缓存、结果
	var cache [151][151]int

	for i := 0; i < rows; i ++ {
		for j := 0; j < cols;j ++ {
			//判断缓存
			if cache[i][j] > 0 {
				if cache[i][j] == 3 {
					result = append(result,[]int{i,j})
				}
			} else {
				path := [151][151]int{}
				fg := pacificAtlanticHS(i,j,rows,cols,matrix,&path,&cache)

				if fg == 3 {
					result = append(result,[]int{i,j})
				}
			}

		}
	}

	//返回
	return result

}
/**

 */
func pacificAtlanticHS(i,j,rows,cols int,matrix [][]int,path *[151][151]int ,cache *[151][151]int) int  {
	//判断
	if i < 0 || j < 0 {
		return 1  //太平洋
	}
	if i >= rows || j >= cols {//大西洋
		return  2
	}

	//判断缓存
	if (*cache)[i][j] == 3 {
		return (*cache)[i][j]
	}

	//当前值
	cur := matrix[i][j]

	//写入当前路径
	(*path)[i][j] = 1

	//定义当前是否流向大西洋和太平洋
	t,d := false,false

	//上
	if i - 1 >= 0 && matrix[i - 1][j] <= cur {

		if matrix[i - 1][j] < cur || ( matrix[i - 1][j] == cur && (*path)[i - 1][j] != 1) {
			//流
			fg := pacificAtlanticHS(i - 1,j,rows,cols,matrix,path,cache)

			if fg == 1 {
				t = true
			} else if fg == 2 {
				d = true
			} else if fg == 3{ //当前结点流向太平洋和大西洋
				t = true
				d = true
			}
		}

	} else if i - 1 < 0 {//太平洋
		t = true
	}

	//下
	if i + 1 < rows && matrix[i + 1][j] <= cur{
		//判断
		if matrix[i + 1][j] < cur || ( matrix[i + 1][j] == cur && (*path)[i + 1][j] != 1){
			//流
			fg := pacificAtlanticHS(i + 1,j,rows,cols,matrix,path,cache)

			if fg == 1 {
				t = true
			} else if fg == 2 {
				d = true
			} else if fg == 3 { //当前结点流向太平洋和大西洋
				t = true
				d = true
			}
		}
	} else if i + 1 >= rows { //大西洋
		d = true
	}

	//左
	if j - 1 >= 0 && matrix[i][j - 1] <= cur {
		//判断
		if matrix[i][j - 1] < cur || ( matrix[i][j - 1] == cur && (*path)[i][j - 1] != 1){
			//流

			fg := pacificAtlanticHS(i,j - 1,rows,cols,matrix,path,cache)

			if fg == 1 {
				t = true
			} else if fg == 2 {
				d = true
			} else if fg == 3{ //当前结点流向太平洋和大西洋
				t = true
				d = true
			}
		}

	} else if j - 1 < 0  {//太平洋
		t = true
	}

	//右
	if j + 1 < cols && matrix[i][j + 1] <= cur{
		//判断
		if matrix[i][j + 1] < cur || ( matrix[i][j + 1] == cur && (*path)[i][j + 1] != 1){
			//流
			fg := pacificAtlanticHS(i,j + 1,rows,cols,matrix,path,cache)

			if fg == 1 {
				t = true
			} else if fg == 2 {
				d = true
			} else if fg == 3 { //当前结点流向太平洋和大西洋
				t = true
				d = true
			}
		}
	} else if j + 1 >= cols {
		d = true

	}

	//判断
	if t == true && d == true {
		(*cache)[i][j] = 3
	} else if t == true && d == false {
		return 1

	} else if t == false && d == true { //大西洋
		return  2
	} else {
		return 4
	}

	//回溯
	(*path)[i][j] = 0

	//返回
	return (*cache)[i][j]
}

func main() {
	matrix := [][]int{
		{10,10,10},
		{10,1,10},
		{10,10,10},
	}
	fmt.Println(pacificAtlantic(matrix))

}


