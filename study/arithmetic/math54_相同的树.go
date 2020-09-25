package tree1

/**
	算法描述
	给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true
示例 2:

输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false
示例 3:

输入:       1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

输出: false
 */




/**
	定义二叉树的数据结构
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
/**
	相同的树
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//定义返回值
	isTrue := true

	//定义2个队列
	var queueP,queueQ []*TreeNode

	//根节点入队
	if (p == nil && q != nil) || q == nil && p != nil {
		return false
	}
	if p != nil {//入队
		queueP = append(queueP, p)
		queueQ = append(queueQ, q)
	}

	//采用层序遍历来判断
	for len(queueP) > 0 && len(queueQ) > 0 {
		//当前出队
		tmp := queueP[0]
		queueP = queueP[1:]

		tmq := queueQ[0]
		queueQ = queueQ[1:]

		//判断当前
		if tmp.Val != tmq.Val {
			isTrue = false
			break
		}

		//判断左
		if (tmp.Left != nil && tmq.Left == nil) || (tmp.Left == nil && tmq.Left != nil) {
			isTrue = false
			break
		}

		//写入左
		if tmp.Left != nil {
			queueP = append(queueP,tmp.Left)
			queueQ = append(queueQ,tmq.Left)
		}

		//判断右
		if (tmp.Right != nil && tmq.Right == nil) || (tmp.Right == nil && tmq.Right != nil) {
			isTrue = false
			break
		}
		//写入左
		if tmp.Right != nil {
			queueP = append(queueP,tmp.Right)
			queueQ = append(queueQ,tmq.Right)
		}

	}

	//返回
	return isTrue

}

/**
	递归实现---leetCode
 */
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	//判断
	if p == nil && q== nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree1(p.Left,q.Left) && isSameTree1(p.Right,q.Right)
	} else {
		return false
	}
}