package main
/**
	算法描述
	设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。

注意：本题相对原题做了扩展

示例:

 输入：4
 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 解释: 4 皇后问题存在如下两个不同的解法。
[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
 */
 /**
  *经典回溯算法
  */
func solveNQueens(n int) [][]string {
	//声明返回值
	var result [][]string

	//判断
	if n <= 0 {
		return result
	}

	//初始化棋盘
	var data  = make ([][]string,n)
	for i := 0 ; i < n ; i ++ {
		data[i] = make([]string,n)
		for j := 0; j < n; j ++ {
			data[i][j] = "." //表示没有皇后
		}
	}

	//回溯求解
	solveQueens(data,1,n,&result)

	//返回
	return result

}

func solveQueens(data [][]string,cur int,n int,result *[][]string) {
	if cur == n + 1 {
		tmp := make([]string,n)
		//生成结果
		for i := 0; i < n; i ++ {
			str := ""
			for j := 0 ; j < n ; j ++ {
				str += data[i][j]
			}
			tmp[i] = str
		}
		*result = append(*result,tmp)
		return
	}

	//决策可以放置的位置
	for i := 0; i < n ; i ++ {
		//判断是否可以放置
		isOk := checkQueens(cur - 1,i,data)

		//如果不可以
		if !isOk {
			continue
		}

		//可以放置
		data[cur - 1][i] = "Q"

		//dfs
		solveQueens(data,cur + 1,n,result)

		//回溯
		data[cur - 1][i] = "."

	}

}
func checkQueens(i,j int,data [][]string) bool{
	//第一行
	if i == 0 {//第一行不冲突
		return true
	}

	//同一行不冲突---回溯了

	//判断同一列
	for k := 0; k < i ; k ++ {
		if data[k][j] == "Q" {
			return false
		}
	}

	//判断左对角线
	a,b := i - 1,j - 1
	for a >= 0 && b >= 0 {
		if data[a][b] == "Q" {
			return false
		}
		a --
		b--
	}

	//判断右对角
	a,b = i - 1,j + 1
	for a >= 0 && b < len(data) {
		if data[a][b] == "Q" {
			return false
		}
		a --
		b ++
	}

	//返回
	return true

}