package main
/**
	算法描述:
	给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

 */
func oddEvenList(head *ListNode) *ListNode {
	//判断
	if head == nil {
		return head
	}
	//定义奇数链表开头
	var odd,even,p *ListNode
	odd = head

	//判断
	if head.Next != nil {
		even = head.Next
	}

	//获取
	if even == nil || even.Next == nil {//判断长度
		return head
	}
	//偶数链表开头
	evenH := even

	//遍历
	p = even.Next
	for p != nil {
		tmp := p.Next
		odd.Next = p
		odd = odd.Next

		//判断
		if tmp != nil {
			even.Next = tmp
			even = even.Next
		} else {
			break
		}

		p = tmp.Next

	}

	//拼接链表
	odd.Next = evenH
	even.Next = nil

	//返回
	return head

}