package main

import "fmt"

/**
	班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。

给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。

 

示例 1：

输入：
[[1,1,0],
 [1,1,0],
 [0,0,1]]
输出：2
解释：已知学生 0 和学生 1 互为朋友，他们在一个朋友圈。
第2个学生自己在一个朋友圈。所以返回 2 。
示例 2：

输入：
[[1,1,0],
 [1,1,1],
 [0,1,1]]
输出：1
解释：已知学生 0 和学生 1 互为朋友，学生 1 和学生 2 互为朋友，所以学生 0 和学生 2 也是朋友，所以他们三个在一个朋友圈，返回 1 。
 */
/**
	采用并查集
 */
type friend struct {
	p *friend
	rank int //秩
	key int
}

func initFriend(x int) *friend{
	f := new(friend)
	f.key = x
	f.p = nil
	f.rank = 0
	return f
}
func findX(x *friend) *friend {
	for x .p != nil {
		x = x.p
	}
	return x

}
func unionXY(x,y *friend) *friend{
	//获取秩
	x1 := x.rank
	y1 := y.rank

	if x1 > y1 {
		y.p = x
		return x
	} else {
		x.p = y
		if x1 == y1 {
			x.rank ++
		}
		return y
	}
}



func findCircleNum(M [][]int) int {
	//获取行和列
	rows := len(M)
	if rows == 0 {
		return 0
	}
	cols := len(M[0])
	if cols == 0 {
		return 0
	}
	//初始化并查集
	hash := make(map[int]*friend)
	for i := 0; i < rows; i++ {
		hash[i] = initFriend(i)
	}

	N := rows
	//遍历朋友圈
	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			//获取i,j--表示为朋友
			if M[i][j] == 1 && findX(hash[i]) != findX(hash[j]) {
				//合并
				unionXY(findX(hash[i]),findX(hash[j]))
				N --

			}

		}
	}

	//统计好友
	return N

}
func main() {
	friends := [][]int{{1,1,0},{1,1,1},
	{0,1,1},
	}
	fmt.Println(findCircleNum(friends))
}

