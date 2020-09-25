package main

import "fmt"

/**
	算法描述：
	给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例:

输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

 */
func generateTrees(n int) []*TreeNode {
	//声明返回值
	var tree []*TreeNode

	//判断
	if n <= 0 {
		return tree
	}
	if n == 1 {
		root := new(TreeNode)
		root.Val = 1
		tree = append(tree,root)
		return tree
	}


	//循环遍历根结点
	for i := 1; i <= n; i ++ {

		//声明左右孩子
		var left,right []*TreeNode

		//获取左孩子
		left = makeBST(1,i - 1)


		//获取右孩子
		right = makeBST(i + 1,n)

		//判断
		if i == 1 {//只有右孩子
			for _,v := range right{
				//构造当前的root
				root := new(TreeNode)
				root.Val = i
				root.Right = v
				tree = append(tree,root)
			}


		} else if i == n {//只有左孩子
			for _,v := range left {
				//构造当前的root
				root := new(TreeNode)
				root.Val = i
				root.Left = v
				tree = append(tree,root)
			}

		} else {//既有左孩子又有右孩子---笛卡尔积
			for _,l := range left {
				for _,r := range right {
					//构造当前的root
					root := new(TreeNode)
					root.Val = i
					root.Left = l
					root.Right = r
					tree = append(tree,root)
				}
			}

		}
	}

	//返回结果
	return tree
}
/**
	构造二叉搜索树
 */
func makeBST(start,end int) []*TreeNode {
	//声明返回
	var tree []*TreeNode

	//判断极端条件
	if start > end {
		return  tree
	}

	//判断是否相等--只有一个结点
	if start == end {
		root := new(TreeNode)
		root.Val = start
		tree = append(tree,root)
		return tree
	}


	//遍历根结点
	for i := start; i <= end;i++ {

		//声明左右孩子
		var left,right []*TreeNode

		//获取左孩子
		left = makeBST(start,i - 1)


		//获取右孩子
		right = makeBST(i + 1,end)

		//判断
		if i == start {//只有右孩子
			for _,v := range right{
				//构造当前的root
				root := new(TreeNode)
				root.Val = i
				root.Right = v
				tree = append(tree,root)
			}


		} else if i == end {//只有左孩子
			for _,v := range left {
				//构造当前的root
				root := new(TreeNode)
				root.Val = i
				root.Left = v
				tree = append(tree,root)
			}

		} else {//既有左孩子又有右孩子---笛卡尔积
			for _,l := range left {
				for _,r := range right {
					//构造当前的root
					root := new(TreeNode)
					root.Val = i
					root.Left = l
					root.Right = r
					tree = append(tree,root)
				}
			}

		}
	}

	//返回
	return tree

}

