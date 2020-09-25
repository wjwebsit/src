package main
/**
	问题描述
	给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
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

	//定义方向 0--从左往右 1--从右往左
	direct := 0

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
			if direct == 0 {//从左往右
				if node.Left != nil {
					nextLeval = append(nextLeval,node.Left)
				}

				if node.Right != nil {
					nextLeval = append(nextLeval,node.Right)
				}
			} else {
				//从右往左时需先遍历右孩子在遍历左孩子
				if node.Right != nil {
					nextLeval = append(nextLeval,node.Right)
				}

				if node.Left != nil {
					nextLeval = append(nextLeval,node.Left)
				}
			}
		}

		//写入结果集
		result = append(result,tmp)

		//重置为下一层node
		level = []*TreeNode{}
		for  i := len(nextLeval) - 1; i >= 0; i --{
			level = append(level,nextLeval[i])
		}

		//重置方向
		if direct == 0 {
			direct = 1
		} else {
			direct = 0
		}
	}

	//返回
	return result
}
