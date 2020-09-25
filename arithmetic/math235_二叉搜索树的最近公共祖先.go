package main
/**
	算法描述
	给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	//声明祖先
	var parent  *TreeNode

	//寻找pq
	cur := root
	for cur != nil {
		//判断
		if cur == p || cur == q {
			parent = cur
			break
		}

		//p,q位于cur左右子树
		if (p.Val < cur.Val && cur.Val < q.Val) || (p.Val > cur.Val && q.Val < cur.Val) {
			parent = cur
			break
		}

		//继续寻找
		if p.Val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	//返回
	return parent
}