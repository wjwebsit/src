package main

import "math"

/**
	算法描述
	给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。


示例：

输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
var pre *TreeNode
func getMinimumDifference(root *TreeNode) int {
	//设置pre
	pre = nil

	//返回值
	min := math.MaxInt64
	getMinAbs(root,&min)
	return  min

}
func getMinAbs(tree *TreeNode,min *int) {
	//判断
	if tree != nil {
		//左孩子
		getMinAbs(tree.Left,min)

		//判断前驱是否存在
		if pre != nil {
			//获取绝对值
			abs := getAbs(pre.Val,tree.Val)

			//比较
			if 	abs < *min {
				*min = abs
			}
		}

		//记录前驱
		pre = tree

		//获取右孩子
		getMinAbs(tree.Right,min)
	}

}

/**
	获取绝对值
 */
func getAbs(a ,b int) int{
	//获取最大值
	if a > b {
		return a - b
	}
	return b - a
}