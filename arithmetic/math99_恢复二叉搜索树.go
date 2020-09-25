package main
/**
	算法描述
	二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。

示例 1:

输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2
示例 2:

输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
进阶:

使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/recover-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
var  pre  *TreeNode
func recoverTree(root *TreeNode)  {
	pre = nil
	//声明
	var nodes []*TreeNode

	midBST(root,&nodes)

	//寻找序对
	s,e := 0,len(nodes) - 1


	//从前面找到一个
	for s < e && s + 1 <= e && nodes[s].Val < nodes[s + 1].Val {
		s ++
	}

	//从后面找到一个
	for s < e && e - 1 >= s && nodes[e].Val < nodes[e - 1].Val {
		e --
	}

	//交换
	nodes[s].Val,nodes[e].Val = nodes[e].Val,nodes[s].Val






}
//中序遍历
func midBST(tree *TreeNode,nodes *[]*TreeNode) {
	if tree != nil {
		midBST(tree.Left, nodes)

		//写入
		*nodes = append(*nodes, tree)

		if pre != nil && pre.Val > tree.Val {
			return
		}

		pre = tree

		//右
		midBST(tree.Right, nodes)
	}
}


