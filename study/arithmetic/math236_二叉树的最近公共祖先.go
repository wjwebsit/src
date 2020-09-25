package main
/**
	算法描述：
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]



 

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
 

说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//判断
	if root == nil {
		return nil
	}

	//声明祖先
	var parents []*TreeNode

	//寻找p的祖先---由于必定存在不用判断是否存在了
	findAncestor(root,p.Val,&parents)

	//倒叙祖先寻找q
	var comParent *TreeNode

	for i := len(parents) - 1; i >= 0; i -- {
		if findAncestor(parents[i],q.Val,&[]*TreeNode{}) {
			comParent = parents[i]
			break
		}

	}

	//返回
	return comParent
}

/**
	利用回溯算法求解---结点的祖先
 */
func findAncestor(node *TreeNode,target int,parents *[]*TreeNode) bool {
	//判断当前值
	if node != nil {//当前结点非空
		//写入当前结点为祖先
		*parents = append(*parents,node)

		//判断是否相等
		if target == node.Val {
			return true
		}

		//从左孩子右孩子寻找
		res := findAncestor(node.Left,target,parents) || findAncestor(node.Right,target,parents)

		//判断
		if res == false {
			//回溯
			*parents = (*parents)[:len(*parents) - 1]
		}

		//返回结果
		return res

	} else {
		return  false
	}
}


