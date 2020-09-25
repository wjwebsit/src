package main
/**
实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。


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
	isOk := true
	isBalancedHs(root,&isOk)
	return isOk
}

func isBalancedHs(tree *TreeNode,isOK *bool) int {
	//判断--如果没有则结束
	if !(*isOK) {
		return -1
	}

	//判断是否为nil
	if tree == nil {
		return  -1
	}

	//获取左子树高度
	left := 1 + isBalancedHs(tree.Left,isOK)

	//获取右子树高度
	right := 1 + isBalancedHs(tree.Right,isOK)

	//判断
	if left > right && left - right > 1 {
		*isOK = false
	}
	if right > left && right - left > 1 {
		*isOK = false
	}

	//返回
	return maxAB(left,right)


}

func maxAB(a,b int) int{
	if a > b {
		return a
	}
	return b
}

