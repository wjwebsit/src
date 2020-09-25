package lbTable6
/**
	算法描述
	给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
 */
/**
	声明单链表结构体
 */
type ListNode struct {
	Val int
	Next *ListNode
}
/**
	思路类似于快排序子函数
 */
func partition(head *ListNode, x int) *ListNode {
	//判断
	if head == nil || head.Next == nil {
		return head
	}

	//定义辅助变量
	var top,pre,firstNode,preLast *ListNode
	first := 1

	//循环
	p := head
	for p != nil {
		temp := p.Next
		//判断当前是否小于x
		if p.Val < x {
			//是否为第一次
			if top == nil {
				firstNode = p
				top = p
			} else  {
				pre.Next = p.Next
				p.Next = top.Next
				top.Next = p
				top = p

			}
			pre = p

		} else {
			first ++
			pre = p
			if top == nil {
				preLast = p
			}
		}
		p = temp

	}

	//判断
	if first > 1 {
		preLast.Next = top.Next
		top.Next = head
	}

	//返回
	return firstNode


}