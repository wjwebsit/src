package main
/**
	算法描述:
	编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	//获取矩阵的行和列
	rows,cols := len(matrix) - 1,len(matrix[0]) - 1

	//二分查找行
	s,e := 0,rows
	targetRow := -1
	for s <= e {
		//获取中间行
		mid := (s + e) / 2

		//判断
		if matrix[mid][0] > target {//在 s-mid-1 行中
			e = mid - 1
		} else if matrix[mid][cols] < target{ //在 mid + 1---e
			s = mid + 1
		} else {//表示在当前
			targetRow = mid
			break
		}
	}

	//判断是否查到行
	if targetRow == -1 {
		return  false
	}

	//在当前行二分查找
	s,e = 0,cols
	rt := -1
	for s <= e {
		mid := (s + e) / 2
		//判断
		if matrix[targetRow][mid] < target {
			s = mid + 1
		} else if matrix[targetRow][mid] > target {
			e = mid - 1
		} else {
			rt = mid
			break
		}
	}
	if rt == -1 {
		return  false
	}

	//返回
	return  true

}