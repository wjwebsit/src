package main
/**
	设计题描述：
	具体描述详见leetCode
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	this := new(BSTIterator)
	if root != nil {
		initStack(this, root)
	}

	//返回
	return  *this
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	//出栈
	top := this.stack[len(this.stack) - 1]
	this.stack = this.stack[:len(this.stack) - 1]

	//判断当前的右孩子
	if top.Right != nil {
		initStack(this,top.Right)
	}
	return  top.Val

}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return  len(this.stack) != 0
}

func initStack(this *BSTIterator,node *TreeNode) {
	//以当前结点为根来寻找最左结点
	this.stack = append(this.stack,node)

	//寻找最左
	left := node.Left
	for left != nil {
		//写入栈
		this.stack = append(this.stack,left)
		left = left.Left
	}
}