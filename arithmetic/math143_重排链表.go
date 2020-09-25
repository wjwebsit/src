package main
/**
	算法描述
	给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

 */
func reorderList(head *ListNode)  {
	//判断head
	if head == nil {
		return
	}

	//声明栈
	var stack []*ListNode

	//顺序遍历 ----入栈
	p := head
	for p != nil {
		stack = append(stack,p)
		p = p.Next
	}

	//利用双指针
	end := len(stack) - 1
	top := stack[end]

	//重组链表
	p = head
	for p.Next != top && p.Next != nil{
		//出栈
		if end - 1 >= 0{
			stack[end - 1].Next = nil
		}

		//连接
		top.Next = p.Next
		p.Next = top

		//重置p
		p = p.Next.Next

		//重置s,top
		end --
		top = stack[end]
	}
}