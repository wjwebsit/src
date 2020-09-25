package main
/**
	检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。设计一个算法，判断 T2 是否为 T1 的子树。

如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，也就是说，从节点 n 处把树砍断，得到的树与 T2 完全相同。

示例1:

 输入：t1 = [1, 2, 3], t2 = [2]
 输出：true
示例2:

 输入：t1 = [1, null, 2, 4], t2 = [3, 2]
 输出：false
提示：

树的节点数目范围为[0, 20000]。
 */
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	//判断
	if t1 == nil {
		return t2 == nil
	}
	if t2 == nil {
		return true
	}
	
	//层序遍历
	var queue []*TreeNode
	queue = append(queue,t1)
	for len(queue) > 0 {
		//出队
		tree := queue[0]
		queue = queue[1:]
		
		//判断子树
		res := theSameTree(tree,t2)
		if res == true {
			return true
		}
		
		//左右子树
		if tree.Left != nil {
			queue = append(queue,tree.Left)
		}
		if tree.Right != nil {
			queue = append(queue,tree.Right)
		}
	}
	
	//不合法
	return false
	
}
/**
	判断相同的树
 */
func theSameTree(t1,t2 *TreeNode) bool{
	//判断不同
	if (t1 == nil && t2 != nil) || (t2 == nil && t1 != nil) {
		return false
	} 
	if t1 == nil && t2 == nil {
		return true
	}
	if t1.Val != t2.Val {
		return false
	}
	
	//左右子树
	return theSameTree(t1.Left,t2.Left) && theSameTree(t1.Right,t2.Right)
}