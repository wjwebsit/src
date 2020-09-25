package main
/**
	题目描述:
	给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]

 */
func preorderTraversal(root *TreeNode) []int {
	//声明返回
	var result []int

	//判断根结点
	if root == nil {
		return result
	}

	//声明栈
	var stack []*TreeNode
	stack = append(stack,root)
	//写入

	//循环
	for len(stack) > 0 {
		//出栈
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		//写入结果
		result = append(result,top.Val)

		//先判断右孩子
		if top.Right != nil {
			stack = append(stack,top.Right)
		}

		//在判断左孩子
		if top.Left != nil {
			stack = append(stack,top.Left)
		}

	}

	//返回
	return result

}