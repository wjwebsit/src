package main
/**
	给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

3
解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
提示：

节点总数 <= 10000

 */
func pathSum(root *TreeNode, sum int) int {
	num := 0
	pathSumHS(root,sum,&num)
	return num
}

func pathSumHS(tree *TreeNode,target int,sum *int) []int {
	//判断
	if tree == nil {
		return  []int{}
	}

	//当前情况
	var result []int

	//左子树和
	left := pathSumHS(tree.Left,target,sum)

	//判断是否为空
	if len(left) != 0 {
		for _,v := range left {
			if v + tree.Val == target {
				*sum ++
			}
			result = append(result,v + tree.Val)

		}

	}
	//右子树和
	right := pathSumHS(tree.Right,target,sum)

	//判断是否为空
	if len(right) != 0 {
		for _,v := range right {
			if v + tree.Val == target {
				*sum ++
			}
			result = append(result,v + tree.Val)

		}
	}
	if tree.Val == target {
		*sum ++
	}
	result = append(result,tree.Val)
	return result
}