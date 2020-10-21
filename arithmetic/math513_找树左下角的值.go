package main

/**
	算法描述
给定一个二叉树，在树的最后一行找到最左边的值。

示例 1:

输入:

    2
   / \
  1   3

输出:
1
 

示例 2:

输入:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

输出:
7
注意: 您可以假设树（即给定的根节点）不为 NULL
 */

func findBottomLeftValue(root *TreeNode) int {
	//寻找最后一层
	var last []*TreeNode
	last = append(last,root)

	for len(last) > 0 {
		//声明临时数组
		var tmp []*TreeNode

		//层序遍历
		for _,tree := range last {
			if tree.Left != nil {
				tmp = append(tmp,tree.Left)
			}
			if tree.Right != nil {
				tmp = append(tmp,tree.Right)
			}
		}


		//最后一层
		if len(tmp) == 0 {
			break
		}

		//next
		last = tmp
	}

	//返回
	return last[0].Val

}