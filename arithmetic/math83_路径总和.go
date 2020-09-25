package main
/**
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例: 
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
 */
func hasPathSum(root *TreeNode, sum int) bool {
	//判断
	if root == nil {
		return false
	}

	//声明返回值 --默认不存在
	var isOk = false

	//采用回溯算法
	hasPathSumHS(root,root.Val,sum,&isOk)

	//返回
	return isOk

}

/**
	回溯算法
 */
 func hasPathSumHS(tree *TreeNode,tmpSum int,sum int,isOk *bool) {
 	//判断
 	if *isOk == true {
 		return
	}

 	//是否为叶子节点并且和是否相等
 	if tree.Left == nil && tree.Right == nil && tmpSum == sum{
 		*isOk = true
 		return
	}

 	//遍历左
 	if tree.Left != nil && *isOk == false {
 		hasPathSumHS(tree.Left,tmpSum + tree.Left.Val,sum,isOk)
	}

 	//遍历右
	 if tree.Right != nil && *isOk == false {
		 hasPathSumHS(tree.Right,tmpSum + tree.Right.Val,sum,isOk)
	 }
 }