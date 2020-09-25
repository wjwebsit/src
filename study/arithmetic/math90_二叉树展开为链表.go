package main

import "fmt"

/**
	算法描述
	给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/
func flatten(root *TreeNode)  {
	//判断根结点
	if root == nil {
		return
	}

	//获取当前的左右子孩子
	left,right := root.Left,root.Right

	//生成链表
	p := makeFlatten(left,root)
	makeFlatten(right,p)
}
/**
	先序遍历
	tree当前结点，父结点
 */
func makeFlatten(tree *TreeNode,p *TreeNode) *TreeNode {
	//判断当前
	if tree != nil {
		//构造链表
		p.Right = tree
		p.Left = nil

		//左孩子和有孩子
		if tree.Left != nil {
			//记录右孩子
			right := tree.Right

			//左子树
			p := makeFlatten(tree.Left, tree)

			//右子树时p应为左子树
			 return makeFlatten(right, p)

		} else {//构造又子树
			return makeFlatten(tree.Right, tree)
		}

	} else {
		return p
	}
}