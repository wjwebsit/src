package main
/**
	算法描述：
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
 */

var pre *TreeNode
/**
	类似线索二叉树思路记录pre 中序遍历为升序故pre.val > current.val时false
 */
func isValidBST(root *TreeNode) bool {
	fg := true
	LDRBTSCheck(root,&fg)

	//返回
	return fg
}
func LDRBTSCheck(root *TreeNode,fg *bool)  {
	//判断当前
	if root != nil && *fg == true {
		//遍历左
		 LDRBTSCheck(root.Left,fg)

		//判断
		if pre != nil && pre.Val >= root.Val {
			*fg = false
		}

		//获取上一任结点
		pre = root

		//遍历右
		 LDRBTSCheck(root.Right,fg)
	}
}

/**
	基于队列来验证---不正确 ---中序遍历为递增
 */
func myCheckBTS(root *TreeNode) bool {
	//声明返回值
	flag := true

	//判断当前的根
	if root == nil {
		return flag
	}

	//获取当前root的值
	rootVal := root.Val

	//判断左树
	if root.Left != nil {
		//声明队列
		var left  [] *TreeNode
		left = append(left,root.Left)

		//判断
		for len(left) > 0 {
			//出队
			p := left[len(left) - 1]
			left = left[0 :len(left) - 1]

			//判断当前
			if p.Val >= rootVal {
				flag = false
				break
			}

			//判断当前的左孩子
			if p.Left != nil {
				//判断值
				if p.Left.Val >= p.Val {
					flag = false
					break
				}

				//入队
				left = append(left,p.Left)
			}

			//判断当前右孩子
			if p.Right != nil {
				//判断值
				if p.Right.Val <= p.Val {
					flag = false
					break
				}

				//入队
				left = append(left,p.Right)
			}

		}

	}
	//判断右树
	if root.Right != nil && flag{
		//声明队列
		var right  [] *TreeNode
		right = append(right,root.Right)

		//判断
		for len(right) > 0 {
			//出队
			p := right[len(right) - 1]
			right = right[0 :len(right) - 1]

			//判断当前
			if p.Val <= rootVal {
				flag = false
				break
			}

			//判断当前的左孩子
			if p.Left != nil {
				//判断值
				if p.Left.Val >= p.Val {
					flag = false
					break
				}

				//入队
				right = append(right,p.Left)
			}

			//判断当前右孩子
			if p.Right != nil {
				//判断值
				if p.Right.Val <= p.Val {
					flag = false
					break
				}

				//入队
				right = append(right,p.Right)
			}

		}

	}

	//返回
	return flag


}
