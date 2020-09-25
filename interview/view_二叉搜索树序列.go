package main
/**
	算法描述
	从左向右遍历一个数组，通过不断将其中的元素插入树中可以逐步地生成一棵二叉搜索树。给定一个由不同节点组成的二叉树，输出所有可能生成此树的数组。

示例:
给定如下二叉树

        2
       / \
      1   3
返回:

[
   [2,1,3],
   [2,3,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bst-sequences-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func BSTSequences(root *TreeNode) [][]int {
	//返回值
	var result [][]int

	//判断
	if root == nil {
		result = append(result,[]int{})
	} else {
		path := []int{root.Val}
		queue := []*TreeNode{}
		DfsBst(root,queue,path,&result)
	}

	return result
}

func DfsBst(tree *TreeNode,queue []*TreeNode, path []int,result *[][]int) {
	if tree.Left != nil {
		queue = append(queue,tree.Left)
	}
	if tree.Right != nil {
		queue = append(queue,tree.Right)
	}

	//判断队列
	if len(queue) == 0 {
		*result = append(*result,path)
		return
	}

	//遍历队列
	for i := 0 ; i < len(queue); i ++ {
		//获取元素
		node := queue[i]

		newQ := append(queue[0:i],queue[i + 1:]...)
		path = append(path,node.Val)

		//遍历
		DfsBst(node,newQ,path,result)

		//回溯
		queue[i] = node
		path = path[:len(path) - 1]
	}

}