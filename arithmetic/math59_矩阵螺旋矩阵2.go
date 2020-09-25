package main

import "fmt"

/**
	算法描述：
	给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
 */

func generateMatrix(n int) [][]int {
	//声明生成的矩阵
	var matrix [][]int

	//初始化
	for i := 0; i < n; i ++ {
		matrix = append(matrix,make([]int,n))
	}

	//生成的矩阵行数和列数 均为n
	rows,cols := n - 1,n - 1

	//获取圈数---一圈圈的
	squary := (n + 1) / 2

	//数字初始化为0
	num := 1

	//从完全开始
	for q := 0; q < squary;q ++ {
		//如果行和列相等时
		if q == cols - q {
			matrix[q][q] = num
			break
		}

		//顺时针第一行
		for k := q; k < cols - q;k ++ {
			matrix[q][k] = num
			num ++
		}

		//顺时针第一列
		for k := q ; k < rows - q; k++ {
			matrix[k][cols - q] = num
			num ++
		}

		//顺时针第二行
		for k := cols - q; k > q; k -- {
			matrix[rows - q][k] = num
			num ++
		}

		//顺时针最后列
		for k := rows - q; k > q; k -- {
			matrix[k][q] = num
			num ++
		}
	}

	//返回
	return matrix

}
func main() {
	fmt.Println(generateMatrix(1))
}

