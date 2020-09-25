package main

import "database/sql"

/**
	算法描述:
	编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func searchMatrix2(matrix [][]int, target int) bool {
	//判断数组
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	//获取行和列
	rows,cols := len(matrix),len(matrix[0])

	//声明返回值
	result := false

	//每行顺时针查找并且弧越来越小
	/*for row := 0; row < rows;row++ {//下一圈比上一圈值大
		//判断
		if matrix[row][0] > target {//当前圈最小值比目标值大，下面的圈均比目标值大
			break
		}

		//判断区间是在水平还是垂直
		if matrix[row][cols - row] >= target {//水平二分查找
			//水平二分查找
			result = binarySearchMatrix(matrix,0,cols - row,0,row,0,target)
		} else if matrix[rows - 1][cols - row] >= target { //垂直二分查找
			//垂直二分查找
			result = binarySearchMatrix(matrix,row,rows - 1,1,row,cols - row,target)
		}

		//判断是否找到
		if result == true {
			break
		}
	}

	//返回
	return result*/
	//从右下角开始
	row,col := rows - 1,0

	for row < rows && col < cols {
		//判断
		if matrix[row][col] > target {
			 row --
		} else if matrix[row][col] < target {
			 col ++
		} else {
			 	result = true
			 	break
		}


	}

	return result
}
/**
	二分查找
 */
func binarySearchMatrix(matrix [][]int,s,e int, way int,row,col int,target int) bool{
	//判断方式
	if way == 0 {//水平查找
		for s <= e {
			mid := (s + e) / 2
			//判断
			if matrix[row][mid] > target {
				e = mid - 1
			} else if matrix[row][mid] < target {
				s = mid + 1
			} else {
				return true
			}
		}
	} else {//垂直查找
		for s <= e {
			mid := (s + e ) / 2

			//判断
			if matrix[mid][col] > target {
				e = mid - 1
			}  else if matrix[mid][col] < target {
				s = mid + 1
			} else {
				return true
			}

		}
	}

	//返回
	return false

}