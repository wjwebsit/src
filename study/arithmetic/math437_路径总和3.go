package main
/**
	算法描述
	给定一个二叉树，它的每个结点都存放着一个整数值。

找出路径和等于给定数值的路径总数。

路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5  -3
   / \    \
  8   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11

 */
func pathSum3(root *TreeNode, sum int) int {
	//利用回溯算法
	count  := 0
	pathSum3HS(root,sum,[]int{},&count)
	return  count
}
func pathSum3HS(tree *TreeNode, sum int,result []int,count *int) {
	//如果树为nil
	if tree == nil {
		return
	}
	//当前结果和之前结果累加 看是否相等 还有把当前值加入结果
	if len(result) == 0 {
		result = append(result,tree.Val)
	} else {
		//累加
		for i := 0; i < len(result);i++ {
			if result[i] + tree.Val == sum {
				*count ++
			}
			result[i] += tree.Val
		}
		//当前
		result = append(result,tree.Val)
	}

	//当前
	if tree.Val == sum {
		*count ++
	}

	//左右子树
	if tree.Left != nil && tree.Right != nil {
		tmp := make([]int,len(result))
		copy(tmp,result)
		pathSum3HS(tree.Left, sum, result, count)
		pathSum3HS(tree.Right, sum, tmp, count)
	} else {
		pathSum3HS(tree.Left, sum, result, count)
		pathSum3HS(tree.Right, sum, result, count)
	}
}