package main
/**
	定义数据结构---链表
 */
type ListNode struct {
	Val int
	Next *ListNode
}

/**
	定义数据结构----树
 */
type TreeNode struct {
	Val int
	Left,Right *TreeNode
}
