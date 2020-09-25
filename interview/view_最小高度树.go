package main
/**
	算法描述
	给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

示例:
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0
         / \
       -3   9
       /   /
     -10  5
 */
func sortedArrayToBST(nums []int) *TreeNode {
	//判断数组
	if len(nums) == 0 {
		return nil
	}
	//获取中点
	s,e := 0 , len(nums) - 1
	mid := s + (e -s) / 2


	//声明新节点
	node := new(TreeNode)
	node.Val = nums[mid]

	//左右
	node.Left = sortedArrayToBST(nums[0:mid])
	node.Right = sortedArrayToBST(nums[mid + 1 : e + 1])

	//返回
	return node

}