package main
/**
	算法描述
	给定一个有环链表，实现一个算法返回环路的开头节点。
有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。


示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：

输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：

输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。

进阶：
你是否可以不用额外空间解决此题？
 */

 /**
 	弗洛伊德的乌龟和兔子
  */
func detectCycle(head *ListNode) *ListNode {
	//判断
	if head == nil {
		return nil
	}

	//快慢指针
	fast,slow := head,head
	for fast != nil && fast.Next != nil {
		if fast == slow {//表示有环
			break
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	//判断是否环
	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast

}