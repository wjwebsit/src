package main

import (
	"fmt"
	"strconv"
)

/**
	算法描述:
	给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.
 */
func exist(board [][]byte, word string) bool {
	//判断
	if len(board) == 0 {
		//返回
		return false
	}

	//获取行和列
	rows ,cols := len(board) - 1,len(board[0]) - 1

	//声明
	var dp [][]int
	for i:= 0; i < len(word); i++ {
		dp = append(dp,[]int{})
	}

	//定义缓存来判断是否重复
	caches := []map[string]int{}


	//初始化dp[i]表示word[i]字母出现的位置dp[0] = [0,0,2,0] 有2个位置（0，0）he (2,0)
	for i := 0; i <= rows; i++ {
		for j := 0; j <= cols;j++ {
			mp := map[string]int{}
			if board[i][j] == word[0] {
				dp[0] = append(dp[0],i)
				dp[0] = append(dp[0],j)
				mp[strconv.Itoa(i)+ strconv.Itoa(j)] = 1
				caches = append(caches,mp)

			}
		}
	}
	//获取单词长度
	length := len(word) - 1

	//判断
	if len(dp[0]) > 0 { //表示有对应的开始
		//遍历字串
		for i := 1; i <= length;i++ {
			//获取字节
			key := word[i]

			//遍历上一步的位置
			for index := 0; index < len(dp[i - 1]); index += 2 {
				//获取横纵坐标
				x,y := dp[i - 1][index], dp[i - 1][index + 1]

				//向左添加只能用一次的限制
				if y - 1 >= 0 && board[x][y - 1] == key && caches[index][strconv.Itoa(x)+ strconv.Itoa(y - 1)] != 1 {
					dp[i] = append(dp[i],x)
					dp[i] = append(dp[i],y -1)
					//写入缓存
					caches[index][strconv.Itoa(x)+ strconv.Itoa(y - 1)] = 1
				}
				// 向右
				if y + 1 <= cols && board[x][y + 1] == key {
					dp[i] = append(dp[i],x)
					dp[i] = append(dp[i],y +1)
					//写入缓存
					caches[index][strconv.Itoa(x)+ strconv.Itoa(y + 1)] = 1
				}
				// 和向下
				if x + 1 <= rows && board[x + 1][y] == key {
					dp[i] = append(dp[i],x + 1)
					dp[i] = append(dp[i],y)
				}
			}
		}
	}

	//返回结果
	return len(dp[length]) > 0

}


/**
	测试
 */
 func main() {
 	var board [][]byte
	board = append(board,[]byte{'C','A','A'})
	board = append(board,[]byte{'A','A','A'})
	board = append(board,[]byte{'B','C','D'})

	fmt.Println(exist(board,"AAB"))

 }