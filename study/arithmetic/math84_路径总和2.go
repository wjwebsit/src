package main
/**
	算法描述
	给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

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

[
   [5,4,11,2],
   [5,8,4,5]
]
 */
func pathSum(root *TreeNode, sum int) [][]int {
	//声明返回结果
	var result [][]int

	//判断
	if root == nil {
		return result
	}

	//声明临时结果集
	var tmpRes []int
	tmpRes = append(tmpRes,root.Val)

	//利用回溯算法
	pathSumHS(root,root.Val,sum,tmpRes,&result)

	//返回结果集
	return result
}

/**
回溯求解
 */
 func pathSumHS(tree *TreeNode,tmpSum,sum int,tmpRes []int,result *[][]int) {

	 //是否为叶子节点并且和是否相等
	 if tree.Left == nil && tree.Right == nil && tmpSum == sum{
		 //写入结果集深拷贝
		 tmpRes1 := make([]int,len(tmpRes))
		 copy(tmpRes1,tmpRes)
		 *result = append(*result,tmpRes1)
		 return
	 }

	 //遍历左
	 if tree.Left != nil{
	 	 //写入结果集缓存
	 	 tmpRes = append(tmpRes,tree.Left.Val)

	 	 //DFS
		 pathSumHS(tree.Left,tmpSum + tree.Left.Val,sum,tmpRes,result)

	 	 //回溯
	 	 tmpRes = tmpRes[0:len(tmpRes) - 1]
	 }

	 //遍历右
	 if tree.Right != nil{
		 //写入结果集缓存
		 tmpRes = append(tmpRes,tree.Right.Val)

		 //DFS
		 pathSumHS(tree.Right,tmpSum + tree.Right.Val,sum,tmpRes,result)

		 //回溯
		 tmpRes = tmpRes[0:len(tmpRes) - 1]
	 }

 }