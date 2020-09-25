package main
/**
	算法描述
	设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。

如果指定节点没有对应的“下一个”节点，则返回null。

示例 1:

输入: root = [2,1,3], p = 1

  2
 / \
1   3

输出: 2
示例 2:

输入: root = [5,3,6,2,4,null,null,1], p = 6

      5
     / \
    3   6
   / \
  2   4
 /
1

输出: null
 */
var Pre *TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	//判断p是否存在右孩子
	if p.Right != nil {//从左孩子中寻找个最小的
		rt := p.Right
		for rt.Left != nil {
			rt = rt.Left
		}

		//返回
		return rt
	}

	Pre = nil
	var res *TreeNode
	midBl(root,p,&res)
	return res

}

//中序遍历
func midBl(tree *TreeNode,p *TreeNode,res **TreeNode){
	//判断
	if tree != nil && *res == nil {
		midBl(tree.Left,p,res)

		//判断前驱
		if Pre != nil && Pre == p {
			*res = tree
		}
		Pre = tree
		midBl(tree.Right,p,res)
	}
}
