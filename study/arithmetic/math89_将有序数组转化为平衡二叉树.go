package main
/**
题目描述
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func sortedArrayToBST(nums []int) *TreeNode {
	//判断数组
	if len(nums) == 0 {
		return nil
	}

	//定义起止
	start,end := 0,len(nums) - 1

	//获取mid
	mid := (start + end) / 2

	//构造根
	var root *TreeNode
	root = new(TreeNode)
	root.Val = nums[mid]

	//构造左右子树
	makeBSTAVL(root,start,mid,nums)
	makeBSTAVL(root,mid,end,nums)

	//返回
	return root
}

func makeBSTAVL(parent *TreeNode,start,end int,nums []int) {
	//判断
	if start < end {
		//获取mid
		mid := (start + end) / 2

		//构造
		newNode := new (TreeNode)
		newNode.Val = nums[mid]

		//判断左右
		if newNode.Val < parent.Val {
			parent.Left = newNode
		} else {
			parent.Right = newNode
		}

		//构造当前的左右
		makeBSTAVL(newNode,start,mid,nums)
		makeBSTAVL(newNode,mid,end,nums)
	}
}