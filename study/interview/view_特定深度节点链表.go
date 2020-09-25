package main
/**
	给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

 

示例：

输入：[1,2,3,4,5,null,7,8]

        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/list-of-depth-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func listOfDepth(tree *TreeNode) []*ListNode {
	//声明返回值
	var result []*ListNode

	//判断根
	if tree == nil {
		return result
	}

	//声明辅助数组
	var nodes []*TreeNode
	nodes = append(nodes,tree) //根写入

	//遍历
	for len(nodes) > 0 {
		//生成链表---采用尾插法
		head := new(ListNode)
		head.Val = nodes[0].Val
		tail := head

		//临时数组存储下一层结点
		var tmp []*TreeNode
		if nodes[0].Left != nil {
			tmp = append(tmp,nodes[0].Left)
		}
		if nodes[0].Right != nil {
			tmp = append(tmp,nodes[0].Right)
		}

		//遍历
		for i := 1; i < len(nodes); i ++ {
			//生成节点
			newNode := new(ListNode)
			newNode.Val = nodes[i].Val
			tail.Next = newNode
			tail = tail.Next

			if nodes[i].Left != nil {
				tmp = append(tmp,nodes[i].Left)
			}
			if nodes[i].Right != nil {
				tmp = append(tmp,nodes[i].Right)
			}
		}

		//重置
		nodes = tmp

		//写入结果集
		result = append(result,head)

	}

	//返回
	return result



}