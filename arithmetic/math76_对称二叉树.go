package main

/**
	定义二叉树的数据结构
 */
 /*type treeNode struct {
 	Val int
 	left,right *treeNode
 }*/

 /**
 	解决
  */
func isSymmetric(root *treeNode) bool {
	//如果为空 返回false
	if root == nil {
		//返回
		return true
	}

	//声明返回变量
	isSym := true

	//回溯解决
	isSymTree(root.left,root.right,&isSym)

	//返回
	return isSym
}

/*
	是否对称判断
 */
 func isSymTree(left,right *treeNode,isSym *bool) {
 	//判断是否终止
 	if *isSym == false || (left == nil && right == nil) {
 		return
	}

 	//不对称情况
 	if (left == nil && right != nil) || (right == nil && left != nil) || left.Val != right.Val {
 		//修改标志变量
 		*isSym = false
 		return
	}

	 //回溯求解 left.left 和 right.right
	 if *isSym == true {
		 isSymTree(left.left, right.right, isSym)
	 }

 	//回溯求解 left.right 和 right.left
 	 if *isSym == true {
		 isSymTree(left.right, right.left, isSym)
	 }
 }
