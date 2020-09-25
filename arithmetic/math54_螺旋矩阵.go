package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	//声明返回值
	var result []int

	//判断是否为空
	if len(matrix) == 0 {
		return result
	}

	//获取行和列
	rows,cols := len(matrix) - 1,len(matrix[0]) - 1

	//从外圈向内依次输出 获取圈数
	var n int
	if rows < cols {
		n = (rows + 2) / 2
	} else {
		n = (cols + 2) / 2
	}


	//开始输出
	for i := 0; i < n;i ++ {
		//判断只有一列
		if i == cols - i {
			for j := i; j <= rows - i;j++ {
				result = append(result, matrix[j][i])
			}
			break
		}

		//判断是否只有一行
		if i == rows - i {
			for j := i; j <= cols - i;j++ {
				result = append(result, matrix[i][j])
			}
			break
		}

		//顺时针第一行
		for j := i;j < cols - i; j ++ {
			//写入结果
			result = append(result,matrix[i][j])
		}


		//顺时针第一列
		for j := i;j < rows - i;j ++ {
			result = append(result,matrix[j][cols - i])
		}

		//顺时针第二行
		for k := cols - i;k > i; k -- {
			result = append(result,matrix[rows - i][k])

		}

		//顺时针最后一列
		for m := rows - i; m > i;m -- {
			result = append(result,matrix[m][i])
		}
	}

	//返回
	return result
}
func main() {
	/*matrix := [][]int{
		[]int{1,11},
		[]int{2,12},
		[]int{3,13},
		[]int{4,14},
		[]int{5,15},
		[]int{6,16},
		[]int{7,17},
		[]int{8,18},
		[]int{9,19},
		[]int{10,20},
	}*/
	matrix := [][]int{[]int{1},[]int{2},[]int{3}}
	fmt.Println(spiralOrder(matrix))
}