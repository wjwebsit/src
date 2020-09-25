package main

/**
	题目描述
	给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

 */
func levelOrder(root *TreeNode) [][]int {
	//声明返回结果
	var result [][]int

	//判断
	if root == nil {
		return result
	}

	//声明层次切片
	var level []*TreeNode

	//根节点写入
	level = append(level,root)

	//遍历
	for len(level) > 0 {
		//声明临时数组
		var tmp []int

		//声明下一层结点
		var nextLeval [] *TreeNode

		//从左向右遍历
		for _,node := range level {
			//写入临时数组
			tmp = append(tmp,node.Val)

			//构造下一层结点
			if node.Left != nil {
				nextLeval = append(nextLeval,node.Left)
			}

			if node.Right != nil {
				nextLeval = append(nextLeval,node.Right)
			}
		}

		//写入结果集
		result = append(result,tmp)

		//重置为下一层node
		level = nextLeval
	}

	//返回
	return result
}
