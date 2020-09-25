package main
/**
	算法描述:
	给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
 */
func rightSideView(root *TreeNode) []int {
	//声明返回值
	var result []int

	//判断树是否为空
	if root == nil {
		return result
	}

	//利用二叉树的层序遍历
	var queue []*TreeNode
	queue = append(queue,root)

	for len(queue) > 0 {
		//右视图即为队列最后一个元素----栈
		result = append(result,queue[len(queue) - 1].Val)

		//构造下一层数据
		var tmp []*TreeNode
		for _,v := range queue {
			//写入数据
			if v.Left != nil {
				tmp = append(tmp,v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp,v.Right)
			}
		}

		//重置
		queue = tmp
	}

	//返回
	return result
}
