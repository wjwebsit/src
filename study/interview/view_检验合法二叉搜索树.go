package main
/**
	算法描述：
	实现一个函数，检查一棵二叉树是否为二叉搜索树。

示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
 */

var pre *TreeNode //前驱节点
/**
	线索二叉树
 */
func isValidBST(root *TreeNode) bool {
	//判断
	if root == nil {
		return true
	}
	pre = nil
	isValid := true
	midBlBST(root,&isValid)
	return isValid
}

func midBlBST(node *TreeNode,isValid *bool) {
	//判断是否合法
	if *isValid == false {
		return
	}

	//判断
	if node != nil {
		//遍历左
		midBlBST(node.Left,isValid)

		//判断前驱
		if pre != nil && pre.Val >= node.Val {
			*isValid = false
			return
		}

		//记录前驱
		pre = node

		//右
		midBlBST(node.Right,isValid)
	}

}