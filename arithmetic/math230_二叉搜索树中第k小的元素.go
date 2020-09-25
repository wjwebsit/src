package main
/**
	算法描述
	给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3

---进阶 ---基于红黑树来构建顺序统计树
 */
func kthSmallest(root *TreeNode, k int) int {
	//中序遍历
	var result ,cur  int
	cur = 1
	inOrderBST(root,k,&cur,&result)

	//返回
	return result
}
func inOrderBST(tree *TreeNode,k int,cur *int,result *int) bool {
	if tree != nil {
		//左孩子
		res := inOrderBST(tree.Left,k,cur,result)

		//判断找到
		if res == true {
			return true
		}

		//当前
		if k == *cur {
			*result = tree.Val
			return true
		}
		*cur ++

		//右孩子
		return  inOrderBST(tree.Right,k,cur,result)

	} else {
		return false
	}

}