package main
/**
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
 */
func isBalanced(root *TreeNode) bool {
	//判断根结点
	if root == nil {
		return true
	}

	//定义返回
	isValid := true

	//回溯求解
	validBalanced(root,&isValid)

	//返回
	return isValid
}

/**
获取子树高度并比较是否为平衡二叉树
 */
 func validBalanced(tree *TreeNode,isValid *bool) int{
 	//判断当前结点是否为叶子结点
 	if tree.Left == nil && tree.Right == nil {
 		return 0
	}
 	//表示已经找到不是高度平衡了--不用在继续了
 	if *isValid == false {
 		return 0
	}

 	//初始化当前结点左子树和右子树的高度
 	left,right := 0,0

 	//寻找左子树高度
 	if tree.Left != nil {
 		left += 1 + validBalanced(tree.Left,isValid)
	}

 	//寻找右子树
 	if tree.Right != nil {
		right += 1 + validBalanced(tree.Right,isValid)
	}

 	//比较高度
 	if left - right > 1 || right - left > 1 {
 		*isValid = false
	}

 	//正常返回最大的
 	if right > left {
 		return right
	} else {
		return left
	}
 }