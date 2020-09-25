package main

import "fmt"

/**
	算法描述:
	根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
 */
 func buildTree(preorder []int, inorder []int) *TreeNode{
 	//声明root
 	var root *TreeNode

 	//判断
 	if len(preorder) == 0 {
 		return root
	}

 	//构造结点
 	root = new (TreeNode)
 	root.Val = preorder[0]

 	//从中序遍历中获取位置
 	var i int
 	for i = 0; i < len(inorder); i ++ {
 		if inorder[i] == root.Val {
 			break
		}
	}

 	//[0,i) 为左子树 ， [i + 1: len(inorder))为右子树
 	root.Left = buildTree(preorder[1:1 + i],inorder[:i])
 	root.Right = buildTree(preorder[1 + i :],inorder[i + 1:])

 	//返回
 	return root
 }
func main() {


}