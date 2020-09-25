package main

import "math"

/**

和二叉树的最大深度相反
 */
func minDepth(root *TreeNode) int {
	//判断根
	if root == nil {
		return 0
	}

	//定义返回值初始为正无穷
	min := math.MaxInt64

	//回溯
	minDepthHS(root,1,&min)

	//返回结果
	return min
}

/*
*回溯算法
 */
 func minDepthHS(tree *TreeNode,tmpDep int, minDep *int) {
 	//判断叶子结点
 	if tree.Left == nil && tree.Right == nil && tmpDep < *minDep {
 		*minDep = tmpDep
 		return
	}

 	//DFS
 	if tree.Left != nil {
		minDepthHS(tree.Left,tmpDep + 1,minDep)
	}

 	//同上
	 if tree.Right != nil {
		 minDepthHS(tree.Right,tmpDep + 1,minDep)
	 }
 }
