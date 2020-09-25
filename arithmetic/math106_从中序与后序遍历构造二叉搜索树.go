package main
/**
	算法描述
	根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
 */
func buildTree2(inorder []int, postorder []int)*TreeNode {
	//声明根
	var root *TreeNode

	//判断后序遍历的最后结点
	if len(postorder) == 0 {
		return root
	}

	//构造根
	root = new(TreeNode)
	root.Val = postorder[len(postorder) - 1]

	//获取在中序遍历的位置
	var i int
	for i = 0; i < len(inorder);i++ {
		if inorder[i] == root.Val {
			break
		}
	}


	//构造左子树和右子树
	root.Left = buildTree2(inorder[:i],postorder[0 : i]);
	root.Right = buildTree2(inorder[i + 1:],postorder[i: len(postorder) - 1])

	//返回
	return root

}