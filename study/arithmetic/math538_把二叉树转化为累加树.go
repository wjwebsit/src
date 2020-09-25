package main
/**
	算法描述:
	给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。

 

例如：

输入: 原始二叉搜索树:
              5
            /   \
           2     13

输出: 转换为累加树:
             18
            /   \
          20     13
 

注意：本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-bst-to-greater-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
var p *TreeNode //前驱结点
func convertBST(root *TreeNode) *TreeNode {
	//判断
	if root == nil {
		return root
	}

	//声明父
	var trees []*TreeNode

	//中序遍历
	convertBSTIn(root,&trees)

	//累加
	for i := len(trees) - 2; i >= 0; i -- {
		trees[i].Val += trees[i + 1].Val
	}

	//返回
	return root

}

//中序遍历---记录
func convertBSTIn(tree *TreeNode, trees *[]*TreeNode) {
	if tree != nil {
		//left
		convertBSTIn(tree.Left,trees)

		//记录
		*trees = append(*trees,tree)

		//right
		convertBSTIn(tree.Right,trees)
	}

}

func BSTIn(root *TreeNode,sum int) {
	if root != nil {
		//右孩子
		BSTIn(root.Right,sum)

		//当前
		root.Val += sum
		sum = root.Val

		//左孩子
		BSTIn(root.Left,sum)
	}

}