package main

/**
	算法描述
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */

 /**
 	定义二叉树的数据结构
  */
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
/**
	二叉树的中序遍历---非递归实现
 */
func inorderTraversal(root *TreeNode) []int {
	//声明返回
	var result []int

	//判断根
	if root == nil {
		return result
	}

	//寻找左--放入栈
	var stack []*TreeNode

	//放入根
	stack = append(stack,root)

	//准备工作寻找左--并放入栈
	p := root.Left
	for p != nil {
		//压入栈
		stack = append(stack,p)

		//继续左
		p = p.Left
	}

	//开始填充结果
	for len(stack) != 0 {
		//栈顶出栈
		topNode := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		//写入当前
		result = append(result,topNode.Val)

		//判断当前的右
		if topNode.Right == nil {
			continue
		}

		//当前右入栈
		stack = append(stack,topNode.Right)

		//获取当前的右的左
		p := topNode.Right.Left
		for p != nil {
			//压入栈
			stack = append(stack,p)

			//继续左
			p = p.Left
		}
	}

	//返回结果
	return result

}




