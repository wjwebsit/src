package m
/**
	算法描述

 */
import "C"
type Node struct {
	     Val bool
	     IsLeaf bool
	     TopLeft *Node
	     TopRight *Node
	     BottomLeft *Node
	     BottomRight *Node
}

func construct(grid [][]int) *Node {
	//获取n
	n := len(grid)
	if n == 0 {
		return nil
	}

	//计算前缀和
	sum := make([][]int,n + 1)

	//初始化
	for i := 0;i < len(sum);i++ {
		sum[i] = make([]int,n + 1)
	}

	//求和
	for i := 1; i <= n; i ++ {
		for j := 1 ; j <= n ; j ++ {
			sum[i][j] = sum[i][j - 1] + sum[i - 1][j] + grid[i - 1][j - 1] - sum[i - 1][j - 1]
		}
	}

	//初始化
	root := new(Node)

	//判断是否为叶子结点
	if sum[n][n] == 0 || sum[n][n] == n * n {
		root.IsLeaf = true
		if sum[n][n] == 0 {
			root.Val = false
		} else {
			root.Val = true
		}
		return root
	}

	//非叶子结点
	root.IsLeaf  = false
	root.Val = true
	root.TopLeft = myConstruct(sum,0,0,n / 2 - 1,n / 2 - 1,n)
	root.TopRight = myConstruct(sum,0,n / 2 ,n / 2 - 1,n - 1,n)
	root.BottomLeft = myConstruct(sum,n / 2,0,n - 1,n /2 - 1,n)
	root.BottomRight = myConstruct(sum,n / 2,n / 2 ,n - 1,n - 1,n)

	//返回
	return root
}

func myConstruct(sum [][]int,sR,eR,sC,eC int,n int) *Node {
	//判断
	if n < 0 {
		return nil
	}

	//拆分
	n /= 2

	//初始化当前节点
	node := new(Node)

	//判断当前是否为叶子结点
	s := sum[sC + 1][eC + 1] - sum[sC + 1][eR] - sum[sR][eC + 1] + sum[sR][eR]

	//判断
	if s == n * n || s == 0 {
		node.IsLeaf = true
		if s == 0 {
			node.Val = false
		} else {
			node.Val = true
		}
	} else {//分割数组
		node.TopLeft = myConstruct(sum,sR,eR,sR + n / 2 - 1,eR + n / 2 - 1,n)
		node.TopRight = myConstruct(sum,sR,eR + n / 2,sR + n / 2 - 1,eC,n)

		node.BottomLeft = myConstruct(sum,sR + n /2,eR,sC,eR + n / 2 - 1,n)
		node.BottomRight = myConstruct(sum,sR + n / 2,eR + n/2 ,sC,eC,n)
		node.Val = true
		node.IsLeaf = false
	}

	return node
}