package main
/**
	给定一颗二叉树求其最大深度
示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

 */

func maxDepth(root *TreeNode) int {
	//判断根结点
	if root == nil {
		return 0
	}

	//定义最大深度
	max := 1

	//回溯求解
	maxDeepHS(root,1,&max)

	//返回最大深度
	return max

}
/**
*@param tree *TreeNode 当前结点
*@param  tmpDeep int 当前深度
*@param max *int 最大深度
*/
func maxDeepHS(tree *TreeNode,tmpDeep int, max *int) {
	//判断当前是否为叶子结点
	if tree.Left == nil && tree.Right == nil {
		//比较
		if tmpDeep > *max {
			*max = tmpDeep
		}
		return
	}

	//回溯
	if tree.Left != nil {
		maxDeepHS(tree.Left,tmpDeep + 1, max)
	}

	//同理
	if tree.Right != nil {
		maxDeepHS(tree.Right,tmpDeep + 1, max)
	}

}
