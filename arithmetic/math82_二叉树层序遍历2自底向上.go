package main
/*
	例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
 */
func levelOrderBottom(root *TreeNode) [][]int {
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

		//从头部写入
		result = append([][]int{tmp},result...)

		//重置为下一层node
		level = nextLeval
	}

	//返回
	return result
}

