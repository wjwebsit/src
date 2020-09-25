package main
/**
	算法描述
	设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。注意：这不一定是二叉搜索树。

例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]

    3
   / \
  5   1
 / \ / \
6  2 0  8
  / \
 7   4
示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
 */
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	//声明hash
	var  tmp,parent []*TreeNode //存储p的祖先--包括自身
	tmp = append(tmp,root)


	getParents(root,p,tmp,&parent)

	//获取祖先
	if len(parent) == 0 {
		return nil
	}

	//寻找q祖先
	for i := len(parent) - 1; i >= 0 ; i -- {
		if getParents(parent[i],q,[]*TreeNode{},&[]*TreeNode{}) {
			return parent[i]
		}

	}

	return nil

}

func getParents(tree *TreeNode,p *TreeNode,tmp1 []*TreeNode,parent *[]*TreeNode) bool{
	//判断
	if tree == nil {
		return false
	}

	//判断
	if tree == p {
		tmp1 = append(tmp1,tree)

		//写入结果
		(*parent) = make([]*TreeNode,len(tmp1))
		copy(*parent,tmp1)

		return true
	}

	var res bool = false

	//寻找左
	if tree.Left != nil {
		tmp1 = append(tmp1,tree.Left)
		res = getParents(tree.Left,p,tmp1,parent)

		//回溯
		if res == false {
			tmp1 = tmp1[0 : len(tmp1)-1]
		}
	}

	//寻找右
	if res == true {
		return  res
	}

	if tree.Right != nil {
		tmp1 = append(tmp1,tree.Right)
		res = getParents(tree.Right,p,tmp1,parent)

		//回溯
		if res == false {
			tmp1 = tmp1[0 : len(tmp1)-1]
		}
	}
	return res

}