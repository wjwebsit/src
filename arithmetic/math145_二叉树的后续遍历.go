package main


/**
	算法描述:
	给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
 */
func postorderTraversal(root *TreeNode) []int {
	//返回值
	var result []int

	//判断树是否为空
	if root == nil {
		return result
	}

	//声明栈和缓存
	var stack []*TreeNode
	var cache []*TreeNode

	//将根结点写入stack
	stack = append(stack,root)

	//循环
	for len(stack) > 0 {
		//获取当前栈顶元素--先不要出栈
		top := stack[len(stack) - 1]

		//定义出栈操作符
		fg1,fg2 := true,true
		//先判断右孩子
		if top.Right != nil && !isPassTreeNode(top.Right,cache) {
			//入栈
			stack = append(stack,top.Right)

			//访问数组写入
			cache = append(cache,top.Right)

			//操作符1为false
			fg1 = false
		}
		if top.Left != nil && !isPassTreeNode(top.Left,cache) {
			//入栈
			stack = append(stack,top.Left)

			//访问数组写入
			cache = append(cache,top.Left)

			//操作符2设为false
			fg2 = false

		}

		//操作符均为true
		if fg1 && fg2 {
			//输出并出栈
			result = append(result,top.Val)
			stack = stack[:len(stack) - 1]
		}

	}

	//返回
	return result

}

func isPassTreeNode(node *TreeNode,cache []*TreeNode) bool {
	//声明返回
	rt := false
	for _,v := range cache {
		if v == node {
			rt = true
			break
		}
	}
	//返回
	return rt
}