package main

import "strconv"

/**
	算法描述:
	给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func binaryTreePaths(root *TreeNode) []string {
	//声明返回值
	var result []string

	//判断
	if root == nil {
		return result
	}

	//求解
	binaryTreePathsHs(root,"",&result)

	//返回
	return result


}

func binaryTreePathsHs(root *TreeNode,path string,result *[]string) {
	//判断
	if root == nil {
		*result = append(*result,path)

	} else {
		if path == "" {
			path += strconv.Itoa(root.Val)
		} else {
			path += "->" + strconv.Itoa(root.Val)
		}

		binaryTreePathsHs(root.Left,path,result)
		binaryTreePathsHs(root.Right,path,result)

	}

}