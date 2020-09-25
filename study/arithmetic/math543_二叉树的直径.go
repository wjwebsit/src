package main
/**
	算法描述:
	给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

 

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

 

注意：两结点之间的路径长度是以它们之间边的数目表示。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//左子树最大路径+右子树最大路径
	max := 0
	getMaxBtreePath(root,&max)

	//返回最大
	return max

}

func getMaxBtreePath(tree *TreeNode,max *int) int {
	//判断叶子结点
	if tree.Left == nil && tree.Right == nil {
		return 0
	}

	//左子树
	L,R := 0,0 //以tree为根找左右路径
	if tree.Left != nil {
		L = 1 + getMaxBtreePath(tree.Left, max)
	}
	if tree.Right != nil {
		R = 1 + getMaxBtreePath(tree.Right, max)
	}

	//比较
	if L + R > *max {
		*max = L + R
	}

	//返回
	return L + R

}
