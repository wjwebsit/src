package main
/**
	算法描述:
	反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
 */
type ListNode struct {
	 Val int
	 Next * ListNode
}

 func reverseList(head *ListNode) *ListNode{
	//判断链表是否为nil
	if head == nil || head.Next == nil{
		return head
	}

	//反转
	var node *ListNode

	//起始为头
	node = head
	p := head.Next
	node.Next = nil
	for p != nil {
		tmp := p.Next
		//反转
		p.Next = node
		node = p

		p = tmp

	}
	return node
 }
